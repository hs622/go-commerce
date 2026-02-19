package constants

type ErrorMessage string

const (
	V_INVALID_EMAIL      ErrorMessage = ""
	V_INVALID_NUMBER     ErrorMessage = ""
	V_INVALID_STRING     ErrorMessage = ""
	V_INVALID_LENGHT     ErrorMessage = ""
	V_INVALID_MIN_LENGHT ErrorMessage = ""
	V_INVALID_MAX_LENGHT ErrorMessage = ""
	V_INVALID_UUID       ErrorMessage = "Invalid resource id."
	V_INVALID_ARGUMENT   ErrorMessage = "invalid argument provided."
)
