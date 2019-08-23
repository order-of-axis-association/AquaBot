package admin_funcs

import (
	"github.com/order-of-axis-association/AquaBot/types"
)

func NewConfig() types.FuncPackageConfig {
	return types.FuncPackageConfig {
		Prefix: "$",

		Commands: []types.Command{
			ClearTable,
			CheckTable,
		},
	}
}
