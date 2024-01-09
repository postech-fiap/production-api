package port

type CustomExceptionInterface interface {
	error
	GetStatusCode() int
	GetMessage() string
}
