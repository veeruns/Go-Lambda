package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/davecgh/go-spew/spew"
)

func CapitalQuestion(country string) string {
	var b bytes.Buffer
	b.WriteString("What is the capital of ")
	b.WriteString(country)
	b.WriteString("?")
	return b.String()
}

func capitalquiz(resp *AlexaResponse, i AlexaRequest) *AlexaResponse {

	var quizanswer int
	resp = CreateResponse(false)
	var questionnumber int
	datanum := i.Session.Attributes
	switch i.Request.DialogState {
	case "STARTED":
		resp.Response.ShouldEndSession = Wrong
		questionnumber = 1
		resp.SessionAttributes = make(map[string]interface{})
		resp.SessionAttributes["questionnumber"] = strconv.Itoa(questionnumber)

		var QuestionToAsk string
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		max := 243
		min := 1
		randindex := r1.Intn(max-min) + min
		cinfo, _ := getItemIdx(randindex)
		resp.SessionAttributes["PreviousAnswer"] = cinfo.City
		QuestionToAsk = CapitalQuestion(cinfo.Country)
		resp.Ssay(QuestionToAsk)

		var intent string
		var b2 bytes.Buffer
		b2.WriteString(`{
        "name": "capitalquiz",
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
		var correctanswers int
		var previousanswer string
		resp.Response.ShouldEndSession = Correct
		for k, v := range datanum {
			switch val := v.(type) {
			case string:
				if k == "CorrectAnswers" {
					correctanswers, _ = strconv.Atoi(val)
				} else if k == "PreviousAnswer" {
					previousanswer = val
				}
			default:
				fmt.Printf("There is default case")
			}
		}
		quizanswer := i.Request.Intent.Slots["Answer"].Value
		if strings.Compare(quizanswer, previousanswer) == 0 {
			correctanswers++
		}
		var builder bytes.Buffer
		builder.WriteString("<p>Aarya, Thank you for playing quiz game</p>")
		builder.WriteString("<p> You have answered ")
		builder.WriteString(strconv.Itoa(correctanswers))
		builder.WriteString(" questions correctly</p>")
		if correctanswers > 4 {
			builder.WriteString("<p> Very good job </p>")
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
					//	fmt.Printf("Did you get questionnumber %d %s\n", questionnumber, k)
				} else if k == "PreviousAnswer" {
					previousanswer, _ = val
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

		if strings.Compare(quizanswer, previousanswer) == 0 {
			correctanswers++
		} else {
			ResponseAlexa.WriteString("<p>That is not the correct Answer, The correct answer is ")
			ResponseAlexa.WriteString(strconv.Itoa(previousanswer))
			ResponseAlexa.WriteString("</p>")
		}
		resp.SessionAttributes["CorrectAnswers"] = strconv.Itoa(correctanswers)
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		max := 243
		min := 1
		randindex := r1.Intn(max-min) + min
		cinfo, _ := getItemIdx(randindex)
		resp.SessionAttributes["PreviousAnswer"] = cinfo.City
		QuestionToAsk = CapitalQuestion(cinfo.Country)
		//	resp.SessionAttributes["PreviousAnswer"] = strconv.Itoa(m1 * m2)
		ResponseAlexa.WriteString("<p> Next Question </p><p>")
		ResponseAlexa.WriteString(QuestionToAsk)
		ResponseAlexa.WriteString("</p>")

		if questionnumber < 6 {
			resp.Ssay(ResponseAlexa.String())
			resp.Response.ShouldEndSession = Wrong
			var intent string
			var b2 bytes.Buffer
			b2.WriteString(`{
          "name": "capitalquiz",
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

	return resp
}
