package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rekognition"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"

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
	//CleanupBucket(s)
	// Upload
	filename := os.Args[1]
	DetectFaces(s, filename)
	/*
		err = AddFileToS3(s, filename)
		if err != nil {
			log.Fatal(err)
		}*/
	//	DetectFaces(s, filename)
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

//DetectFaces Labels, Now I do not want to upload to S3 if no humans are detected. So
func DetectFaces(s *session.Session, filename string) {
	// Read the file to buffer
	imgFile, err := os.Open(filename)

	if err != nil {
		fmt.Printf("Oops some error %s\n", err.Error())
		os.Exit(1)
	}
	defer imgFile.Close()

	/*	fInfo, _ := imgFile.Stat() // So that we know the size of buffer to create
		var size int64 = fInfo.Size()
		buf := make([]byte, size) */ // Make a buffer with size we got earlier

	fReader := bufio.NewReader(imgFile) //Use bufio to read it to buffer
	content, _ := ioutil.ReadAll(fReader)

	//imgBase64Str, _ := base64.StdEncoding.DecodeString(content) //base64 encoded string

	svc := rekognition.New(s)
	/*
		input := &rekognition.DetectLabelsInput{
			Image: &rekognition.Image{
				S3Object: &rekognition.S3Object{
					Bucket: aws.String(Conf.s3_bucket),
					Name:   aws.String(filename),
				},
			},
			MaxLabels:     aws.Int64(123),
			MinConfidence: aws.Float64(70.000000),
		} */
	input := &rekognition.DetectFacesInput{
		Image: &rekognition.Image{
			Bytes: []byte(content),
		},
	}

	result, err := svc.DetectFaces(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case rekognition.ErrCodeInvalidS3ObjectException:
				fmt.Println(rekognition.ErrCodeInvalidS3ObjectException, aerr.Error())
			case rekognition.ErrCodeInvalidParameterException:
				fmt.Println(rekognition.ErrCodeInvalidParameterException, aerr.Error())
			case rekognition.ErrCodeImageTooLargeException:
				fmt.Println(rekognition.ErrCodeImageTooLargeException, aerr.Error())
			case rekognition.ErrCodeAccessDeniedException:
				fmt.Println(rekognition.ErrCodeAccessDeniedException, aerr.Error())
			case rekognition.ErrCodeInternalServerError:
				fmt.Println(rekognition.ErrCodeInternalServerError, aerr.Error())
			case rekognition.ErrCodeThrottlingException:
				fmt.Println(rekognition.ErrCodeThrottlingException, aerr.Error())
			case rekognition.ErrCodeProvisionedThroughputExceededException:
				fmt.Println(rekognition.ErrCodeProvisionedThroughputExceededException, aerr.Error())
			case rekognition.ErrCodeInvalidImageFormatException:
				fmt.Println(rekognition.ErrCodeInvalidImageFormatException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}
