package config_cmds

import (
	"github.com/order-of-axis-association/AquaBot/types"
)

var package_prefix = "#"

var commands = []types.Command{
	ConfigCmd,
}

func NewConfig() types.FuncPackageConfig {
	return types.FuncPackageConfig {
		Prefix:   package_prefix,
		Commands: commands,
	}
}
