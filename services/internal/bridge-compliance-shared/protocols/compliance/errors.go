package compliance

import (
	"net/http"

	"github.com/dfc/go/services/internal/bridge-compliance-shared/http/helpers"
)

var (
	// /receive

	// TransactionNotFoundError is an error response
	TransactionNotFoundError = &helpers.ErrorResponse{Code: "transaction_not_found", Mdfcge: "Transaction not found.", Status: http.StatusNotFound}

	// /send

	// CannotResolveDestination is an error response
	CannotResolveDestination = &helpers.ErrorResponse{Code: "cannot_resolve_destination", Mdfcge: "Cannot resolve federated Stellar address.", Status: http.StatusBadRequest}
	// AuthServerNotDefined is an error response
	AuthServerNotDefined = &helpers.ErrorResponse{Code: "auth_server_not_defined", Mdfcge: "No AUTH_SERVER defined in dfc.toml file.", Status: http.StatusBadRequest}
)
