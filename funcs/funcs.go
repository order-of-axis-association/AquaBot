package funcs

import (
	_ "fmt"
	"github.com/order-of-axis-association/AquaBot/types"
	"github.com/order-of-axis-association/AquaBot/utils"
	"github.com/bwmarrin/discordgo"
)

func Ping(cmd_args types.CmdArgs, s *discordgo.Session, m *discordgo.MessageCreate, global_state types.G_State) error {
	return utils.Say("Pong!", s, m)
}

func Pong(cmd_args types.CmdArgs, s *discordgo.Session, m *discordgo.MessageCreate, global_state types.G_State) error {
	return utils.Say("Ping!", s, m)
}

func Help(cmd_args types.CmdArgs, s *discordgo.Session, m *discordgo.MessageCreate, global_state types.G_State) error {
	return utils.Say("I'm a cute useless godess with a great ass. Leave me alone.", s, m)
}
