package fibonacci

import "testing"

func TestGetnth(t *testing.T) {
	var testSet = []struct {
		input    int
		expected int64
	}{
		{1, 0},
		{2, 1},
		{15, 610},
		{36, 14930352},
	}
	for i, d := range testSet {
		got := Getnth(d.input)
		if got != d.expected {
			t.Errorf("Test Case: %d Expected: %d; Got: %d", i, d.expected, got)
		}
	}
}

func BenchmarkGetnthDepth(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Getnth(i)
	}
}

func BenchmarkGetnthSingle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Getnth(38)
	}
}
