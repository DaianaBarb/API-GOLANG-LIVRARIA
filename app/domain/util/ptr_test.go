package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPtr_BoolToPtr(t *testing.T) {
	tBool := true
	fBool := false

	type test struct {
		name string
		args bool
		want *bool
	}
	tests := []test{
		{
			name: "when value is true, should return a bool pointer with value true",
			args: true,
			want: &tBool,
		},
		{
			name: "when value is false, should return a bool pointer with value false",
			args: false,
			want: &fBool,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BoolToPtr(tt.args)
			assert.Equalf(t, tt.want, got, "BoolToPtr(...) = %v, want = %v", got, tt.want)
		})
	}
}
