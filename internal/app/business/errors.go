package business

import "strconv"

// Supported values for Error
const (
	_ Error = iota
	ErrMissingPurchaseID
	ErrMissingPurchaseOrderID
	ErrDuplicatedOrderID
)

type Error uint8

func (e Error) Error() string {
	return "0" + strconv.FormatUint(uint64(e), 10)
}
