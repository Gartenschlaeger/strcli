package utilities

func ClampI(v, min, max int) int {
	if v < min {
		return min
	}
	if v > max {
		return max
	}

	return v
}

func MinI(a int, b int) int {
	if a < b {
		return a
	}

	return b
}

func MaxI(a int, b int) int {
	if a > b {
		return a
	}

	return b
}

func AbsI(v int) int {
	if v < 0 {
		return -v
	}

	return v
}

func ModI(v int, m int) int {
	result := v % m
	if result < 0 {
		result += m
	}
	return result
}
