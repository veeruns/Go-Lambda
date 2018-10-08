package alexalib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"math/rand"
	"reflect"
)

//Correct is the new name for true
const Correct = "true"

//Wrong is the  new name for false
const Wrong = "false"

//AlexaRequest Structure
type AlexaRequest struct {
	Version string   `json:"version"`
	Session *Session `json:"session"`
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

//Session structure
type Session struct {
	New        bool                   `json:"new"`
	SessionID  string                 `json:"sessionId"`
	Attributes map[string]interface{} `json:"attributes"`
	User       struct {
		UserID      string `json:"userId"`
		AccessToken string `json:"accessToken"`
	} `json:"user"`
	Application struct {
		ApplicationID string `json:"applicationId"`
	} `json:"application"`
}

//Slot Structure
type Slot struct {
	Name               string      `json:"name"`
	Value              string      `json:"value,omitempty"`
	ConfirmationStatus string      `json:"confirmationStatus"`
	Resolutions        interface{} `json:"resolutions,omitempty"`
}

//OutputSpeech structure
type OutputSpeech struct {
	Type string `json:"type,omitempty"`
	Text string `json:"text,omitempty"`
	SSML string `json:"ssml,omitempty"`
}

//Response structure
type Response struct {
	OutputSpeech     *OutputSpeech `json:"outputSpeech,omitempty"`
	Directives       []interface{} `json:"directives,omitempty"`
	ShouldEndSession string        `json:"shouldEndSession,omitempty"`
}

//AlexaResponse Structure
type AlexaResponse struct {
	Version           string                 `json:"version"`
	SessionAttributes map[string]interface{} `json:"sessionAttributes,omitempty"`
	Response          Response               `json:"response,omitempty"`
}

//Intent Structure
type Intent struct {
	Name               string                `json:"name"`
	ConfirmationStatus string                `json:"confirmationStatus,omitempty"`
	Slots              map[string]IntentSlot `json:"slots"`
}

//IntentSlot is the Alexa IntentSlots
type IntentSlot struct {
	Name               string `json:"name"`
	ConfirmationStatus string `json:"confirmationStatus,omitempty"`
	Value              string `json:"value"`
	ID                 string `json:"id,omitempty"`
}

//DialogDirective is the structure that has DialogDirective
type DialogDirective struct {
	Type          string  `json:"type"`
	SlotToElicit  string  `json:"slotToElicit,omitempty"`
	SlotToConfirm string  `json:"slotToConfirm,omitempty"`
	UpdatedIntent *Intent `json:"updatedIntent,omitempty"`
}

//CreateResponse with flag to create either an SSML or plaintext outputSpeech
func CreateResponse(flag bool) *AlexaResponse {
	var resp AlexaResponse
	resp.Version = "1.0"
	var speech OutputSpeech

	resp.Response.ShouldEndSession = Correct
	if flag {
		speech = OutputSpeech{
			Type: "PlainText",
			Text: "Please over ride",
		}

		resp.Response.OutputSpeech = &speech

	} else {

		speech = OutputSpeech{
			Type: "SSML",
			SSML: "<speak>Please over ride </speak>",
		}
		resp.Response.OutputSpeech = &speech
	}
	return &resp
}

//Say functions just output plaintext speech
func (resp *AlexaResponse) Say(text string) {
	var speech OutputSpeech
	speech = OutputSpeech{
		Type: "PlainText",
		Text: text,
	}
	sm, _ := json.Marshal(resp)
	fmt.Printf("In Say the response is %s\n", sm)
	resp.Response.OutputSpeech = &speech
}

//EndResponse function clears everything
func (resp *AlexaResponse) EndResponse() {
	clear(resp)
	resp.Version = "1.0"
	resp.Response.ShouldEndSession = Wrong
	var dtype string
	dtype = "Dialog.Delegate"
	d := DialogDirective{
		Type: dtype,
	}
	resp.Response.Directives = append(resp.Response.Directives, d)
}

//AddDialogDirective adds a Dialog Directive to response
func (resp *AlexaResponse) AddDialogDirective(dialogType, slotToElicit, slotToConfirm string, intent *Intent) {
	d := DialogDirective{
		Type:          dialogType,
		SlotToElicit:  slotToElicit,
		SlotToConfirm: slotToConfirm,
		UpdatedIntent: intent,
	}
	resp.Response.Directives = append(resp.Response.Directives, d)
}

// Ssay functions says something in SSML
func (resp *AlexaResponse) Ssay(text string) {
	var b bytes.Buffer
	b.WriteString("<speak>")
	b.WriteString(text)
	b.WriteString("</speak>")
	var speech OutputSpeech
	var op string
	op = b.String()
	speech = OutputSpeech{
		Type: "SSML",
		SSML: string(op),
	}

	resp.Response.OutputSpeech = &speech
}

func clear(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
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

	var op string
	op = b.String()

	resp.Response.OutputSpeech = &OutputSpeech{
		Type: "SSML",
		SSML: op,
	}
}

//CreatePairs creates a pair of multiplier and mutliplicand less than 16
func CreatePairs() (int, int) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	max := 12
	min := 1
	multiplier := r1.Intn(max-min) + min
	multiplicant := r1.Intn(max-min) + min
	return multiplier, multiplicant
}

//CreateQuestion functions
func CreateQuestion(multiplicand, multiplier int) string {
	var b bytes.Buffer
	b.WriteString("What is the answer to ")
	b.WriteString(strconv.Itoa(multiplicand))
	b.WriteString(" multiplied by ")
	b.WriteString(strconv.Itoa(multiplier))
	b.WriteString(" ")
	return b.String()
}
