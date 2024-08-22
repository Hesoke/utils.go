package rng

import (
	"testing"
)

func TestPiBased(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"test 1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log(PiBased().Int63())
		})
	}
}
