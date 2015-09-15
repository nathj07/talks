package fibonacci

import "testing"

type testData struct {
	input    int
	expected int64
}

var testSet = []testData{
	testData{
		input:    1,
		expected: 0,
	},
	testData{
		input:    2,
		expected: 1,
	},
	testData{
		input:    15,
		expected: 610,
	},
	testData{
		input:    36,
		expected: 14930352,
	},
}

func TestGetnth(t *testing.T) {
	for i, d := range testSet {
		got := Getnth(d.input)
		if got != d.expected {
			t.Errorf("Test Case: %d Expected: %d; Got: %d", i, d.expected, got)
		}
	}
}

func BenchmarkGetnth(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Getnth(i)
	}
}
