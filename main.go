package main

import (
	"PawelK2012/go-chat/src/repository"
	aichat "PawelK2012/go-chat/src/services/ai_chat"
	"fmt"
	"net/http"
	"text/template"
)

type ContactDetails struct {
	Email   string
	Subject string
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
	aichat.StartChat()

	tmpl := template.Must(template.ParseFiles("forms.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}

		details := ContactDetails{
			Email:   r.FormValue("email"),
			Subject: r.FormValue("subject"),
			Message: r.FormValue("message"),
		}

		// do something with details
		_ = details

		fmt.Printf("message = %s", details.Message)

		// tmpl.Execute(w, struct{ Success bool }{true})

		llmRsp := LLMResp{
			Resp: "default resp",
		}
		llmRsp.Resp = "LLM Msg" + details.Email
		tmpl.Execute(w, llmRsp)

	})

	http.ListenAndServe(":8080", nil)

}
