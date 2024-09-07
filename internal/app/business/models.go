package business

import "fmt"

type Purchase struct {
	ID      uint64
	OrderID uint64
}

func (p Purchase) Validate() error {
	if p.OrderID == 0 {
		return fmt.Errorf("missing order id to confirm the purchase (%w)", ErrMissingPurchaseID)
	}

	return nil
}

type Message struct {
	ID      uint64
	Key     string
	Topic   string
	Header  []byte
	Content []byte
}
