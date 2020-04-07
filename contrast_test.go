package imageTransform

import (
	"testing"
)

func TestHigherContrast(t *testing.T) {
	t.Run("higherContrast(2, 5)", func(t *testing.T) { 
		got := higherContrast(2, 5)
		if got != 0.4 {
			t.Errorf("higherContrast(2, 5) = %f; want 1", got)
		}
	})
	t.Run("higherContrast(55535, 2)", func(t *testing.T) { 
		got := higherContrast(55535, 2)
		if got != 65535 {
			t.Errorf("higherContrast(55535, 2) = %f; want 1", got)
		}
	})
	t.Run("higherContrast(0, 2)", func(t *testing.T) { 
		got := higherContrast(0, 2)
		if got != 0 {
			t.Errorf("higherContrast(0, 3) = %f; want 1", got)
		}
	})
}