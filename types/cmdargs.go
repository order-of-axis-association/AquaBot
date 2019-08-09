package types

import "github.com/jinzhu/gorm"

type CmdArgs struct {
	Cmd				string
	PosArgs			map[int]string	  // Key is positional int
	ShortFlagArgs	map[string]string // Key is short arg without -
	LongFlagArgs	map[string]string // Key is long arg without --
}
