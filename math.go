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

// MaxN 取 N 个数字的最大值
func MaxN[T GenNumber](args ...T) T {
	max := args[0]
	for _, arg := range args {
		if arg > max {
			max = arg
		}
	}

	return max
}

// MinN 取 N 个数字的最小值
func MinN[T GenNumber](args ...T) T {
	min := args[0]
	for _, arg := range args {
		if arg < min {
			min = arg
		}
	}

	return min
}
