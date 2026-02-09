package constants

type PaymentMethod string
type PaymemtStatus string

const (
	ORDER_CREDIT_CARD PaymentMethod = "credit_card"
	ORDER_DOC         PaymentMethod = "cod"
)

const (
	ORDER_PENDING   PaymemtStatus = "pedning"
	ORDER_COMPLETED PaymemtStatus = "completed"
	ORDER_REJECT    PaymemtStatus = "reject"
)
