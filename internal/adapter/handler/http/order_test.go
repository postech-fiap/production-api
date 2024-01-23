package http

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/postech-fiap/production-api/internal/adapter/handler/http/middlewares"
	"github.com/postech-fiap/production-api/internal/core/port"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_NewOrderService(t *testing.T) {
	t.Run("Create Order Service", func(t *testing.T) {
		got := NewOrderService(&orderUseCaseMock{})
		if got == nil {
			t.Errorf("NewOrderService() should have created orderService")
		}
	})
}

func Test_orderService_List(t *testing.T) {
	type fields struct {
		orderUseCase port.OrderUseCaseInterface
	}

	tests := []struct {
		name     string
		fields   fields
		wantCode int
		wantBody string
	}{
		{
			name: "Should 200",
			fields: fields{
				orderUseCase: func() port.OrderUseCaseInterface {
					b := new(orderUseCaseMock)
					b.On("List").Return(mockDomainOrderList(), nil)
					return b
				}(),
			},
			wantCode: http.StatusOK,
			wantBody: mockOrderListResponse(),
		},
		{
			name: "Should 424",
			fields: fields{
				orderUseCase: func() port.OrderUseCaseInterface {
					b := new(orderUseCaseMock)
					b.On("List").Return(nil, mockFailedDependencyException())
					return b
				}(),
			},
			wantCode: http.StatusFailedDependency,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			orderService := &orderService{
				orderUseCase: tt.fields.orderUseCase,
			}

			router := gin.New()
			router.Use(middlewares.ErrorService)
			router.GET("/order", orderService.List)

			response := httptest.NewRecorder()
			request := httptest.NewRequest("GET", "/order", nil)
			router.ServeHTTP(response, request)

			code := response.Code
			if tt.wantCode != 0 && tt.wantCode != code {
				t.Errorf("List() code = %v, want %v", code, tt.wantCode)
			}

			body := response.Body.String()
			if len(tt.wantBody) > 0 && tt.wantBody != body {
				t.Errorf("List() body = %v, want %v", body, tt.wantBody)
			}
		})
	}
}

func Test_orderService_Insert(t *testing.T) {
	type fields struct {
		orderUseCase port.OrderUseCaseInterface
	}

	type args struct {
		body string
	}

	tests := []struct {
		name     string
		fields   fields
		args     args
		wantCode int
	}{
		{
			name: "Should 201",
			fields: fields{
				orderUseCase: func() port.OrderUseCaseInterface {
					b := new(orderUseCaseMock)
					b.On("Insert", mock.Anything).Return(nil)
					return b
				}(),
			},
			args: args{
				body: mockInsertOrderRequest,
			},
			wantCode: http.StatusCreated,
		},
		{
			name: "Should 400",
			fields: fields{
				orderUseCase: func() port.OrderUseCaseInterface {
					return new(orderUseCaseMock)
				}(),
			},
			args: args{
				body: "{}",
			},
			wantCode: http.StatusBadRequest,
		},
		{
			name: "Should 424",
			fields: fields{
				orderUseCase: func() port.OrderUseCaseInterface {
					b := new(orderUseCaseMock)
					b.On("Insert", mock.Anything).Return(mockFailedDependencyException())
					return b
				}(),
			},
			args: args{
				body: mockInsertOrderRequest,
			},
			wantCode: http.StatusFailedDependency,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			orderService := &orderService{
				orderUseCase: tt.fields.orderUseCase,
			}

			router := gin.New()
			router.Use(middlewares.ErrorService)
			router.POST("/order", orderService.Insert)

			response := httptest.NewRecorder()
			request := httptest.NewRequest("POST", "/order", bytes.NewBufferString(tt.args.body))
			router.ServeHTTP(response, request)

			code := response.Code
			if tt.wantCode != 0 && tt.wantCode != code {
				t.Errorf("Insert() code = %v, want %v", code, tt.wantCode)
			}
		})
	}
}

func Test_orderService_UpdateStatus(t *testing.T) {
	type fields struct {
		orderUseCase port.OrderUseCaseInterface
	}

	type args struct {
		id   int64
		body string
	}

	tests := []struct {
		name     string
		fields   fields
		args     args
		wantCode int
	}{
		{
			name: "Should 202",
			fields: fields{
				orderUseCase: func() port.OrderUseCaseInterface {
					b := new(orderUseCaseMock)
					b.On("UpdateStatus", mockOrderId, mockStatus).Return(nil)
					return b
				}(),
			},
			args: args{
				id:   mockOrderId,
				body: mockUpdateOrderStatusRequest,
			},
			wantCode: http.StatusAccepted,
		},
		{
			name: "Should 400 due path param",
			fields: fields{
				orderUseCase: func() port.OrderUseCaseInterface {
					return new(orderUseCaseMock)
				}(),
			},
			args: args{
				id:   0,
				body: mockUpdateOrderStatusRequest,
			},
			wantCode: http.StatusBadRequest,
		},
		{
			name: "Should 400 due path param",
			fields: fields{
				orderUseCase: func() port.OrderUseCaseInterface {
					return new(orderUseCaseMock)
				}(),
			},
			args: args{
				id:   mockOrderId,
				body: `{}`,
			},
			wantCode: http.StatusBadRequest,
		},
		{
			name: "Should 424",
			fields: fields{
				orderUseCase: func() port.OrderUseCaseInterface {
					b := new(orderUseCaseMock)
					b.On("UpdateStatus", mockOrderId, mockStatus).Return(mockFailedDependencyException())
					return b
				}(),
			},
			args: args{
				id:   mockOrderId,
				body: mockUpdateOrderStatusRequest,
			},
			wantCode: http.StatusFailedDependency,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			orderService := &orderService{
				orderUseCase: tt.fields.orderUseCase,
			}

			router := gin.New()
			router.Use(middlewares.ErrorService)
			router.PUT("/order/:id/status", orderService.UpdateStatus)

			response := httptest.NewRecorder()
			uri := fmt.Sprintf("/order/%d/status", tt.args.id)
			request := httptest.NewRequest("PUT", uri, bytes.NewBufferString(tt.args.body))
			router.ServeHTTP(response, request)

			code := response.Code
			if tt.wantCode != 0 && tt.wantCode != code {
				t.Errorf("UpdateStatus() code = %v, want %v", code, tt.wantCode)
			}
		})
	}
}
