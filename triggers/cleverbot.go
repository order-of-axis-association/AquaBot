package triggers

import (
	_ "fmt"

	"github.com/order-of-axis-association/AquaBot/types"
)

func InvokeCleverbot(message string, state types.MessageState, payload_chan chan types.CBPayload) {
	payload := types.CBPayload{
		Msg:      message,
		MsgState: &state,
	}

	payload_chan <- payload
}
