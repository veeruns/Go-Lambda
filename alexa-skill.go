package main

import (
	"bytes"
	"context"
	"fmt"
	"strconv"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/davecgh/go-spew/spew"
)

type AlexaRequest struct {
	Version string `json:"version"`
	Request struct {
		Type      string `json:"type"`
		Time      string `json:"timestamp"`
		Locale    string `json:"locale"`
		RequestID string `json:"requestId"`
		Intent    struct {
			Name               string          `json:"name"`
			ConfirmationStatus string          `json:"confirmationstatus"`
			Slots              map[string]Slot `json:"slots"`
		} `json:"intent"`
		DialogState string `json:"dialogState"`
	} `json:"request"`
}

type Slot struct {
	Name               string      `json:"name"`
	Value              string      `json:"value,omitempty"`
	ConfirmationStatus string      `json:"confirmationStatus"`
	Resolutions        interface{} `json:"resolutions,omitempty"`
}

type AlexaResponse struct {
	Version  string `json:"version"`
	Response struct {
		OutputSpeech struct {
			Type string `json:"type"`
			Text string `json:"text,omitempty"`
			SSML string `json:"ssml,omitempty"`
		} `json:"outputSpeech,omitemtpty"`
		Directives       []interface{} `json:"directives,omitempty"`
		ShouldEndSession string        `json:"shouldEndSession"`
	} `json:"response"`
}

type Intent struct {
	Name               string                `json:"name"`
	ConfirmationStatus string                `json:"confirmationStatus,omitempty"`
	Slots              map[string]IntentSlot `json:"slots"`
}

type IntentSlot struct {
	Name               string `json:"name"`
	ConfirmationStatus string `json:"confirmationStatus,omitempty"`
	Value              string `json:"value"`
	ID                 string `json:"id,omitempty"`
}

type DialogDirective struct {
	Type          string  `json:"type"`
	SlotToElicit  string  `json:"slotToElicit,omitempty"`
	SlotToConfirm string  `json:"slotToConfirm,omitempty"`
	UpdatedIntent *Intent `json:"updatedIntent,omitempty"`
}

func CreateResponse(flag bool) *AlexaResponse {
	var resp AlexaResponse
	resp.Version = "1.0"
	resp.Response.ShouldEndSession = "true"
	if flag {

		resp.Response.OutputSpeech.Type = "PlainText"
		resp.Response.OutputSpeech.Text = "Hello.  Please override this default output."

	} else {
		resp.Response.OutputSpeech.Type = "SSML"

		resp.Response.OutputSpeech.SSML = "<speak> Hello, Please override this default SSML output. </speak>"
	}
	return &resp
}

//Say functions just output plaintext speech
func (resp *AlexaResponse) Say(text string) {
	resp.Response.OutputSpeech.Text = text
}

//AddDialogDirective adds a Dialog Directive to response
func (resp *AlexaResponse) AddDialogDirective(dialogType, slotToElicit, slotToConfirm string, intent *Intent) {
	d := DialogDirective{
		Type:          dialogType,
		SlotToElicit:  slotToElicit,
		SlotToConfirm: slotToConfirm,
		UpdatedIntent: intent,
	}
	resp.Response.Directives = append(r.Response.Directives, d)
}

// Ssay functions says something in SSML
func (resp *AlexaResponse) Ssay(text string) {
	var b bytes.Buffer
	b.WriteString("<speak>")
	b.WriteString(text)
	b.WriteString("</speak>")
	resp.Response.OutputSpeech.SSML = b.String()
}

//NSsay function repeats something N times
func (resp *AlexaResponse) NSsay(text string, number int) {
	var b bytes.Buffer
	b.WriteString("<speak>")
	for i := 0; i < number; i++ {
		b.WriteString("<p>")
		b.WriteString(text)
		b.WriteString("</p>")
	}
	b.WriteString("</speak>")
	resp.Response.OutputSpeech.SSML = b.String()
}

//HandleRequest function is the one which handles the request from alexa and gives response back
func HandleRequest(ctx context.Context, i AlexaRequest) (AlexaResponse, error) {
	// Use Spew to output the request for debugging purposes:
	fmt.Println("---- Dumping Input Map: ----")
	spew.Dump(i)
	fmt.Println("---- Done. ----")

	// Example of accessing map value via index:
	//log.Printf("Request type is %s\n ", i.Request.Intent.Name)
	//fmt.Println("Times is %s\n", i.Request.Intent.Slots.Times.Value)
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
		resp.Response.ShouldEndSession = "true"
		resp.Say("Hello there, Lambda appears to be working properly.")
	case "chew":
		resp = CreateResponse(false)

		numberOfTime, _ := strconv.Atoi(i.Request.Intent.Slots["times"].Value)
		resp.NSsay("Aarya Please <emphasis level='strong'> chew the food </emphasis> ", numberOfTime)
	case "AMAZON.HelpIntent":
		resp = CreateResponse(true)
		resp.Say("Helping aarya with some things")
	case "quiz":

		switch i.Request.DialogState {
		case "STARTED":

		case "COMPLETED":
			//	resp.Response.ShouldEndSession = "true"
			resp.Ssay("Completed")

		case "IN_PROGRESS":

		default:
			resp.Ssay("Some random default, it did not catch any of it")
		}
	default:
		resp = CreateResponse(true)
		resp.Say("I'm sorry, the input does not look like something I understand.")
	}

	return *resp, nil
}

func main() {
	lambda.Start(HandleRequest)
}
