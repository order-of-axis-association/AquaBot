package config_funcs

import (
	"github.com/order-of-axis-association/AquaBot/types"
)

var package_prefix = "#"

var commands = []types.Command{
	ConfigCmd,
}

var Config = types.FuncPackageConfig{
	Prefix:   package_prefix,
	Commands: commands,
}
