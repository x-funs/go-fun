package fun

// Similarity 计算两个原始字符串的相似度
func Similarity(a, b string) float64 {
	len1 := len([]rune(a))
	len2 := len([]rune(b))

	if len1 == 0 || len2 == 0 {
		return 1
	}

	lcs := float64(LongestCommonSubString(a, b))
	max := float64(Max(len1, len2))

	return lcs / max
}

// SimilarityText 计算两个字符串移除特殊符号后的相似度
func SimilarityText(a, b string) float64 {
	a = RemoveSign(a)
	b = RemoveSign(b)

	return Similarity(a, b)
}

// LongestCommonSubString 计算两个字符串最大公共子串长度
func LongestCommonSubString(x, y string) int {
	rm := []rune(x)
	rn := []rune(y)

	m := len(rm)
	n := len(rn)

	if m == 0 || n == 0 {
		return 0
	}

	// 初始化二维数组
	var opt = make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		opt[i] = make([]int, n+1)
	}

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if rm[i] == rn[j] {
				opt[i][j] = opt[i+1][j+1] + 1
			} else {
				opt[i][j] = Max(opt[i+1][j], opt[i][j+1])
			}
		}
	}

	return opt[0][0]
}
