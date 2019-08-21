package cb

import (
	"fmt"
	"time"
	"strings"
	"regexp"
	"io/ioutil"

	"github.com/order-of-axis-association/AquaBot/types"
	"github.com/order-of-axis-association/AquaBot/utils"
	"github.com/ugjka/cleverbot-go"
)

var SESSION_LENGTH int64 = 60 * 10 // 10 minutes
var CLEVERBOT_API_FILE_PATH = "secrets/cleverbot_api_key"

type SessionID struct {
	Session				*cleverbot.Session
	LastInteraction		int64
}

func (s SessionID) IsStale() bool {
	return time.Now().After(time.Unix(s.LastInteraction + SESSION_LENGTH, 0))
}

func getToken() string {
	data, err := ioutil.ReadFile(CLEVERBOT_API_FILE_PATH)
	if err != nil {
		fmt.Println("Could not read Cleverbot API file! Err:", err)
	}

	return strings.TrimSpace(string(data))
}

var QUESTION_REPLACEMENTS map[string]string = map[string]string {
	`<@603252075006001152>`: "Cleverbot",
	`(?i)\sAqua\s`: "Cleverbot",
}

var RESPONSE_REPLACEMENTS map[string]string = map[string]string {
	`(?i)cleverbot`: "Aqua",
}

func replaceAll(msg string, replacement_map map[string]string) string {
	for regex, rep := range replacement_map {
		re := regexp.MustCompile(regex)
		msg = re.ReplaceAllString(msg, rep)
	}

	return msg
}

func StartCBDaemon (g_state types.G_State, msg_chan chan types.CBPayload) {
	api_token := getToken()
	session_pool := make(map[string]SessionID, 0)

	var payload types.CBPayload
	payload = <-msg_chan

	for ; true; payload = <-msg_chan {
		msg := payload.Msg
		msg = replaceAll(msg, QUESTION_REPLACEMENTS)

		msg_state := payload.MsgState

		username := msg_state.M.Author.Username + "#" + msg_state.M.Author.Discriminator

		session, exists := session_pool[username]
		if ! exists || session.IsStale() {
			session := SessionID{
				Session: cleverbot.New(api_token),
				LastInteraction: time.Now().Unix(),
			}
			session_pool[username] = session
		}
		session, _ = session_pool[username]

		resp, err := session.Session.Ask(msg)
		if err != nil {
			fmt.Println("Failed to ask question. Err:", err)
			return
		}
		resp = replaceAll(resp, RESPONSE_REPLACEMENTS)

		mention := fmt.Sprintf("<@%s> ", msg_state.M.Author.ID)
		err = utils.Say(mention + resp, *msg_state)
		if err != nil {
			fmt.Println("Error saying cleverbot response. Err:", err)
		}
	}
}
