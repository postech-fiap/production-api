package http

import (
	"github.com/gin-gonic/gin"
	"github.com/postech-fiap/production-api/internal/adapter/handler/http/dto"
	"github.com/postech-fiap/production-api/internal/core/domain"
	"github.com/postech-fiap/production-api/internal/core/exception"
	"github.com/postech-fiap/production-api/internal/core/port"
	"net/http"
)

type orderHandler struct {
	orderService port.OrderServiceInterface
}

func NewOrderHandler(orderService port.OrderServiceInterface) *orderHandler {
	return &orderHandler{
		orderService: orderService,
	}
}

func (o *orderHandler) List(c *gin.Context) {
	orders, err := o.orderService.List()
	if err != nil {
		c.Error(exception.NewInvalidDataException("invalid body", err))
		return
	}
	c.JSON(http.StatusOK, orders)
}

func (o *orderHandler) Insert(c *gin.Context) {
	var requestBody dto.OrderInsertRequest
	err := c.ShouldBind(&requestBody)
	if err != nil {
		c.Error(exception.NewInvalidDataException("invalid body", err))
		return
	}

	newOrder := o.mapInsertDTOToDomain(&requestBody)

	err = o.orderService.Insert(newOrder)
	if err != nil {
		c.Error(err)
		return
	}

	c.Status(http.StatusCreated)
}

func (o *orderHandler) SetStatus(c *gin.Context) {
	var requestURIParams dto.UpdateStatusRequestURI
	err := c.ShouldBindUri(&requestURIParams)
	if err != nil {
		c.Error(exception.NewInvalidDataException("invalid id param", err))
		return
	}

	var requestBody dto.UpdateStatusRequestBody
	err = c.ShouldBind(&requestBody)
	if err != nil {
		c.Error(exception.NewInvalidDataException("invalid body", err))
		return
	}

	statusToSet := domain.Status(requestBody.Status)
	err = o.orderService.UpdateStatus(requestURIParams.ID, statusToSet)
	if err != nil {
		c.Error(err)
		return
	}

	c.Status(http.StatusAccepted)
}

func (o *orderHandler) mapInsertDTOToDomain(orderDTO *dto.OrderInsertRequest) *domain.Order {
	items := make([]domain.Item, 0)
	for _, itemDTO := range orderDTO.Items {
		item := domain.Item{
			Name:     itemDTO.Name,
			Quantity: itemDTO.Quantity,
			Comment:  itemDTO.Comment,
		}
		items = append(items, item)
	}

	return &domain.Order{
		ID:          orderDTO.ID,
		Number:      orderDTO.Number,
		CreatedDate: orderDTO.CreatedDate,
		Items:       items,
	}
}
