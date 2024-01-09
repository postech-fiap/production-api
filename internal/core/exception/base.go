package exception

type baseException struct {
	statusCode int
	message    string
	error      error
}

func (b *baseException) GetStatusCode() int {
	return b.statusCode
}

func (b *baseException) GetMessage() string {
	return b.message
}

func (b *baseException) Error() string {
	if b.error != nil {
		return b.error.Error()
	}
	return b.message
}
