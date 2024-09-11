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

type Headers = []Header

type Header struct {
	Key   string
	Value []byte
}

type Message struct {
	ID      uint64
	Topic   string
	Key     []byte
	Value   []byte
	Headers Headers
}
