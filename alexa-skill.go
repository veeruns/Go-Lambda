package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/davecgh/go-spew/spew"
	"bytes"
	"strconv"
)

type AlexaRequest struct {
	Version string `json:"version"`
	Request struct {
		Type   string `json:"type"`
		Time   string `json:"timestamp"`
		Intent struct {
			Name               string `json:"name"`
			ConfirmationStatus string `json:"confirmationstatus"`
			Slots struct {
				Times struct {
					Name string `json:name`
					Value string `json:value`
					ConfirmationStatus string `json:"confirmationstatus"`
					} `json:times`
			} `json:slots`
		} `json:"intent"`
	} `json:"request"`
}

type AlexaResponse struct {
	Version  string `json:"version"`
	Response struct {
		OutputSpeech struct {
			Type string `json:"type"`
			Text string `json:"text,omitempty"`
			SSML string `json:"ssml,omitempty"`
		} `json:"outputSpeech,omitemtpty"`
	} `json:"response"`
}

func CreateResponse(flag bool) *AlexaResponse {
	var resp AlexaResponse
	resp.Version = "1.0"
	if(flag){

	resp.Response.OutputSpeech.Type = "PlainText"
	resp.Response.OutputSpeech.Text = "Hello.  Please override this default output."

}else{
	resp.Response.OutputSpeech.Type="SSML"

	resp.Response.OutputSpeech.SSML="<speak> Hello, Please override this default SSML output. </speak>"
}
	return &resp
}



func (resp *AlexaResponse) Say(text string) {
	resp.Response.OutputSpeech.Text = text
}

func (resp *AlexaResponse) Ssay(text string){
	var b bytes.Buffer
	b.WriteString("<speak>")
	b.WriteString(text)
	b.WriteString("</speak>")
	resp.Response.OutputSpeech.SSML = b.String()
}

func (resp *AlexaResponse) NSsay(text string,number int){
	var b bytes.Buffer
	b.WriteString("<speak>")
	for i:=0;i < number;i++ {
	b.WriteString("<p>")
	b.WriteString(text)
	b.WriteString("</p>")
}
	b.WriteString("</speak>")
	resp.Response.OutputSpeech.SSML = b.String()
}

func HandleRequest(ctx context.Context, i AlexaRequest) (AlexaResponse, error) {
	// Use Spew to output the request for debugging purposes:
	fmt.Println("---- Dumping Input Map: ----")
	spew.Dump(i)
	fmt.Println("---- Done. ----")

	// Example of accessing map value via index:
	log.Printf("Request type is ", i.Request.Intent.Name)
  fmt.Println("Times is %s\n",i.Request.Intent.Slots.Times.Value)
	// Create a response object
	var resp *AlexaResponse

	// Customize the response for each Alexa Intent
	switch i.Request.Intent.Name {
	case "Eat":
		resp = CreateResponse(false)
		//resp.Say("Aarya Please <emphasis level='strong'> eat the food </emphasis>")
		resp.Ssay("Aarya Please <emphasis level='strong'> eat the food </emphasis>")
	case "hello":
		resp = CreateResponse(true)
		resp.Say("Hello there, Lambda appears to be working properly.")
	case "chew":
		resp=CreateResponse(false)
		number_of_time,_:=strconv.Atoi(i.Request.Intent.Slots.Times.Value)
		resp.NSsay("Aarya Please <emphasis level='strong'> chew the food </emphasis> ", number_of_time)
	case "AMAZON.HelpIntent":
		resp=CreateResponse(true)
		resp.Say("This app is easy to use, just say: ask the office how warm it is")

	default:
		resp=CreateResponse(true)
		resp.Say("I'm sorry, the input does not look like something I understand.")
	}

	return *resp, nil
}

func main() {
	lambda.Start(HandleRequest)
}
