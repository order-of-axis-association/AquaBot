package config_funcs

import (
	"fmt"
	"strings"
	"time"

	"github.com/order-of-axis-association/AquaBot/db"
	"github.com/order-of-axis-association/AquaBot/types"
	"github.com/order-of-axis-association/AquaBot/utils"
)

var ConfigCmd = types.Command {
	Cmd:     "config",
	Version: "0.0.1",

	Func:    ConfigFunc,
	Flags:   nil,
	Usage:   ConfigUsage,
	Configs: ConfigConfigs, // *barf*
}

var ConfigUsage string = `
config add <key> <val>
	- Help text
config rem <key>
	- Help text
`

var ConfigConfigs = types.ConfigFlags{
	"configs.allowed_users": validate_allowedusers,
	"configs.blocked_users": nil,
}

func ConfigFunc(cmd_args types.CmdArgs, state types.MessageState) error {
	pos_args := cmd_args.PosArgs
	if len(pos_args) == 0 {
		return printErrorAndUsage("You must provide a valid subcommand!", state)
	}
	subcommand := strings.ToLower(cmd_args.PosArgs[0])

	if utils.StrContains(subcommand, []string{"add"}) {
		return addConfig(cmd_args, state)
	} else if utils.StrContains(subcommand, []string{"remove"}) {
		return remConfig(cmd_args, state)
	} else {
		return printErrorAndUsage("Invalid subcommand!", state)
	}

	return nil
}

func addConfig(cmd_args types.CmdArgs, state types.MessageState) error {
	pos_args := cmd_args.PosArgs
	if len(pos_args) < 2 {
		return printErrorAndUsage("Must provide a config", state)
	}

	config_key := pos_args[1]
	valid_config_keys := getValidConfigKeys()

	fmt.Printf("Valid conf keys: %+v\n", valid_config_keys)

	if utils.StrContains(config_key, valid_config_keys) {
		return utils.TempSay("Asdf", state)
	}

	return nil
}

func remConfig(cmd_args types.CmdArgs, state types.MessageState) error {
	return nil
}

func add(key string, val string, state types.MessageState) error {
	conf := db.Config{
		GuildId:       &state.M.GuildID,
		ChannelId:     &state.M.ChannelID,
		LastUserToSet: state.M.Author.String(),
		LastUpdated:   time.Now(),

		ConfigName:  key,
		ConfigValue: val,
	}

	fmt.Printf("%+v\n", conf)

	//db.G.DBConn.Create(

	return nil
}

func get(key string, state types.MessageState) (string, error) {
	return "", nil
}
