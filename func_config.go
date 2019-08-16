package main

import (
	"github.com/order-of-axis-association/AquaBot/types"

	"github.com/order-of-axis-association/AquaBot/admin_funcs"
	"github.com/order-of-axis-association/AquaBot/util_funcs"
)

var FUNC_CONFIG = []types.FuncPackageConfig {
	admin_funcs.Config,
	util_funcs.Config,
}
