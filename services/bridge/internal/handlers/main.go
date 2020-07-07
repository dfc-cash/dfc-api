package handlers

import (
	"github.com/dfc/go/clients/federation"
	"github.com/dfc/go/clients/horizon"
	"github.com/dfc/go/clients/dfctoml"
	"github.com/dfc/go/services/bridge/internal/config"
	"github.com/dfc/go/services/bridge/internal/db"
	"github.com/dfc/go/services/bridge/internal/listener"
	"github.com/dfc/go/services/bridge/internal/submitter"
	"github.com/dfc/go/support/http"
)

// RequestHandler implements bridge server request handlers
type RequestHandler struct {
	Config               *config.Config                          `inject:""`
	Client               http.SimpleHTTPClientInterface          `inject:""`
	Horizon              horizon.ClientInterface                 `inject:""`
	Database             db.Database                             `inject:""`
	StellarTomlResolver  dfctoml.ClientInterface             `inject:""`
	FederationResolver   federation.ClientInterface              `inject:""`
	TransactionSubmitter submitter.TransactionSubmitterInterface `inject:""`
	PaymentListener      *listener.PaymentListener               `inject:""`
}

func (rh *RequestHandler) isAssetAllowed(code string, issuer string) bool {
	for _, asset := range rh.Config.Assets {
		if asset.Code == code && asset.Issuer == issuer {
			return true
		}
	}
	return false
}
