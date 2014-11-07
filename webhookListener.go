package webhooks

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type WebhookListener struct {
	OnCommitComment func(*CommitCommentEvent)
}

func (l *WebhookListener) GetHttpListener() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		delivery := r.Header.Get("X-Github-Delivery")
		userAgent := r.UserAgent()
		event := r.Header.Get("X-Github-Event")
		signature := r.Header.Get("X-Hub-Signature")
		fmt.Println(delivery, userAgent, event, signature)
		body, _ := ioutil.ReadAll(r.Body)
		switch event {
		case "commit_comment":
			data := CommitCommentEvent{}
			json.Unmarshal(body, &data)
			if l.OnCommitComment != nil {
				l.OnCommitComment(&data)
			}
		}
	}
}
