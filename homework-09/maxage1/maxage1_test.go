package main

import "testing"

func TestMaxAge(t *testing.T) {
	type args struct {
		p []ager
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "when customers only",
			args: args{
				p: []ager{
					&customer{Age: 33},
					&customer{Age: 55},
					&customer{Age: 44},
				},
			},
			want: 55,
		},
		{
			name: "when employees only",
			args: args{
				[]ager{
					&employee{Age: 33},
					&employee{Age: 55},
					&employee{Age: 44},
				},
			},
			want: 55,
		},
		{
			name: "when employees and customers",
			args: args{
				[]ager{
					&customer{Age: 77},
					&employee{Age: 33},
					&customer{Age: 55},
					&employee{Age: 44},
					&customer{Age: 66},
				},
			},
			want: 77,
		},
		{
			name: "when colleection is empty",
			args: args{
				[]ager{},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxAge(tt.args.p...); got != tt.want {
				t.Errorf("MaxAge() = %v, want %v", got, tt.want)
			}
		})
	}
}
