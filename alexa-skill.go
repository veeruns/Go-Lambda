package main

import (
	"bytes"
	"context"
	"encoding/json"
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
	Value              string      `json:"value"`
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

func (resp *AlexaResponse) Say(text string) {
	resp.Response.OutputSpeech.Text = text
}

func (r *AlexaResponse) AddDialogDirective(dialogType, slotToElicit, slotToConfirm string, intent *Intent) {
	d := DialogDirective{
		Type:          dialogType,
		SlotToElicit:  slotToElicit,
		SlotToConfirm: slotToConfirm,
		UpdatedIntent: intent,
	}
	r.Response.Directives = append(r.Response.Directives, d)
}

func (resp *AlexaResponse) Ssay(text string) {
	var b bytes.Buffer
	b.WriteString("<speak>")
	b.WriteString(text)
	b.WriteString("</speak>")
	resp.Response.OutputSpeech.SSML = b.String()
}

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
		var number1, number2 int
		number1 = 5
		number2 = 6
		//	var answerCalc int
		//		answerCalc = number1 * number2
		resp = CreateResponse(false)
		switch i.Request.DialogState {
		case "STARTED":
			var b bytes.Buffer
			b.WriteString("What is the answer to ")
			b.WriteString(strconv.Itoa(number1))
			b.WriteString(" multiplied by ")
			b.WriteString(strconv.Itoa(number2))
			b.WriteString(" ")
			resp.Ssay(b.String())
			resp.Response.ShouldEndSession = "false"
			var intent string
			var b2 bytes.Buffer
			b2.WriteString(`{

	"name": "quiz",
	"confirmationStatus": "NONE",
	"slots": {
		"Answer": {
			"name": "Answer",
			"confirmationStatus": "NONE"
		},
		"Question": {
			"name": "Question",
			"confirmationStatus": "NONE",
			"value": "`)
			b2.WriteString(strconv.Itoa(number1))
			b2.WriteString(" multiplied by ")
			b2.WriteString(strconv.Itoa(number2))
			b2.WriteString(`"}
	}

}`)

			intent = b2.String()

			fmt.Printf("Intent is %s\n", intent)
			updatedintent := Intent{}
			json.Unmarshal([]byte(intent), &updatedintent)

			fmt.Printf("%v\n", updatedintent)
			resp.AddDialogDirective("Dialog.ElicitSlot", "Question", "", &updatedintent)
		case "COMPLETED":
			//	resp.Response.ShouldEndSession = "true"
			resp.Ssay("Completed")
			//resp.AddDialogDirective(dialogType, slotToElicit, slotToConfirm, intent)
		case "IN_PROGRESS":
			resp.Response.ShouldEndSession = "false"
			/*   			var intent string
						var b2 bytes.Buffer
						b2.WriteString(`{

				"name": "quiz",
				"confirmationStatus": "NONE",
				"slots": {
					"Answer": {
						"name": "Answer",
						"confirmationStatus": "NONE"
						"Value": "Done"
					},`)
						b2.WriteString(`"Question": {
						"name": "Question",
						"confirmationStatus": "NONE",
						"value":"`)
						b2.WriteString(strconv.Itoa(number1))
						b2.WriteString(" multiplied by ")
						b2.WriteString(strconv.Itoa(number2))
						b2.WriteString(`"}
				}

			}`)

						intent = b2.String()
						updatedintent := Intent{}
						json.Unmarshal([]byte(intent), &updatedintent)
						given_answer, _ := strconv.Atoi(i.Request.Intent.Slots["Answer"].Value)
						if answerCalc == given_answer {
							resp.Ssay("Its correct")
						} else {
							resp.Ssay("Sorry its wrong")
						}
						resp.AddDialogDirective("Dialog.ElicitSlot", "Question", "", &updatedintent)*/
			resp.Ssay("In progress worked")
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
