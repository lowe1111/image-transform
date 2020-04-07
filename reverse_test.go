package imageTransform

import (
	"image/color"
	"testing"
)

func TestReversColors(t *testing.T) {
	t.Run("reverseColors(255,255,255,23)", func(t *testing.T) { 
		got := reverseColors(255,255,255,23)
		want := color.RGBA64{65280, 65280, 65280, 23}
		if got != want {
			t.Errorf("reverseColors(255,255,255,23) = %v; want %v", got, want)
		}
	})
	t.Run("reverseColors(32334,23000,1000, 3)", func(t *testing.T) { 
		got := reverseColors(32334,23000,1000, 3)
		want := color.RGBA64{33201, 42535, 64535, 3}
		if got != want {
			t.Errorf("reverseColors(32334,23000,1000, 3) = %v; want %v", got, want)
		}
	})
	t.Run("reverseColors(0,0,0,0)", func(t *testing.T) { 
		got := reverseColors(0,0,0, 0)
		want := color.RGBA64{65535, 65535, 65535, 0}
		if got != want {
			t.Errorf("reverseColors(0,0,0,0) = %v; want %v", got, want)
		}
	})
}