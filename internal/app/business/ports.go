package business

import (
	"context"
	"io"
)

// Ports for drive adapters
type (
	// PurchaseConfirmer defines a way to confirm a Purchase
	PurchaseConfirmer interface {
		ConfirmPurchase(context.Context, *Purchase) error
	}

	// PurchasesNotifier defines a way to notify purchases
	PurchasesNotifier interface {
		NotifyPurchases(context.Context) error
	}
)

// Ports for driven adapters
type (
	// PurchaseSaver defines a way to save a Purchase record
	PurchaseSaver interface {
		SavePurchase(context.Context, *Purchase) error
	}

	// PurchaseMessagesReader defines a way to read the pending PurchaseMessage(s)
	PurchaseMessagesReader interface {
		io.Closer
		ReadPurchaseMessages(context.Context) ([]PurchaseMessage, error)
	}

	// PurchaseMessageSender defines a way to send a PurchaseMessage
	PurchaseMessageSender interface {
		SendPurchaseMessage(context.Context, PurchaseMessage) error
	}

	MessageDeliveryConfirmer interface {
		ConfirmMessageDelivery(context.Context, PurchaseMessage) error
	}
)
