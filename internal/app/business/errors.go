package business

import "strconv"

// Supported values for Error
const (
	_ Error = iota
	ErrMissingPurchaseID
	ErrMissingPurchaseOrderID
)

type Error uint8

func (e Error) Error() string {
	return "error:" + strconv.FormatUint(uint64(e), 10)
}
