package resourceadapter

import (
	"context"

	"github.com/dfc/go/amount"
	. "github.com/dfc/go/protocols/horizon"
	"github.com/dfc/go/services/horizon/internal/assets"
	"github.com/dfc/go/services/horizon/internal/db2/core"
	"github.com/dfc/go/xdr"
)

func PopulateBalance(ctx context.Context, dest *Balance, row core.Trustline) (err error) {
	dest.Type, err = assets.String(row.Assettype)
	if err != nil {
		return
	}

	dest.Balance = amount.String(row.Balance)
	dest.BuyingLiabilities = amount.String(row.BuyingLiabilities)
	dest.SellingLiabilities = amount.String(row.SellingLiabilities)
	dest.Limit = amount.String(row.Tlimit)
	dest.Issuer = row.Issuer
	dest.Code = row.Assetcode
	dest.LastModifiedLedger = row.LastModified
	return
}

func PopulateNativeBalance(dest *Balance, stroops, buyingLiabilities, sellingLiabilities xdr.Int64) (err error) {
	dest.Type, err = assets.String(xdr.AssetTypeAssetTypeNative)
	if err != nil {
		return
	}

	dest.Balance = amount.String(stroops)
	dest.BuyingLiabilities = amount.String(buyingLiabilities)
	dest.SellingLiabilities = amount.String(sellingLiabilities)
	dest.LastModifiedLedger = 0
	dest.Limit = ""
	dest.Issuer = ""
	dest.Code = ""
	return
}
