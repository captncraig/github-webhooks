package webhooks

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

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
	}
}

func verifySha(signature string, body []byte) bool {
	secret := []byte("qwertyuiop")
	mac := hmac.New(sha1.New, secret)
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
