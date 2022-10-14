// Package tire TireTree 实现
package tire

import "bytes"

type tire struct {
	data map[rune]*tire
	end  bool
}

type Option struct {
	// 匹配次数限制, -1 不限制
	Limit int
	// 贪婪模式, 尽可能的多匹配词语. 如关键词定义 ["上海", "上海游玩"], 对于句子 "他到上海游玩". true 则匹配 ["上海", "上海游玩"], false 则只会匹配 ["上海"]
	Greed bool
	// 密度, 匹配出词中词. 如关键词定义 ["到上海", "上海"], 对于句子 "他到上海游玩". true 则匹配 ["到上海", "上海"], false 则只会匹配 ["到上海"]
	Density bool
	// 是否是单词语系
	WordMode bool
}

// NewTire new TireTree
func NewTire() *tire {
	return &tire{
		data: nil,
		end:  false,
	}
}

// AddAll 批量添加词语
func (t *tire) AddAll(words []string) *tire {
	if len(words) > 0 {
		for _, word := range words {
			t.Add(word)
		}
	}

	return t
}

// Add 添加词语
func (t *tire) Add(word string) *tire {
	if word == "" {
		return t
	}

	var child *tire = nil
	var current = t

	charList := []rune(word)
	length := len(charList)
	for i := 0; i < length; i++ {
		char := charList[i]
		child = current.getChild(char)
		if child == nil {
			// 无子类，新建一个子节点后存放下一个字符
			child = NewTire()
			current.addChild(char, child)
		}

		current = child
	}

	current.setEnd()

	return t
}

// Contains 是否包含
func (t *tire) Contains(text string, wordMode bool) bool {
	word := t.FindWithOptions(text, Option{
		Limit:    1,
		Greed:    false,
		Density:  false,
		WordMode: wordMode,
	})
	return len(word) != 0
}

// Find 包含词语列表
func (t *tire) Find(text string, wordMode bool) map[string]int {
	word := t.FindWithOptions(text, Option{
		Limit:    -1,
		Greed:    false,
		Density:  false,
		WordMode: wordMode,
	})
	return word
}

// FindWithOptions 包含词语列表, 并统计次数
func (t *tire) FindWithOptions(text string, opt Option) map[string]int {
	var foundWordList = make([]string, 0)
	var curNode *tire

	var word bytes.Buffer
	charList := []rune(text)
	length := len(charList)
	for i := 0; i < length; i++ {
		word.Reset()
		curNode = t
		for j := i; j < length; j++ {
			char := charList[j]

			curNode = curNode.getChild(char)
			if curNode == nil {
				break
			}

			// 关键词是否是全量字母
			// 若关键词是全量字母且在它之前的字符是字母, 则该词无需被记录
			if opt.WordMode && t.isSeparator(char) && i > 0 && !t.isSeparator(charList[i-1]) {
				break
			}

			word.WriteRune(char)

			if curNode.isEnd() {
				if opt.WordMode && j < length-1 && !t.isSeparator(charList[j+1]) {
					break
				}

				foundWordList = append(foundWordList, word.String())

				if opt.Limit > 0 && len(foundWordList) >= opt.Limit {
					return t.statFoundWord(foundWordList)
				}

				if !opt.Density {
					i = j
					break
				}

				if !opt.Greed {
					break
				}
			}
		}
	}

	return t.statFoundWord(foundWordList)
}

// 获取子节点
func (t *tire) getChild(char rune) *tire {
	tire, ok := t.data[char]
	if ok {
		return tire
	}
	return nil
}

// 添加子节点
func (t *tire) addChild(char rune, child *tire) {
	if t.data == nil {
		t.data = make(map[rune]*tire, 0)
	}
	t.data[char] = child
}

// 设置结束
func (t *tire) setEnd() {
	t.end = true
}

// 是否结束
func (t *tire) isEnd() bool {
	return t.end
}

// 是否是分隔符
func (t *tire) isSeparator(c rune) bool {
	// 32( ), 33(!), 34("), 35(#), 37(%), 38(&), 39('), 40((), 41()), 42(*), 42(*), 44(,), 45(-), 46(.), 47(/), 58(:)
	// 59(;), 60(<), 61(=), 62(>), 64(@), 91([), 93(]), 96(`), 123({), 124(|), 125(}), 126(~), 183(·), 8216( ‘)
	// 8217(’), 8220(“), 8221(”), 8230(…), 12289(、), 12290(。), 12298(《), 12299(》), 12304(【), 12305(】), 65281(！)
	// 65288(（), 65289(）), 65292(，), 65306(：), 65307(；)
	return c == 32 || c == 33 || c == 34 || c == 35 || c == 37 || c == 38 || c == 39 || c == 40 || c == 41 || c == 42 || c == 44 || c == 45 || c == 46 || c == 47 || c == 58 || c == 59 || c == 60 || c == 61 || c == 62 || c == 64 || c == 91 || c == 93 || c == 96 || c == 123 || c == 124 || c == 125 || c == 126 || c == 183 || c == 8216 || c == 8217 || c == 8220 || c == 8221 || c == 8230 || c == 12289 || c == 12290 || c == 12298 || c == 12299 || c == 12304 || c == 12305 || c == 65281 || c == 65288 || c == 65289 || c == 65292 || c == 65306 || c == 65307
}

// 词数统计
func (t *tire) statFoundWord(list []string) map[string]int {
	foundWordMap := make(map[string]int, 0)
	for _, word := range list {
		wordStr := word
		if _, ok := foundWordMap[wordStr]; !ok {
			foundWordMap[wordStr] = 1
		} else {
			foundWordMap[wordStr] += 1
		}
	}
	return foundWordMap
}
