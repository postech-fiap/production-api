package middlewares

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/postech-fiap/producao/internal/core/port"
	"net/http"
)

type customExceptionResponse struct {
	statusCode int
	Message    string `json:"message,omitempty"`
}

func ErrorHandler(c *gin.Context) {
	c.Next()

	if len(c.Errors) == 0 {
		return
	}

	err := c.Errors[0]
	response := customExceptionResponse{
		statusCode: http.StatusInternalServerError,
	}

	var x port.CustomExceptionInterface
	isACustomExceptionType := errors.As(err, &x)
	if isACustomExceptionType {
		response.statusCode = x.GetStatusCode()
		response.Message = x.GetMessage()
	}

	fmt.Println(err)
	c.JSON(response.statusCode, response)
	return
}
