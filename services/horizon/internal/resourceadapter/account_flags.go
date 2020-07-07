package resourceadapter

import (
	. "github.com/dfc/go/protocols/horizon"
	"github.com/dfc/go/services/horizon/internal/db2/core"
)

func PopulateAccountFlags(dest *AccountFlags, row core.Account) {
	dest.AuthRequired = row.IsAuthRequired()
	dest.AuthRevocable = row.IsAuthRevocable()
	dest.AuthImmutable = row.IsAuthImmutable()
}
