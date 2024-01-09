package exception

import (
	"github.com/postech-fiap/producao/internal/core/port"
	"net/http"
)

type notFoundException struct {
	baseException
}

func NewNotFoundException(message string, error error) port.CustomExceptionInterface {
	return &notFoundException{
		baseException{
			statusCode: http.StatusNotFound,
			message:    message,
			error:      error,
		},
	}
}
