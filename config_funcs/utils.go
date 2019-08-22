package config_funcs

import (
	"fmt"
	_ "strings"

	"github.com/order-of-axis-association/AquaBot/types"

	"github.com/order-of-axis-association/AquaBot/admin_funcs"
	"github.com/order-of-axis-association/AquaBot/util_funcs"
	"github.com/order-of-axis-association/AquaBot/utils"
)

func printErrorAndUsage(error_msg string, state types.MessageState) error {
	msg := fmt.Sprintf("%s\nCommand usage is: ```%s```", error_msg, ConfigUsage)
	return utils.TempSay(msg, state)
}

func getValidConfigKeys() []string {
	// Cannot include config_funcs.Config because of initialization loops.
	// config_funcs.Config will be a one-off that's explicitly added.
	// help_funcs is not included due to import cycles. I _really_ need to find an alternative approach.
	var EnabledFuncPackages = []types.FuncPackageConfig{
		admin_funcs.Config,
		util_funcs.Config,
	}

	config_keys := make([]string, 0)
	for _, package_config := range EnabledFuncPackages {
		for _, command := range package_config.Commands {
			config_flags := command.Configs
			if config_flags != nil {
				for config_key, _ := range config_flags {
					config_keys = append(config_keys, config_key)
				}
			}
		}
	}

	// config_func configs... because initialization loops/import cycles...
	config_cmds := []types.ConfigFlags{
		ConfigConfigs,
	}

	for _, config := range config_cmds {
		for key, _ := range config {
			config_keys = append(config_keys, key)
		}
	}

	return config_keys
}
