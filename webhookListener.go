package webhooks

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"reflect"
)

type WebhookContext struct {
	Delivery  string
	UserAgent string
	Event     string
	Signature string
	W         http.ResponseWriter
	R         *http.Request
	Body      []byte
}

type eventTypeDefinition struct {
	Method   string
	Template interface{}
}

type WebhookListener struct {
	OnCommitComment func(*CommitCommentEvent, *WebhookContext)
	OnCreate        func(*CreateEvent, *WebhookContext)
	OnDelete        func(*DeleteEvent, *WebhookContext)
	OnDeployment    func(*DeploymentEvent, *WebhookContext)
	OnIssueComment  func(*IssueCommentEvent, *WebhookContext)
	OnPing          func(*PingEvent, *WebhookContext)
	OnPullRequest   func(*PullRequestEvent, *WebhookContext)
	OnPush          func(*PushEvent, *WebhookContext)
}

// to avoid repetition, invoke handlers via reflection.
// event type header to Method name, and a prototype struct for the first argument
var eventTypes = map[string]eventTypeDefinition{
	"commit_comment": {"OnCommitComment", CommitCommentEvent{}},
	"create":         {"OnCreate", CreateEvent{}},
	"delete":         {"OnDelete", DeleteEvent{}},
	"deployment":     {"OnDeployment", DeploymentEvent{}},
	"issue_comment":  {"OnIssueComment", IssueCommentEvent{}},
	"ping":           {"OnPing", PingEvent{}},
	"pull_request":   {"OnPullRequest", PullRequestEvent{}},
	"push":           {"OnPush", PushEvent{}},
}

func (l *WebhookListener) GetHttpListener() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := &WebhookContext{W: w, R: r}
		ctx.Delivery = r.Header.Get("X-Github-Delivery")
		ctx.UserAgent = r.UserAgent()
		ctx.Event = r.Header.Get("X-Github-Event")
		ctx.Signature = r.Header.Get("X-Hub-Signature")
		body, _ := ioutil.ReadAll(r.Body)
		ctx.Body = body

		if !Validator(ctx) {
			w.WriteHeader(400)
			w.Write([]byte("Signature failure"))
			return
		}

		eventDef, ok := eventTypes[ctx.Event]
		if !ok {
			log.Printf("Unknown event type: %s", ctx.Event)
			w.WriteHeader(500)
			w.Write([]byte("Unknown Event type"))
			return
		}

		// find method to call
		listener := reflect.ValueOf(*l)
		field := listener.FieldByName(eventDef.Method)
		if !field.IsValid() {
			log.Printf("Invalid Field: %s", eventDef.Method)
			w.WriteHeader(500)
			return
		}
		if field.IsNil() {
			return
		}

		// make new object and deserialize payload
		typ := reflect.TypeOf(eventDef.Template)
		instanceValue := reflect.New(typ)
		instance := instanceValue.Interface()
		err := json.Unmarshal(body, instance)
		if err != nil {
			log.Printf("Error deserializing: %s", err)
			w.WriteHeader(500)
			return
		}

		// finally call the function.
		field.Call([]reflect.Value{instanceValue, reflect.ValueOf(ctx)})
	}
}

type ValidatorFunc func(*WebhookContext) bool

var Validator ValidatorFunc = func(*WebhookContext) bool {
	return true
}

const GithubHookSecretEnv = "github-hook-secret"

func init() {
	// If environment variable is set, replace no-op validator with one that verifies signature.
	if secret := os.Getenv(GithubHookSecretEnv); secret != "" {
		Validator = func(ctx *WebhookContext) bool {
			mac := hmac.New(sha1.New, []byte(secret))
			_, err := mac.Write(ctx.Body)
			if err != nil {
				return false
			}
			sig, err := hex.DecodeString(ctx.Delivery[5:])
			if err != nil {
				return false
			}
			return hmac.Equal(mac.Sum(nil), sig)
		}
	}
}
