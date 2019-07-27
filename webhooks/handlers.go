package webhooks

import (
	"fmt"
	"strconv"

	"gopkg.in/go-playground/webhooks.v5/github"

	"github.com/bwmarrin/discordgo"
	"github.com/remiscarlet/shorturl/gitio"
)

// Guild ID =>
//   Channel ID =>
//     Repos to announce about
var AlertConfig = map[string]map[string][]string{
	"604536071891714069": map[string][]string{
		"604599071063408680": []string{
			"test-repo",
			"AquaBot",
		},
	},
	"602320820835975198": map[string][]string{
		"604621897078013955": []string{
			"AquaBot",
			"terraform",
			"OA_web",
		},
	},
}

func handlePushPayload(push github.PushPayload, dg *discordgo.Session) {
	sender_info := push.Sender
	sender_login := sender_info.Login

	commits := push.Commits
	num_commits := strconv.Itoa(len(commits)) + " commit"
	if len(commits) > 1 {
		num_commits = num_commits + "s"
	}

	repository := push.Repository
	full_repo_name := repository.FullName
	repo_name := repository.Name

	message := sender_login + " just pushed " + num_commits + " to *" + full_repo_name + "*\n"

	shortener := gitio.New()

	for _, commit := range commits {
		short_sha := commit.ID[:7]
		short_url_bytes, err := shortener.Shorten(commit.URL)
		short_url := string(short_url_bytes)

		if err != nil {
			fmt.Println("Error while shortening github commit url:", err)
		}

		// TODO: Bah. Use fmt.Sprintf
		message = message +
			"- `" + short_sha + "`" + ": " + "`" + commit.Message + "` (" + short_url + ")\n"
	}

	message = message + "\n Don't mind me. I'm just a useless goddess anyway. :("

	for _, channel_config := range AlertConfig {
		for channel_id, enabled_repos := range channel_config {

			valid_repo := false

			for _, enabled_repo := range enabled_repos {
				if enabled_repo == repo_name {
					valid_repo = true
				}
			}

			if valid_repo {
				dg.ChannelMessageSend(channel_id, message)
			}
		}
	}
}
