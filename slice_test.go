package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestTrueIndex(t *testing.T) {

	tests := []struct {
		name     string
		capacity int
		start    int
		distance int
		want     int
	}{
		{
			name:     "no wrap",
			capacity: 10,
			start:    0,
			distance: 1,
			want:     1,
		},
		{
			name:     "wrap",
			capacity: 10,
			start:    9,
			distance: 2,
			want:     1,
		},
		{
			name:     "no movement",
			capacity: 10,
			start:    9,
			distance: 0,
			want:     9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSlice(tt.capacity, false)
			ind := s.trueIndex(tt.start, tt.distance)
			require.Equal(t, tt.want, ind)
		})
	}
}
