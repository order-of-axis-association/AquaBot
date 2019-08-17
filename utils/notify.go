package utils

// Functions to "notify" people with.

import (
	"fmt"

	"github.com/order-of-axis-association/AquaBot/db"
	"github.com/order-of-axis-association/AquaBot/types"
)

func Say(message string, state types.MessageState) error {
	c, err := state.S.State.Channel(state.M.ChannelID)
	if err != nil {
		return err
	}

	_, err = state.S.ChannelMessageSend(c.ID, message)
	if err != nil {
		fmt.Println("Error sending message:", err)
		return err
	}

	return nil
}

func TempSay(message string, state types.MessageState) error {
	c, err := state.S.State.Channel(state.M.ChannelID)
	if err != nil {
		return err
	}

	msg, err := state.S.ChannelMessageSend(c.ID, message)
	if err != nil {
		fmt.Println("Error sending message:", err)
		return err
	}

	tempmessage := db.TempMessage{
		MessageId: msg.ID,
		ChannelId: msg.ChannelID,
		Length: 15,
	}
	if err := state.G.DBConn.Create(&tempmessage).Error; err != nil {
		return err
	}

	return nil
}

// Sends the message as "monospace", ie codeblocks
func Mono(message string, state types.MessageState) error {
	message = fmt.Sprintf("```%s```", message)
	return TempSay(message, state)
}

// Sends message with a nice unicode X prepended at the front of message
func Error(message string, state types.MessageState) error {
	message = fmt.Sprintf("❌ *%s*", message)
	ApplyErrorReaction(state)
	return TempSay(message, state)
}
