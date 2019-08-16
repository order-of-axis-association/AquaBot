package admin_funcs

import (
	"github.com/order-of-axis-association/AquaBot/types"
)

var Config = types.FuncPackageConfig {
	Prefix:   "$",

	Commands: []types.Command {
		ClearTable,
		CheckTable,
	},
}
