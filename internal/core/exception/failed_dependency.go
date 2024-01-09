package exception

import (
	"github.com/postech-fiap/producao/internal/core/port"
	"net/http"
)

type failedDependencyException struct {
	baseException
}

func NewFailedDependencyException(error error) port.CustomExceptionInterface {
	return &failedDependencyException{
		baseException{
			statusCode: http.StatusFailedDependency,
			message:    "failed internal dependency",
			error:      error,
		},
	}
}
