package client

import (
	"github.com/gorilla/mux"

	"github.com/Evrynetlabs/evrsdk/client/context"
	"github.com/Evrynetlabs/evrsdk/client/rpc"
)

// Register routes
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
	rpc.RegisterRPCRoutes(cliCtx, r)
}
