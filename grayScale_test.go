package imageTransform

import (
	"image/color"
	"testing"
)

func TestConvertToGray(t *testing.T) {
	t.Run("convertToGray(255,255,255)", func(t *testing.T) { 
		got := convertToGray(255, 255, 255)
		want := color.Gray16{255}
		if got != want {
			t.Errorf("convertToGray(255,255,255) = %v; want {255}", got)
		}
	})
	t.Run("convertToGray(0,0,0)", func(t *testing.T) { 
		got := convertToGray(0, 0, 0)
		want := color.Gray16{0}
		if got != want {
			t.Errorf("convertToGray(0,0,0) = %v; want {0}}", got)
		}
	})
	t.Run("convertToGray(255,255,255)", func(t *testing.T) { 
		got := convertToGray(23, 222, 87)
		want := color.Gray16{170}
		if got != want {
			t.Errorf("convertToGray(23, 222, 87) = %v; want 170", got)
		}
	})
}