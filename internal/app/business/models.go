package business

type Purchase struct {
	ID      uint64
	OrderID uint64
}

func (p Purchase) Validate() error {
	if p.OrderID == 0 {
		return ErrMissingPurchaseID
	}

	return nil
}
