package main

import (
	"net/http"

	"github.com/dfc/go/clients/horizon"
	"github.com/dfc/go/services/friendbot/internal"
	"github.com/dfc/go/strkey"
)

func initFriendbot(
	friendbotSecret string,
	networkPassphrase string,
	horizonURL string,
	startingBalance string,
) *internal.Bot {

	if friendbotSecret == "" || networkPassphrase == "" || horizonURL == "" || startingBalance == "" {
		return nil
	}

	// ensure its a seed if its not blank
	strkey.MustDecode(strkey.VersionByteSeed, friendbotSecret)

	return &internal.Bot{
		Secret: friendbotSecret,
		Horizon: &horizon.Client{
			URL:  horizonURL,
			HTTP: http.DefaultClient,
		},
		Network:           networkPassphrase,
		StartingBalance:   startingBalance,
		SubmitTransaction: internal.AsyncSubmitTransaction,
	}
}