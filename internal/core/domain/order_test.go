package domain

import (
	"testing"
)

func Test_Order_IsValidStatus(t *testing.T) {
	type args struct {
		order     Order
		newStatus Status
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Should return true on change from PENDING TO RECEIVED",
			args: args{
				order: Order{
					Status: PENDING,
				},
				newStatus: RECEIVED,
			},
			want: true,
		},
		{
			name: "Should return true on change from RECEIVED TO IN_PREPARE",
			args: args{
				order: Order{
					Status: RECEIVED,
				},
				newStatus: IN_PREPARE,
			},
			want: true,
		},
		{
			name: "Should return true on change from IN_PREPARE TO DONE",
			args: args{
				order: Order{
					Status: IN_PREPARE,
				},
				newStatus: DONE,
			},
			want: true,
		},
		{
			name: "Should return true on change from FINISHED TO DONE",
			args: args{
				order: Order{
					Status: DONE,
				},
				newStatus: FINISHED,
			},
			want: true,
		},
		{
			name: "Should return false on incorrect change",
			args: args{
				order: Order{
					Status: RECEIVED,
				},
				newStatus: DONE,
			},
			want: false,
		},
		{
			name: "Should return false on invalid new status",
			args: args{
				order: Order{
					Status: RECEIVED,
				},
				newStatus: "abc",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.order.IsValidStatus(tt.args.newStatus)

			if got != tt.want {
				t.Errorf("IsValidStatus() got = %v, want %v", got, tt.want)
				return
			}
		})
	}
}
