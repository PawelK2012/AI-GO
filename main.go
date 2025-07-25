package main

import (
	"PawelK2012/go-chat/src/repository"
	aichat "PawelK2012/go-chat/src/services/ai_chat"
	"fmt"
	"html/template"
	"net/http"
)

type UserInput struct {
	Message string
}

type LLMResp struct {
	Resp string
}

func main() {
	// Instantiate clients repository to open the OpenAI connection
	repository := repository.New()

	// comment out if you would like to try Judge service
	// judge := judge.New(repository)
	// judge.JudgeLLMResult()

	aichat := aichat.New(repository)

	tmpl := template.Must(template.ParseFiles("forms.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}

		user_input := UserInput{
			Message: r.FormValue("message"),
		}

		fmt.Printf("message = %s", user_input.Message)
		resp, _ := aichat.AskQuestion(user_input.Message)

		// LLM output is displayed in front end based on this struct
		llmRsp := LLMResp{
			Resp: "default resp",
		}
		llmRsp.Resp = resp
		tmpl.Execute(w, llmRsp)

	})

	http.ListenAndServe(":8080", nil)

}
