package keeper

import (
	"github.com/cosmostation/mintstation/x/mns/types"
)

var _ types.QueryServer = Keeper{}
