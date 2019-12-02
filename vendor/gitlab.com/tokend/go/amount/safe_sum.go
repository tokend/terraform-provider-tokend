package amount

// Returns true if ok
func SafePositiveSum(a, b int64) (int64, bool) {
	if a < 0 || b < 0 {
		return 0, false
	}

	result := a + b
	return result, result >= a && result >= b
}

func SafePositiveSumForSlice(nums ...int64) (int64, bool) {
	var result int64
	for _, n := range nums {
		if n < 0 {
			return 0, false
		}

		if result < 0 {
			return 0, false
		}

		result += n
	}

	for _, n := range nums {
		if result < n {
			return 0, false
		}
	}

	return result, true
}
