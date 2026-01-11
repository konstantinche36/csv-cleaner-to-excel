package model

type Record struct {
	Data   string
	Name   string
	Number string
	City   string
	Amount float64
	Status string
}

type ErrorCode string

const (
	ErrPhoneEmpty     ErrorCode = "PHONE_EMPTY"
	ErrPhoneInvalid   ErrorCode = "PHONE_INVALID"
	ErrDateInvalid    ErrorCode = "DATE_INVALID"
	ErrAmountEmpty    ErrorCode = "AMOUNT_EMPTY"
	ErrAmountInvalid  ErrorCode = "AMOUNT_INVALID"
	ErrAmountNegative ErrorCode = "AMOUNT_NEGATIVE"
)

type InvalidRecord struct {
	LineNumber   int
	ErrorCode    ErrorCode
	ErrorMessage string
	RawLine      string
}
