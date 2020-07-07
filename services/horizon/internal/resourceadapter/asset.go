package resourceadapter

import (
	"context"

	. "github.com/dfc/go/protocols/horizon"
	"github.com/dfc/go/xdr"
)

func PopulateAsset(ctx context.Context, dest *Asset, asset xdr.Asset) error {
	return asset.Extract(&dest.Type, &dest.Code, &dest.Issuer)
}
