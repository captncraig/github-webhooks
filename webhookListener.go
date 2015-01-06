package webhooks

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
)

var secret = os.Getenv("github-hook-secret")

type WebhookContext struct {
	Delivery  string
	UserAgent string
	Event     string
	W         http.ResponseWriter
	R         *http.Request
}

type WebhookListener struct {
	OnCommitComment func(*CommitCommentEvent, *WebhookContext)
	OnPush          func(*PushEvent, *WebhookContext)
	OnAny           func(*PayloadBase, *WebhookContext)
}

func (l *WebhookListener) GetHttpListener() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := WebhookContext{W: w, R: r}
		ctx.Delivery = r.Header.Get("X-Github-Delivery")
		ctx.UserAgent = r.UserAgent()
		ctx.Event = r.Header.Get("X-Github-Event")
		signature := r.Header.Get("X-Hub-Signature")
		body, _ := ioutil.ReadAll(r.Body)
		if !verifySha(signature, body) {
			w.WriteHeader(400)
			w.Write([]byte("Signature failure"))
			return
		}
		switch ctx.Event {
		case "commit_comment":
			data := CommitCommentEvent{}
			json.Unmarshal(body, &data)
			if l.OnCommitComment != nil {
				l.OnCommitComment(&data, &ctx)
			}
		case "push":
			data := PushEvent{}
			json.Unmarshal(body, &data)
			if l.OnPush != nil {
				l.OnPush(&data, &ctx)
			}
		}
		if l.OnAny != nil {
			data := PayloadBase{}
			json.Unmarshal(body, &data)
			l.OnAny(&data, &ctx)
		}
	}
}

func verifySha(signature string, body []byte) bool {
	if secret == "" {
		return true
	}
	mac := hmac.New(sha1.New, []byte(secret))
	_, err := mac.Write(body)
	if err != nil {
		return false
	}
	sig, err := hex.DecodeString(signature[5:])
	if err != nil {
		return false
	}
	return hmac.Equal(mac.Sum(nil), sig)
}
