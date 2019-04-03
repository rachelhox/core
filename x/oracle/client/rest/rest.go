package rest

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/gorilla/mux"
)

//nolint
const (
	RestDenom = "denom"
	RestVoter = "voter"
	RestPrice = "price"
)

// RegisterRoutes registers oracle-related REST handlers to a router
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router, cdc *codec.Codec) {
	resgisterTxRoute(cliCtx, r, cdc)
	registerQueryRoute(cliCtx, r, cdc)
}
