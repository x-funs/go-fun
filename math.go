package fun

// Max 取 int 最大值
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Min 取 int 最小值
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// MaxInt64 取 int64 最大值
func MaxInt64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

// MinInt64 取 int64 最小值
func MinInt64(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}
