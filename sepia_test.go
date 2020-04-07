package imageTransform

import (
	"image/color"
	"testing"
)

func TestNormalize(t *testing.T) {
	t.Run("normalize(65536)", func(t *testing.T) { 
		var want uint16 = 65535
		got := normalize(65536);
		if want != got {
			t.Errorf("normalize(65536) = %v; want %v", got, want)
		}
	})
	t.Run("normalize(4444.231)", func(t *testing.T) { 
		var want uint16 = 4444
		got := normalize(4444.231);
		if want != got {
			t.Errorf("normalize(4444.231) = %v; want %v", got, want)
		}
	})
	t.Run("normalize(0)", func(t *testing.T) { 
		var want uint16 = 0
		got := normalize(0);
		if want != got {
			t.Errorf("normalize(0) = %v; want %v", got, want)
		}
	})
}

func TestSepianizeColor(t *testing.T) {
	t.Run("sepianizeColor(40000, 40000, 40000, 3)", func(t *testing.T) { 
		want := color.RGBA64{54040, 48120, 37480, 3}
		got := sepianizeColor(40000, 40000, 40000, 3);
		if want != got {
			t.Errorf("sepianizeColor(40000, 40000, 40000, 3) = %v; want %v", got, want)
		}
	})
	t.Run("sepianizeColor(0, 0, 0, 0)", func(t *testing.T) { 
		want := color.RGBA64{0, 0, 0, 0}
		got := sepianizeColor(0, 0, 0, 0);
		if want != got {
			t.Errorf("sepianizeColor(0, 0, 0, 0) = %v; want %v", got, want)
		}
	})
	t.Run("sepianizeColor(65535, 65535, 65535, 1)", func(t *testing.T) { 
		want := color.RGBA64{65535, 65535, 61406, 1}
		got := sepianizeColor(65535, 65535, 65535, 1);
		if want != got {
			t.Errorf("sepianizeColor(65535, 65535, 65535, 1) = %v; want %v", got, want)
		}
	})
}


