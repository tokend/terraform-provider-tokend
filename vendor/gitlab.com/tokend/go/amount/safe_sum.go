package amount

// Returns true if ok
func SafePositiveSum(a, b int64) (int64, bool) {
	if a < 0 || b < 0 {
		return 0, false
	}

	result := a + b
	return result, result >= a && result >= b
}
