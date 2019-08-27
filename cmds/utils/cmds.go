package util_cmds

import (
	_ "fmt"
	"github.com/order-of-axis-association/AquaBot/types"
	"github.com/order-of-axis-association/AquaBot/utils"
)

var Ping = types.Command{
	Cmd:     "ping",
	Version: "1.0.0",

	Func:  PingFunc,
	Flags: nil,
	Usage: "",
}

func PingFunc(cmd_args types.CmdArgs, state types.MessageState) error {
	return utils.TempSay("Pong!", state)
}

var Pong = types.Command{
	Cmd:     "pong",
	Version: "1.0.0",

	Func:  PongFunc,
	Flags: nil,
	Usage: "",
}

func PongFunc(cmd_args types.CmdArgs, state types.MessageState) error {
	return utils.TempSay("Ping!", state)
}
