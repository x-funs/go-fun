package tree

import (
	"bytes"
)

type Tire struct {
	data map[rune]*Tire
	end  bool
}

type Opt struct {
	Limit    int  // 限制, 匹配到多少个不一样的词语后结束匹配
	Greed    bool // 贪婪, 尽可能的多匹配词语. 如关键词定义 ["上海", "上海游玩"], 对于句子 "他到上海游玩". true 则匹配 ["上海", "上海游玩"], false 则只会匹配 ["上海"]
	Density  bool // 密度, 匹配出词中词. 如关键词定义 ["到上海", "上海"], 对于句子 "他到上海游玩". true 则匹配 ["到上海", "上海"], false 则只会匹配 ["到上海"]
	HasGroup bool // 是否有词组, 建议当文章有英文开启, 否则会影响匹配的效率
}

// Add 添加词语
func (t *Tire) Add(word string) *Tire {
	if word == "" {
		return t
	}

	var child *Tire = nil
	var current = t

	charList := []rune(word)
	length := len(charList)
	for i := 0; i < length; i++ {
		char := charList[i]
		child = current.getChild(char)
		if child == nil {
			// 无子类，新建一个子节点后存放下一个字符
			child = new(Tire)
			current.addChild(char, child)
		}

		current = child
	}

	current.setEnd()

	return t
}

// 获取子节点
func (t *Tire) getChild(char rune) *Tire {
	tire, ok := t.data[char]
	if ok {
		return tire
	}
	return nil
}

// 添加子节点
func (t *Tire) addChild(char rune, child *Tire) {
	if t.data == nil {
		t.data = make(map[rune]*Tire, 0)
	}
	t.data[char] = child
}

// 设置结束
func (t *Tire) setEnd() {
	t.end = true
}

// 是否结束
func (t *Tire) isEnd() bool {
	return t.end
}

// Contains 是否包含词语
func (t *Tire) Contains(text string) bool {
	word := t.FindAll(text, Opt{
		Limit:    1,
		Greed:    false,
		Density:  false,
		HasGroup: true,
	})
	return len(word) != 0
}

// FindAll 匹配全部, 返回(匹配词 => 出现次数)的映射
func (t *Tire) FindAll(text string, opt Opt) map[string]int {
	var foundWordList = make([]string, 0)
	var curNode *Tire

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
			if opt.HasGroup && t.isSeparator(char) && i > 0 && !t.isSeparator(charList[i-1]) {
				break
			}

			word.WriteRune(char)

			if curNode.isEnd() {
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

// 是否是分隔符, 目前取 ASCII 中的标点符号
func (t *Tire) isSeparator(char rune) bool {
	return char < 65 || (char > 90 && char < 97) || char > 122
}

// 词数统计
func (t *Tire) statFoundWord(list []string) map[string]int {
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
