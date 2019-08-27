package config_cmds

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
config set <key> <val>
	- Help text
config unset <key>
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

	if utils.StrContains(subcommand, []string{"set"}) {
		return setConfig(cmd_args, state)
	} else if utils.StrContains(subcommand, []string{"unset"}) {
		return unsetConfig(cmd_args, state)
	} else {
		return printErrorAndUsage("Invalid subcommand!", state)
	}

	return nil
}

func setConfig(cmd_args types.CmdArgs, state types.MessageState) error {
	pos_args := cmd_args.PosArgs
	if len(pos_args) < 2 {
		return printErrorAndUsage("Must provide a config to modify!", state)
	}

	config_key := pos_args[1]
	valid_config_keys := getValidConfigKeys()

	fmt.Printf("Valid conf keys: %+v\n", valid_config_keys)

	if len(pos_args) < 3 {
		return printErrorAndUsage("Must provide a value to set!", state)
	}

	if utils.StrContains(config_key, valid_config_keys) {
		config_val := pos_args[2]
		return set(config_key, config_val, state)
	}

	return nil
}

func unsetConfig(cmd_args types.CmdArgs, state types.MessageState) error {
	return nil
}

func set(key string, val string, state types.MessageState) error {
	dbconn := state.G.DBConn

	conf := db.Config{
		GuildId:       &state.M.GuildID,
		ChannelId:     &state.M.ChannelID,
		LastUserToSet: state.M.Author.String(),
		LastUpdated:   time.Now(),

		ConfigName:  key,
	}

	if not_found := dbconn.Find(conf).RecordNotFound(); not_found {

	}

	fmt.Printf("%+v\n", conf)

	//db.G.DBConn.Create(

	return nil
}

func get(key string, state types.MessageState) (string, error) {
	return "", nil
}
