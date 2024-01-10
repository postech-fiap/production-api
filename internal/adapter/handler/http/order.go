package http

import (
	"github.com/gin-gonic/gin"
	"github.com/postech-fiap/production-api/internal/adapter/handler/http/dto"
	"github.com/postech-fiap/production-api/internal/adapter/handler/http/mapper"
	"github.com/postech-fiap/production-api/internal/core/domain"
	"github.com/postech-fiap/production-api/internal/core/exception"
	"github.com/postech-fiap/production-api/internal/core/port"
	"net/http"
)

type orderService struct {
	orderUseCase port.OrderUseCaseInterface
}

func NewOrderService(orderUseCase port.OrderUseCaseInterface) *orderService {
	return &orderService{
		orderUseCase: orderUseCase,
	}
}

func (o *orderService) List(c *gin.Context) {
	orders, err := o.orderUseCase.List()
	if err != nil {
		c.Error(exception.NewInvalidDataException("invalid body", err))
		return
	}
	ordersResponse := mapper.MapDomainToOrderWrapperDto(orders)
	c.JSON(http.StatusOK, ordersResponse)
}

func (o *orderService) Insert(c *gin.Context) {
	var requestBody dto.OrderInsertRequest
	err := c.ShouldBind(&requestBody)
	if err != nil {
		c.Error(exception.NewInvalidDataException("invalid body", err))
		return
	}

	newOrder := mapper.MapInsertDTOToDomain(&requestBody)

	err = o.orderUseCase.Insert(newOrder)
	if err != nil {
		c.Error(err)
		return
	}

	c.Status(http.StatusCreated)
}

func (o *orderService) SetStatus(c *gin.Context) {
	var requestURIParams dto.UpdateStatusRequestURI
	err := c.ShouldBindUri(&requestURIParams)
	if err != nil {
		c.Error(exception.NewInvalidDataException("invalid param id", err))
		return
	}

	var requestBody dto.UpdateStatusRequestBody
	err = c.ShouldBind(&requestBody)
	if err != nil {
		c.Error(exception.NewInvalidDataException("invalid body", err))
		return
	}

	statusToSet := domain.Status(requestBody.Status)
	err = o.orderUseCase.UpdateStatus(requestURIParams.ID, statusToSet)
	if err != nil {
		c.Error(err)
		return
	}

	c.Status(http.StatusAccepted)
}
