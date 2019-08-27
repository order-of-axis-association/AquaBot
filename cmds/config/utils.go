package config_cmds

import (
	"fmt"
	_ "strings"

	"github.com/order-of-axis-association/AquaBot/types"

    "github.com/order-of-axis-association/AquaBot/cmds/admin"
    "github.com/order-of-axis-association/AquaBot/cmds/utils"

    "github.com/order-of-axis-association/AquaBot/utils"
)

func printErrorAndUsage(error_msg string, state types.MessageState) error {
    msg := fmt.Sprintf("%s\nCommand usage is: ```%s```", error_msg, ConfigUsage)
    return utils.TempSay(msg, state)
}

func getValidConfigKeys() []string {
	// Cannot include config_cmds.Config because of initialization loops.
	// config_cmds.Config will be a one-off that's explicitly added.
	// help_cmds is not included due to import cycles. I _really_ need to find an alternative approach.
	var EnabledFuncPackages = []types.FuncPackageConfig {
		admin_cmds.NewConfig(),
		util_cmds.NewConfig(),
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
	config_cmds := []types.ConfigFlags {
		ConfigConfigs,
	}

	for _, config := range config_cmds {
		for key, _ := range config {
			config_keys = append(config_keys, key)
		}
	}


	return config_keys
}
