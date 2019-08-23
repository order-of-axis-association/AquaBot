package config

import (
	"github.com/order-of-axis-association/AquaBot/types"

	"github.com/order-of-axis-association/AquaBot/funcs/admin"
	"github.com/order-of-axis-association/AquaBot/funcs/config"
	"github.com/order-of-axis-association/AquaBot/funcs/help"
	"github.com/order-of-axis-association/AquaBot/funcs/utils"
)

var EnabledFuncPackages = []types.FuncPackageConfig{
	admin_funcs.NewConfig(),
	config_funcs.NewConfig(),
	help_funcs.NewConfig(),
	util_funcs.NewConfig(),
}
