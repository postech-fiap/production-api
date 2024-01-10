package exception

import (
	"github.com/postech-fiap/production-api/internal/core/port"
	"net/http"
)

type conflictException struct {
	baseException
}

func NewConflictException(message string, error error) port.CustomExceptionInterface {
	return &conflictException{
		baseException{
			statusCode: http.StatusConflict,
			message:    message,
			error:      error,
		},
	}
}
