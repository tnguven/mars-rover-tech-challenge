package plateau_test

import (
   "testing"

   "github.com/stretchr/testify/assert"
   "github.com/tnguven/mars-rover-challenge/internal/plateau"
)

func TestNewPlateau(t *testing.T) {
   p := plateau.New(5, 5)

   assert.NotNil(t, p)
   assert.Equal(t, 5, p.MaxX)
   assert.Equal(t, 5, p.MaxY)
}

func TestIsInBounds(t *testing.T) {
   p := plateau.New(5, 5)

   tests := []struct {
      x, y     int
      expected bool
   }{
      {0, 0, true},
      {5, 5, true},
      {3, 2, true},
      {-1, 0, false},
      {0, -1, false},
      {6, 5, false},
      {5, 6, false},
   }

   for _, tt := range tests {
      assert.Equal(t, tt.expected, p.IsinBounds(tt.x, tt.y), "x=%d y=%d", tt.x, tt.y)
   }
}
