package business

import "context"

// Ports for drive adapters
type (
	// PurchaseConfirmer defines a way to confirm a Purchase
	PurchaseConfirmer interface {
		ConfirmPurchase(context.Context, *Purchase) error
	}
)

// Ports for driven adapters
type (
	// PurchaseSaver defines a way to save a Purchase record
	PurchaseSaver interface {
		SavePurchase(context.Context, *Purchase) error
	}
)
