package main

import (
	"fmt"
	"github.com/captncraig/github-webhooks"
	"net/http"
)

func main() {
	listener := webhooks.WebhookListener{}

	listener.OnCommitComment = func(e *webhooks.CommitCommentEvent, ctx *webhooks.WebhookContext) {
		fmt.Println(e.Comment.User.Login, e.Comment.Body)
	}

	listener.OnPush = func(e *webhooks.PushEvent, ctx *webhooks.WebhookContext) {
		fmt.Println(e)
	}

	http.HandleFunc("/qwertyuiop", listener.GetHttpListener())
	http.ListenAndServe(":12345", nil)
}
