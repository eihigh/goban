package goban

import (
	"reflect"
	"testing"
)

func TestFit(t *testing.T) {
	a := NewBox(0, 0, 10, 10)
	b := NewBox(10, 10, 10, 10)

	tests := []struct {
		got, want *Box
	}{
		{a.Fit(b, -1, -1, -1, -1), NewBox(10, 10, 10, 10)},
		{a.Fit(b, -1, -1, 0, 0), NewBox(5, 5, 10, 10)},
		{a.Fit(b, 0, 0, 0, 0), NewBox(10, 10, 10, 10)},
		{a.Fit(b, 1, 1, -1, -1), NewBox(20, 20, 10, 10)},
	}

	for _, tt := range tests {
		if !reflect.DeepEqual(tt.want, tt.got) {
			t.Errorf("got %v; want %v", tt.got, tt.want)
		}
	}
}
