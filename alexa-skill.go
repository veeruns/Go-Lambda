package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"math/rand"
	"reflect"

	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/davecgh/go-spew/spew"
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
	Name               string       `json:"name"`
	Value              string       `json:"value,omitempty"`
	ConfirmationStatus string       `json:"confirmationStatus"`
	Resolutions        *Resolutions `json:"resolutions,omitempty"`
}

//Resolutions structure
type Resolutions struct {
	ResolutionsPerAuthority []struct {
		Authority string `json:"authority"`
		Status    struct {
			Code string `json:"code"`
		} `json:"status"`
		Values []struct {
			Value struct {
				Name string `json:"name"`
				ID   string `json:"id"`
			} `json:"value"`
		} `json:"values"`
	} `json:"resolutionsPerAuthority"`
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

//ResolutionValue finds the first value of Resolutions
func ResolutionValue(input interface{}) string {
	var returnvalue string
	switch input.(type) {
	case Slot:
		fmt.Printf("IT is a slot type hence")
		fmt.Printf("Slot value is %s\n", input.(Resolutions).ResolutionsPerAuthority[0].Values[0].Value.Name)
	case *Resolutions:
		fmt.Printf("It is a pointer to resolution")
		var slotvalue *Resolutions
		slotvalue = input.(*Resolutions)
		returnvalue = slotvalue.ResolutionsPerAuthority[0].Values[0].Value.Name
		fmt.Printf("Slot value is %s\n", slotvalue.ResolutionsPerAuthority[0].Values[0].Value.Name)
	case Resolutions:
		fmt.Printf("It is resolution itself")
	default:
		fmt.Println("Do not care ", reflect.TypeOf(input).String())
	}

	return returnvalue
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

//createurl for channels will encode the url the right way.
func createurl(callingfunc string, para string) string {
	var Url *url.URL
	Url, err := url.Parse("https://veeruns.raghavanonline.com/roku")
	if err != nil {
		panic("boom")
	}

	parameters := url.Values{}
	parameters.Add("channel", para)
	Url.RawQuery = parameters.Encode()

	return Url.String()
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

//CallEndPoint functions calls "The" endpoint
func CallEndPoint(api string) string {
	if api == "" {
		api = "PowerOff"
	}
	cert, err := tls.LoadX509KeyPair("ssl/certs/cert.pem", "ssl/keys/key.pem")
	if err != nil {
		fmt.Printf("It did not work %s\n", err.Error())
	}
	caCert, err := ioutil.ReadFile("ssl/certs/cacert.pem")
	if err != nil {
		fmt.Printf("Could not load CA cert %s\n", err.Error())
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// Setup HTTPS client
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      caCertPool,
	}
	tlsConfig.BuildNameToCertificate()
	transport := &http.Transport{TLSClientConfig: tlsConfig}
	client := &http.Client{Transport: transport}
	var url bytes.Buffer
	url.WriteString("https://veeruns.raghavanonline.com:8080/roku?")
	url.WriteString(api)
	resp, err := client.Get(url.String())
	if err != nil {
		fmt.Printf("HTTP failed %s\n", err.Error())
		return "OhOh"
	}
	if resp.StatusCode != 200 {
		return "OhOh Wrong Status code"
	} else {
		var b bytes.Buffer
		b.WriteString("We are good ")
		b.WriteString(strconv.Itoa(resp.StatusCode))
		b.WriteString(" That is the resp code")
		fmt.Printf("We do not have an issue %d\n", resp.StatusCode)
		return b.String()
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
		/* Need to clean up quiz, general way dialog works */

	case "rokucontrol":
		resp = CreateResponse(false)
		functocall := i.Request.Intent.Slots["func"].Value
		fmt.Printf("The request intent value is %s\n", functocall)
		var b bytes.Buffer
		b.WriteString("func=")
		b.WriteString(functocall)
		say := CallEndPoint(b.String())
		fmt.Printf("Output from Endpoint is %s\n", say)
		resp.Ssay("The request has been send to the device")
	case "changechannel":
		resp = CreateResponse(false)
		channelname := i.Request.Intent.Slots["Channel"].Value
		var resolutions string
		resolutions = ResolutionValue(i.Request.Intent.Slots["Channel"].Resolutions)
		fmt.Printf("Resolution is %s\n", resolutions)

		//		fmt.Printf("Slot value is %s\n", slotvalue)
		fmt.Printf("The channel name is %s\n", channelname)
		var b string
		b = createurl("roku", resolutions)
		fmt.Printf("Created url is %s\n", b)
		op := CallEndPoint(b)
		fmt.Printf("Output from launch channel endpoint is %s\n", op)
		resp.Ssay("Request send to the device")
	case "quiz":
		var quizanswer int
		resp = CreateResponse(false)
		var questionnumber int
		datanum := i.Session.Attributes
		fmt.Println("DATANNUM OP")
		spew.Dump(datanum)
		fmt.Println("DATANUM OP DONE")
		switch i.Request.DialogState {
		case "STARTED":
			resp.Response.ShouldEndSession = Wrong
			questionnumber = 1
			resp.SessionAttributes = make(map[string]interface{})
			resp.SessionAttributes["questionnumber"] = strconv.Itoa(questionnumber)

			var QuestionToAsk string
			multiplier, multiplicant := CreatePairs()
			resp.SessionAttributes["PreviousAnswer"] = strconv.Itoa(multiplier * multiplicant)
			QuestionToAsk = CreateQuestion(multiplier, multiplicant)
			resp.Ssay(QuestionToAsk)

			var intent string
			var b2 bytes.Buffer
			b2.WriteString(`{
					"name": "quiz",
						"confirmationStatus": "NONE",
							"slots": {
									"Answer": {
											"name": "Answer",
												"confirmationStatus": "NONE"
											}
										}
										}`)

			intent = b2.String()
			updatedintent := Intent{}
			json.Unmarshal([]byte(intent), &updatedintent)
			resp.AddDialogDirective("Dialog.ElicitSlot", "Answer", "", &updatedintent)
			jsonresp, _ := json.Marshal(resp)
			fmt.Printf("When dialog started the resp is %s\n", jsonresp)

		case "COMPLETED":
			var previousanswer, correctanswers int

			resp.Response.ShouldEndSession = "true"
			for k, v := range datanum {
				switch val := v.(type) {
				case string:
					if k == "CorrectAnswers" {
						correctanswers, _ = strconv.Atoi(val)
					} else if k == "PreviousAnswer" {
						previousanswer, _ = strconv.Atoi(val)
					}
				default:
					fmt.Printf("There is default case")
				}
			}
			quizanswer, _ = strconv.Atoi(i.Request.Intent.Slots["Answer"].Value)
			if quizanswer == previousanswer {
				correctanswers++
			}
			var builder bytes.Buffer
			builder.WriteString("<p>Aarya, Thank you for playing quiz game</p>")
			builder.WriteString("<p> You have answered ")
			builder.WriteString(strconv.Itoa(correctanswers))
			builder.WriteString(" questions correctly</p>")
			if correctanswers > 4 {
				builder.WriteString("<audio src='soundbank://soundlibrary/human/amzn_sfx_crowd_applause_05'/>")
			}
			resp.Ssay(builder.String())
		case "IN_PROGRESS":
			datanum := i.Session.Attributes
			var previousanswer, correctanswers int
			fmt.Println("DATANNUM OP")
			spew.Dump(datanum)
			fmt.Println("DATANUM OP DONE")
			for k, v := range datanum {
				switch val := v.(type) {
				case string:
					if k == "questionnumber" {
						questionnumber, _ = strconv.Atoi(val)
						fmt.Printf("Did you get questionnumber %d %s\n", questionnumber, k)
					} else if k == "PreviousAnswer" {
						previousanswer, _ = strconv.Atoi(val)
					} else if k == "CorrectAnswers" {
						correctanswers, _ = strconv.Atoi(val)
					}
				default:
					fmt.Printf("There is default case")
				}

			}
			qanswer, _ := strconv.Atoi(i.Request.Intent.Slots["Answer"].Value)
			questionnumber++
			resp.SessionAttributes = make(map[string]interface{})

			resp.SessionAttributes["questionnumber"] = strconv.Itoa(questionnumber)
			var ResponseAlexa bytes.Buffer

			if qanswer == previousanswer {
				ResponseAlexa.WriteString("<p>That is the correct Answer</p>")
				correctanswers++

			} else {
				ResponseAlexa.WriteString("<p>That is not the correct Answer, The correct answer is ")
				ResponseAlexa.WriteString(strconv.Itoa(previousanswer))
				ResponseAlexa.WriteString("</p>")
			}
			resp.SessionAttributes["CorrectAnswers"] = strconv.Itoa(correctanswers)
			m1, m2 := CreatePairs()
			qtoa := CreateQuestion(m1, m2)
			resp.SessionAttributes["PreviousAnswer"] = strconv.Itoa(m1 * m2)
			ResponseAlexa.WriteString("<p> Next Question </p><p>")
			ResponseAlexa.WriteString(qtoa)
			ResponseAlexa.WriteString("</p>")

			if questionnumber < 6 {
				resp.Ssay(ResponseAlexa.String())
				resp.Response.ShouldEndSession = Wrong
				var intent string
				var b2 bytes.Buffer
				b2.WriteString(`{
						"name": "quiz",
							"confirmationStatus": "NONE",
								"slots": {
										"Answer": {
												"name": "Answer",
													"confirmationStatus": "NONE"
												}
											}
											}`)

				intent = b2.String()
				updatedintent := Intent{}
				json.Unmarshal([]byte(intent), &updatedintent)
				resp.AddDialogDirective("Dialog.ElicitSlot", "Answer", "", &updatedintent)
			} else {

				var intent string
				intent = `{
	"version": "1.0",
	"response": {
		"directives": [{
			"type": "Dialog.Delegate"
		}],
		"shouldEndSession": "False"
	}
}`
				//updatedintent := Al{}
				clear(resp)
				json.Unmarshal([]byte(intent), resp)
			}
			//resp.AddDialogDirective("Dialog.ElicitSlot", "Answer", "", )
		default:
			resp.Ssay("Some random default, it did not catch any of it")
		}
	case "capitals":
		resp = CreateResponse(false)
		countryname := i.Request.Intent.Slots["Question"].Value
		capitalname, err := getItem(countryname)
		var b bytes.Buffer
		if err != nil {
			b.WriteString(err.Error())
		} else {
			b.WriteString("Capital of ")
			b.WriteString(countryname)
			b.WriteString(" is ")
			b.WriteString(capitalname.City)
		}
		resp.Ssay(b.String())
	case "capitalquiz":
		resp = CreateResponse(true)
		resp = capitalquiz(resp, i)
	default:
		resp = CreateResponse(true)
		resp.Say("I'm sorry, the input does not look like something I understand.")
	}

	return *resp, nil
}

func main() {
	lambda.Start(HandleRequest)
}
