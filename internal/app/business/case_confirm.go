package business

import (
	"context"
)

func NewPurchaseConfirmer(saver PurchaseSaver) PurchaseConfirmer {
	return purchaseConfirmer{
		saver: saver,
	}
}

type purchaseConfirmer struct {
	saver PurchaseSaver
}

func (p purchaseConfirmer) ConfirmPurchase(ctx context.Context, purchase *Purchase) error {
	if err := purchase.Validate(); err != nil {
		return err
	}

	return p.saver.SavePurchase(ctx, purchase)
}
