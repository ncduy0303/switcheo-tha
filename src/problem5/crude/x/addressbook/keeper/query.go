package keeper

import (
	"crude/x/addressbook/types"
)

var _ types.QueryServer = Keeper{}
