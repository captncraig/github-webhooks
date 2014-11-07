package main

import (
	"encoding/json"
	"fmt"
	"github.com/captncraig/github-webhooks"
	"net/http"
)

func main() {
	listener := webhooks.WebhookListener{}
	listener.OnCommitComment = func(c *webhooks.CommitCommentEvent) {
		b, _ := json.MarshalIndent(c, "", "  ")
		fmt.Println(string(b))
	}
	http.HandleFunc("/qwertyuiop", listener.GetHttpListener())
	http.ListenAndServe(":12345", nil)
}
