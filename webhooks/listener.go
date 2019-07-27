package webhooks

import (
	"fmt"
	"os"

	"net/http"

	"github.com/bwmarrin/discordgo"

	"gopkg.in/go-playground/webhooks.v5/github"
)

func InitWebhookServer(dg *discordgo.Session) {
	webhook_secret := os.Getenv("WEBHOOK_SERVER_SECRET")
	fmt.Println(webhook_secret)
	hook, _ := github.New(github.Options.Secret(webhook_secret))

	const (
		path = "/"
	)

	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		payload, err := hook.Parse(r, github.PushEvent)
		if err != nil {
			if err == github.ErrEventNotFound {
				// ok event wasn;t one of the ones asked to be parsed
			}
		}

		switch payload.(type) {
		case github.PushPayload:
			push := payload.(github.PushPayload)
			handlePushPayload(push, dg)
		}
	})
	fmt.Println("Starting server...")
	http.ListenAndServe(":25100", nil)

	fmt.Println("Listening on port :25100")
}
