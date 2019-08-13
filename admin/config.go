package admin

import (
	"github.com/order-of-axis-association/AquaBot/types"
)

var Prefix = "$"

var FuncMap = types.FuncMap {
	"cleartable": ClearTable,
}

var FlagMap = types.FlagMap {
	"cleartable": ClearTableFlags(),
}
