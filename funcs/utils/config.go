package util_funcs

import (
	"github.com/order-of-axis-association/AquaBot/types"
)

var package_prefix = "!"

func NewConfig() types.FuncPackageConfig {
	return types.FuncPackageConfig {
		Prefix: package_prefix,
		Commands: []types.Command{
			Ping,
			Pong,
			Todo,
		},
	}
}
