package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client/context"
)

// RegisterRoutes registers voter-related REST handlers to a router
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
	// this line is used by starport scaffolding
	r.HandleFunc("/voter/poll", ListPollHandler(cliCtx, "voter")).Methods("GET")
	r.HandleFunc("/voter/poll", createPollHandler(cliCtx)).Methods("POST")
}
