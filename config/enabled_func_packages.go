package config

import (
	"github.com/order-of-axis-association/AquaBot/types"

	"github.com/order-of-axis-association/AquaBot/cmds/admin"
	"github.com/order-of-axis-association/AquaBot/cmds/config"
	"github.com/order-of-axis-association/AquaBot/cmds/help"
	"github.com/order-of-axis-association/AquaBot/cmds/utils"
)

var EnabledFuncPackages = []types.FuncPackageConfig{
	admin_cmds.NewConfig(),
	config_cmds.NewConfig(),
	help_cmds.NewConfig(),
	util_cmds.NewConfig(),
}
