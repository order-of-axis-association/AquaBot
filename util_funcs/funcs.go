package util_funcs

import (
	_ "fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/order-of-axis-association/AquaBot/types"
	"github.com/order-of-axis-association/AquaBot/utils"
)

var Ping = types.Command {
	Cmd: "ping",
	Version: "1.0.0",

	Func: PingFunc,
	Flags: nil,
	Usage: "",
}

func PingFunc(cmd_args types.CmdArgs, s *discordgo.Session, m *discordgo.MessageCreate, global_state types.G_State) error {
	return utils.Say("Pong!", s, m)
}

var Pong = types.Command {
	Cmd: "pong",
	Version: "1.0.0",

	Func: PongFunc,
	Flags: nil,
	Usage: "",
}

func PongFunc(cmd_args types.CmdArgs, s *discordgo.Session, m *discordgo.MessageCreate, global_state types.G_State) error {
	return utils.Say("Ping!", s, m)
}
