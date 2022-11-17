package fun

import "testing"

func TestSimilarity(t *testing.T) {
	t.Log(Similarity("", ""))
	t.Log(Similarity("ABC", "B"))
	t.Log(Similarity("ABC", "AC"))
	t.Log(Similarity("ABC", "CA"))
	t.Log(Similarity("ABCDEF", "CD"))
	t.Log(Similarity("ABCDEFGHIG", "CEGML"))
	t.Log(Similarity("ABCDEFGHIG", "GIH"))
	t.Log(Similarity("我是中国人我说中国话", "我米说"))

	t.Log(Similarity("国家统计局上半年GDP同比增长5", "上半年GDP同比增长"))
}

func TestLcs(t *testing.T) {
	t.Log(LongestCommonSubString("ABC", "B"))
	t.Log(LongestCommonSubString("ABC", "AC"))
	t.Log(LongestCommonSubString("ABC", "CA"))
	t.Log(LongestCommonSubString("ABCDEF", "CD"))
	t.Log(LongestCommonSubString("ABCDEFGHIG", "CEGML"))
	t.Log(LongestCommonSubString("ABCDEFGHIG", "GIH"))
	t.Log(LongestCommonSubString("我是中国人我说中国话", "中国人米说"))
}

func TestSimilarityText(t *testing.T) {
	t.Log(SimilarityText("国家统计局：上半年GDP同比增长5%", "上半年GDP同比增长"))
}

func TestReSimilarityText(t *testing.T) {
	t.Log(SimilarityText("人民日报仲音：算大账 看优势", "算大账 看优势"))
	t.Log(SimilarityText("人民日报仲音：算大账 看优势", " 看优势"))
}
