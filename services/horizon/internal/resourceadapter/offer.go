package resourceadapter

import (
	"context"

	"github.com/dfc/go/amount"
	. "github.com/dfc/go/protocols/horizon"
	"github.com/dfc/go/services/horizon/internal/db2/core"
	"github.com/dfc/go/services/horizon/internal/db2/history"
	"github.com/dfc/go/services/horizon/internal/httpx"
	"github.com/dfc/go/support/render/hal"
)

func PopulateOffer(ctx context.Context, dest *Offer, row core.Offer, ledger *history.Ledger) {
	dest.ID = row.OfferID
	dest.PT = row.PagingToken()
	dest.Seller = row.SellerID
	dest.Amount = amount.String(row.Amount)
	dest.PriceR.N = row.Pricen
	dest.PriceR.D = row.Priced
	dest.Price = row.PriceAsString()

	row.SellingAsset.MustExtract(&dest.Selling.Type, &dest.Selling.Code, &dest.Selling.Issuer)
	row.BuyingAsset.MustExtract(&dest.Buying.Type, &dest.Buying.Code, &dest.Buying.Issuer)

	dest.LastModifiedLedger = row.Lastmodified
	if ledger != nil {
		dest.LastModifiedTime = &ledger.ClosedAt
	}
	lb := hal.LinkBuilder{httpx.BaseURL(ctx)}
	dest.Links.Self = lb.Linkf("/offers/%d", row.OfferID)
	dest.Links.OfferMaker = lb.Linkf("/accounts/%s", row.SellerID)
	return
}
