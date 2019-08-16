package config

import (
	"github.com/order-of-axis-association/AquaBot/types"

	"github.com/order-of-axis-association/AquaBot/admin_funcs"
	"github.com/order-of-axis-association/AquaBot/help_funcs"
	"github.com/order-of-axis-association/AquaBot/util_funcs"
)

var EnabledFuncPackages = []types.FuncPackageConfig {
	admin_funcs.Config,
	help_funcs.Config,
	util_funcs.Config,
}
