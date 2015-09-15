package fibonacci

// Getnth simply returns the number at that position in the
// fibonacci sequence
func Getnth(n int) int64 {
	var fib int64 = 1
	var prev int64 = 0
	if n == 1 {
		return 0
	}

	for i := 1; i < n; i++ {
		next := fib + prev
		prev = fib
		fib = next
	}
	return fib
}
