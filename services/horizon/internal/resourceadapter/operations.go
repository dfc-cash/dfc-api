package resourceadapter

import (
	"context"
	"fmt"

	"github.com/dfc/go/protocols/horizon/operations"
	"github.com/dfc/go/services/horizon/internal/db2/history"
	"github.com/dfc/go/services/horizon/internal/httpx"
	"github.com/dfc/go/support/render/hal"
	"github.com/dfc/go/xdr"
)

// NewOperation creates a new operation resource, finding the appropriate type to use
// based upon the row's type.
func NewOperation(
	ctx context.Context,
	row history.Operation,
	ledger history.Ledger,
) (result hal.Pageable, err error) {

	base := operations.Base{}
	PopulateBaseOperation(ctx, &base, row, ledger)

	switch row.Type {
	case xdr.OperationTypeBumpSequence:
		e := operations.BumpSequence{Base: base}
		err = row.UnmarshalDetails(&e)
		result = e
	case xdr.OperationTypeCreateAccount:
		e := operations.CreateAccount{Base: base}
		err = row.UnmarshalDetails(&e)
		result = e
	case xdr.OperationTypePayment:
		e := operations.Payment{Base: base}
		err = row.UnmarshalDetails(&e)
		result = e
	case xdr.OperationTypePathPayment:
		e := operations.PathPayment{}
		e.Payment.Base = base
		err = row.UnmarshalDetails(&e)
		result = e
	case xdr.OperationTypeManageOffer:
		e := operations.ManageOffer{}
		e.CreatePassiveOffer.Base = base
		err = row.UnmarshalDetails(&e)
		result = e
	case xdr.OperationTypeCreatePassiveOffer:
		e := operations.CreatePassiveOffer{Base: base}
		err = row.UnmarshalDetails(&e)
		result = e
	case xdr.OperationTypeSetOptions:
		e := operations.SetOptions{Base: base}
		err = row.UnmarshalDetails(&e)
		result = e
	case xdr.OperationTypeChangeTrust:
		e := operations.ChangeTrust{Base: base}
		err = row.UnmarshalDetails(&e)
		result = e
	case xdr.OperationTypeAllowTrust:
		e := operations.AllowTrust{Base: base}
		err = row.UnmarshalDetails(&e)
		result = e
	case xdr.OperationTypeAccountMerge:
		e := operations.AccountMerge{Base: base}
		err = row.UnmarshalDetails(&e)
		result = e
	case xdr.OperationTypeInflation:
		e := operations.Inflation{Base: base}
		err = row.UnmarshalDetails(&e)
		result = e
	case xdr.OperationTypeManageData:
		e := operations.ManageData{Base: base}
		err = row.UnmarshalDetails(&e)
		result = e
	default:
		result = base
	}

	return
}

// Populate fills out this resource using `row` as the source.
func PopulateBaseOperation(
	ctx context.Context,
	dest *operations.Base,
	row history.Operation,
	ledger history.Ledger,
) {
	dest.ID = fmt.Sprintf("%d", row.ID)
	dest.PT = row.PagingToken()
	// Check db2/history.Transaction.Successful field comment for more information.
	if row.TransactionSuccessful == nil {
		dest.TransactionSuccessful = true
	} else {
		dest.TransactionSuccessful = *row.TransactionSuccessful
	}
	dest.SourceAccount = row.SourceAccount
	populateOperationType(dest, row)
	dest.LedgerCloseTime = ledger.ClosedAt
	dest.TransactionHash = row.TransactionHash

	lb := hal.LinkBuilder{Base: httpx.BaseURL(ctx)}
	self := fmt.Sprintf("/operations/%d", row.ID)
	dest.Links.Self = lb.Link(self)
	dest.Links.Succeeds = lb.Linkf("/effects?order=desc&cursor=%s", dest.PT)
	dest.Links.Precedes = lb.Linkf("/effects?order=asc&cursor=%s", dest.PT)
	dest.Links.Transaction = lb.Linkf("/transactions/%s", row.TransactionHash)
	dest.Links.Effects = lb.Link(self, "effects")
}

func populateOperationType(dest *operations.Base, row history.Operation) {
	var ok bool
	dest.TypeI = int32(row.Type)
	dest.Type, ok = operations.TypeNames[row.Type]

	if !ok {
		dest.Type = "unknown"
	}
}
