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

	listener.OnPing = func(e *webhooks.PingEvent, ctx *webhooks.WebhookContext) {

		fmt.Println(e.HookId, e.Zen)
	}

	http.HandleFunc("/hook", listener.GetHttpListener())
	http.ListenAndServe(":8787", nil)
}
