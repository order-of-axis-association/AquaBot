package admin

import (
	"github.com/order-of-axis-association/AquaBot/types"
)

var FuncMap = types.FuncMap {
	"cleardb": ClearDB,
}

var FlagMap = types.FlagMap {
	"cleardb": nil,
}
