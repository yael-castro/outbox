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

	// MessagesRelay defines a way to notify purchases
	MessagesRelay interface {
		RelayMessages(context.Context) error
	}
)

// Ports for driven adapters
type (
	// PurchaseSaver defines a way to save a Purchase record
	PurchaseSaver interface {
		SavePurchase(context.Context, *Purchase) error
	}

	// MessagesReader defines a way to read the pending PurchaseMessage(s)
	MessagesReader interface {
		io.Closer
		ReadMessages(context.Context) ([]Message, error)
	}

	// MessageSender defines a way to send a Message
	MessageSender interface {
		SendMessage(context.Context, *Message) error
	}

	MessageDeliveryConfirmer interface {
		ConfirmMessageDelivery(context.Context, uint64) error
	}
)
