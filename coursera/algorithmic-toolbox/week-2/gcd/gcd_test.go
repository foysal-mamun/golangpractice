package gcd

import "testing"

func TestNativeAlgo(t *testing.T) {
	r := NativeAlgo(3918848, 1653264)
	var ar int64 = 61232
	if r != ar {
		t.Errorf("Wrong Result %d, Correct Result is %d", r, ar)
	}
}

func BenchmarkNativeAlgo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NativeAlgo(3918848, 1653264)
	}
}

func TestEffectiveAlgo(t *testing.T) {
	r := EffectiveAlgo(3918848, 1653264)
	var ar int64 = 61232
	if r != ar {
		t.Errorf("Wrong Result %d, Correct Result is %d", r, ar)
	}
}

func BenchmarkEffectiveAlgo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EffectiveAlgo(3918848, 1653264)
	}
}
