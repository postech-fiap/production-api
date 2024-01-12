package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type pingService struct{}

func NewPingService() *pingService {
	return &pingService{}
}

func (p *pingService) Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
