package constants

type PaymentMethod string
type PaymemtStatus string

const (
	// payment method
	ORDER_CREDIT_CARD PaymentMethod = "credit_card"
	ORDER_DOC         PaymentMethod = "cod"

	// payment status
	ORDER_PENDING   PaymemtStatus = "pedning"
	ORDER_COMPLETED PaymemtStatus = "completed"
	ORDER_REJECT    PaymemtStatus = "reject"
	ORDER_WITHDRAW  PaymemtStatus = "withdraw"

	//

)

// work with iota
// func (ps PaymemtStatus) String() string {
// 	return [...]string{"pending", "completed", "reject"}[ps]
// }
