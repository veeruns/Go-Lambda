package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"

	"path/filepath"

	"github.com/spf13/viper"
)

// TODO fill these in!
const (
	S3_REGION = "us-east-1"
	S3_BUCKET = "image.dump"
)

type Config struct {
	s3_bucket         string
	s3_region         string
	access_key_id     string
	access_key_secret string
}

var Conf Config

func main() {
	readconfig(&Conf, "/etc/motion", "creds")
	// Create a single AWS session (we can re use this if we're uploading many files)
	s, err := session.NewSession(&aws.Config{Credentials: credentials.NewStaticCredentials(Conf.access_key_id, Conf.access_key_secret, ""), Region: aws.String(Conf.s3_region)})
	if err != nil {
		log.Fatal(err)
	}
	CleanupBucket(s)
	// Upload
	filename := os.Args[1]
	err = AddFileToS3(s, filename)
	if err != nil {
		log.Fatal(err)
	}
}
func CleanupBucket(s *session.Session) bool {
	svc := s3.New(s)
	iter := s3manager.NewDeleteListIterator(svc, &s3.ListObjectsInput{
		Bucket: aws.String(Conf.s3_bucket),
	})

	// Traverse iterator deleting each object
	if err := s3manager.NewBatchDeleteWithClient(svc).Delete(aws.BackgroundContext(), iter); err != nil {
		fmt.Printf("Unable to delete objects from bucket %q, %v", Conf.s3_bucket, err)
		return false
	}
	return true

}

// AddFileToS3 will upload a single file to S3, it will require a pre-built aws session
// and will set file info like content type and encryption on the uploaded file.
func AddFileToS3(s *session.Session, fileDir string) error {

	// Open the file for use
	file, err := os.Open(fileDir)
	if err != nil {
		return err
	}
	defer file.Close()

	// Get file size and read the file content into a buffer
	fileInfo, _ := file.Stat()
	var size int64 = fileInfo.Size()
	buffer := make([]byte, size)
	file.Read(buffer)

	// Config settings: this is where you choose the bucket, filename, content-type etc.
	// of the file you're uploading.
	basepath := filepath.Base(fileDir)
	_, err = s3.New(s).PutObject(&s3.PutObjectInput{
		Bucket:               aws.String(Conf.s3_bucket),
		Key:                  aws.String(basepath),
		ACL:                  aws.String("private"),
		Body:                 bytes.NewReader(buffer),
		ContentLength:        aws.Int64(size),
		ContentType:          aws.String(http.DetectContentType(buffer)),
		ContentDisposition:   aws.String("attachment"),
		ServerSideEncryption: aws.String("AES256"),
	})
	return err
}

//Readconfig File
func readconfig(cfg *Config, confdir string, confname string) bool {
	viper.SetConfigName(confname)
	viper.AddConfigPath(confdir)
	err := viper.ReadInConfig()

	if err != nil {
		fmt.Printf("Config file not found...%s\n", err.Error())
		return false
	}
	//Server section
	cfg.access_key_id = viper.GetString("aws.access_key_id")
	cfg.access_key_secret = viper.GetString("aws.access_key_secret")
	cfg.s3_bucket = viper.GetString("aws.s3_bucket")
	cfg.s3_region = viper.GetString("aws.s3_region")

	return true

}
