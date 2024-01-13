package usecase

import (
	"github.com/postech-fiap/production-api/internal/core/domain"
	"github.com/postech-fiap/production-api/internal/core/exception"
	"github.com/postech-fiap/production-api/internal/core/port"
	"reflect"
	"testing"
)

func Test_NewOrderUserCase(t *testing.T) {
	t.Run("Create OrderUseCaseInterface", func(t *testing.T) {
		got := NewOrderUserCase(&orderRepositoryMock{})
		_, ok := got.(port.OrderUseCaseInterface)
		if !ok {
			t.Errorf("NewOrderUserCase() should have created port OrderUseCaseInterface")
		}
	})
}

func Test_orderUseCase_List(t *testing.T) {
	mockOrders := mockOrders()

	type fields struct {
		orderRepositoryInterface port.OrderRepositoryInterface
	}

	tests := []struct {
		name    string
		fields  fields
		want    []domain.Order
		wantErr error
	}{
		{
			name: "Should return orders",
			fields: fields{
				orderRepositoryInterface: func() port.OrderRepositoryInterface {
					o := new(orderRepositoryMock)
					o.On("List").Return(mockOrders, nil)
					return o
				}(),
			},
			want: mockOrders,
		},
		{
			name: "Should return error",
			fields: fields{
				orderRepositoryInterface: func() port.OrderRepositoryInterface {
					o := new(orderRepositoryMock)
					o.On("List").Return(nil, mockError())
					return o
				}(),
			},
			wantErr: mockFailedDependencyException(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			orderUseCase := NewOrderUserCase(tt.fields.orderRepositoryInterface)

			gotOrders, gotErr := orderUseCase.List()

			errType := reflect.TypeOf(gotErr)
			wantErrType := reflect.TypeOf(tt.wantErr)

			if errType != wantErrType {
				t.Errorf("List() got error = %v, want error %v", errType, wantErrType)
				return
			}

			if !reflect.DeepEqual(gotOrders, tt.want) {
				t.Errorf("List() got orders = %v, want orders %v", gotOrders, tt.want)
				return
			}
		})
	}
}

func Test_orderUseCase_Insert(t *testing.T) {
	type fields struct {
		orderRepositoryInterface port.OrderRepositoryInterface
	}

	type args struct {
		order *domain.Order
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr error
	}{
		{
			name: "Should insert order",
			fields: fields{
				orderRepositoryInterface: func() port.OrderRepositoryInterface {
					o := new(orderRepositoryMock)
					o.On("Get", mockOrderID).Return(nil, nil)
					o.On("Insert", mockOrder(domain.PENDING)).Return(nil)
					return o
				}(),
			},
			args: args{
				order: mockOrder(""),
			},
		},
		{
			name: "Should not insert order because repository return error on check if exists order",
			fields: fields{
				orderRepositoryInterface: func() port.OrderRepositoryInterface {
					o := new(orderRepositoryMock)
					o.On("Get", mockOrderID).Return(nil, mockError())
					return o
				}(),
			},
			args: args{
				order: mockOrder(""),
			},
			wantErr: mockFailedDependencyException(),
		},
		{
			name: "Should not insert order because exists order",
			fields: fields{
				orderRepositoryInterface: func() port.OrderRepositoryInterface {
					o := new(orderRepositoryMock)
					o.On("Get", mockOrderID).Return(mockOrder(domain.PENDING), nil)
					return o
				}(),
			},
			args: args{
				order: mockOrder(""),
			},
			wantErr: exception.NewConflictException("order already exists", nil),
		},
		{
			name: "Should not insert order because repository return error on insert order",
			fields: fields{
				orderRepositoryInterface: func() port.OrderRepositoryInterface {
					o := new(orderRepositoryMock)
					o.On("Get", mockOrderID).Return(nil, nil)
					o.On("Insert", mockOrder(domain.PENDING)).Return(mockError())
					return o
				}(),
			},
			args: args{
				order: mockOrder(""),
			},
			wantErr: mockFailedDependencyException(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			orderUseCase := NewOrderUserCase(tt.fields.orderRepositoryInterface)

			gotErr := orderUseCase.Insert(tt.args.order)

			errType := reflect.TypeOf(gotErr)
			wantErrType := reflect.TypeOf(tt.wantErr)

			if errType != wantErrType {
				t.Errorf("Insert() got error = %v, want error %v", errType, wantErrType)
				return
			}
		})
	}
}

func Test_orderUseCase_UpdateStatus(t *testing.T) {
	type fields struct {
		orderRepositoryInterface port.OrderRepositoryInterface
	}

	type args struct {
		id        int64
		newStatus domain.Status
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr error
	}{
		{
			name: "Should update status",
			fields: fields{
				orderRepositoryInterface: func() port.OrderRepositoryInterface {
					o := new(orderRepositoryMock)
					o.On("Get", mockOrderID).Return(mockOrder(domain.PENDING), nil)
					o.On("UpdateStatus", mockOrder(domain.RECEIVED)).Return(nil)
					return o
				}(),
			},
			args: args{
				id:        mockOrderID,
				newStatus: domain.RECEIVED,
			},
		},
		{
			name: "Should not update status because repository return error on get order",
			fields: fields{
				orderRepositoryInterface: func() port.OrderRepositoryInterface {
					o := new(orderRepositoryMock)
					o.On("Get", mockOrderID).Return(nil, mockError())
					return o
				}(),
			},
			args: args{
				id:        mockOrderID,
				newStatus: domain.RECEIVED,
			},
			wantErr: mockFailedDependencyException(),
		},
		{
			name: "Should not update status because order not found",
			fields: fields{
				orderRepositoryInterface: func() port.OrderRepositoryInterface {
					o := new(orderRepositoryMock)
					o.On("Get", mockOrderID).Return(nil, nil)
					return o
				}(),
			},
			args: args{
				id:        mockOrderID,
				newStatus: domain.RECEIVED,
			},
			wantErr: exception.NewNotFoundException("order not found", nil),
		},
		{
			name: "Should not update status because not a valid new status",
			fields: fields{
				orderRepositoryInterface: func() port.OrderRepositoryInterface {
					o := new(orderRepositoryMock)
					o.On("Get", mockOrderID).Return(mockOrder(domain.PENDING), nil)
					return o
				}(),
			},
			args: args{
				id:        mockOrderID,
				newStatus: domain.DONE,
			},
			wantErr: exception.NewInvalidDataException("invalid status transition", nil),
		},
		{
			name: "Should not update status because not a valid new status",
			fields: fields{
				orderRepositoryInterface: func() port.OrderRepositoryInterface {
					o := new(orderRepositoryMock)
					o.On("Get", mockOrderID).Return(mockOrder(domain.PENDING), nil)
					o.On("UpdateStatus", mockOrder(domain.RECEIVED)).Return(mockError())
					return o
				}(),
			},
			args: args{
				id:        mockOrderID,
				newStatus: domain.RECEIVED,
			},
			wantErr: mockFailedDependencyException(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			orderUseCase := NewOrderUserCase(tt.fields.orderRepositoryInterface)

			gotErr := orderUseCase.UpdateStatus(tt.args.id, tt.args.newStatus)

			errType := reflect.TypeOf(gotErr)
			wantErrType := reflect.TypeOf(tt.wantErr)

			if errType != wantErrType {
				t.Errorf("UpdateStatus() got error = %v, want error %v", errType, wantErrType)
				return
			}
		})
	}
}
