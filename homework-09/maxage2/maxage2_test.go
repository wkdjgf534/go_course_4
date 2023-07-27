package main

import (
	"reflect"
	"testing"
)

func TestMaxAge(t *testing.T) {
	type args struct {
		users []any
	}
	tests := []struct {
		name string
		args args
		want any
	}{
		{
			name: "when customers only",
			args: args{
				users: []any{
					customer{Age: 22},
					customer{Age: 44},
					customer{Age: 33},
				}},
			want: customer{Age: 44},
		},
		{
			name: "when employees only",
			args: args{
				users: []any{
					employee{Age: 22},
					employee{Age: 44},
					employee{Age: 33},
				}},
			want: employee{Age: 44},
		},
		{
			name: "when employees and customers",
			args: args{
				users: []any{
					customer{Age: 22},
					employee{Age: 33},
					customer{Age: 44},
					employee{Age: 66},
					customer{Age: 55},
				}},
			want: employee{Age: 66},
		},
		{
			name: "when collection is empty",
			args: args{
				users: []any{},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxAge(tt.args.users...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MaxAge() = %v, want %v", got, tt.want)
			}
		})
	}
}
