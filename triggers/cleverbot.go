package triggers

import (
	_ "fmt"
	"regexp"

	"github.com/order-of-axis-association/AquaBot/types"
)

func InvokeCleverbot(message string, state types.MessageState, payload_chan chan types.CBPayload) {
	replacements := map[string]string {
		`<@603252075006001152>`: "Cleverbot",
		`(?i)\sAqua\s`: "Cleverbot",
	}

	for regex, rep := range replacements {
		re := regexp.MustCompile(regex)
		message = re.ReplaceAllString(message, rep)
	}

	payload := types.CBPayload {
		Msg: message,
		MsgState: &state,
	}

	payload_chan <- payload
}
