package main

import (
	"bytes"
	"testing"
)

func TestWriteToFile(t *testing.T) {
	type args struct {
		data []any
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "when strings",
			args: args{data: []any{"Hello", "World"}},
			want: "HelloWorld",
		},
		{
			name: "when numbers",
			args: args{data: []any{1, 2, 3}},
			want: "",
		},
		{
			name: "when booleans",
			args: args{data: []any{true, false}},
			want: "",
		},
		{
			name: "when an empty input",
			args: args{data: []any{}},
			want: "",
		},
		{
			name: "when numbers, strings and booleans",
			args: args{data: []any{true, 1, "One", 2, "Two", 3, "Three", false}},
			want: "OneTwoThree",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			WriteToFile(w, tt.args.data...)
			if got := w.String(); got != tt.want {
				t.Errorf("WriteToFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
