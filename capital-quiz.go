package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/davecgh/go-spew/spew"
)

func capitalquiz() {

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
}
