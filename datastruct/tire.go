package datastruct

import (
	"strings"

	"github.com/x-funs/go-fun"
)

type Tire struct {
	data map[string]*Tire
	end  bool
}

type Opt struct {
	Limit   int  // 限制, 匹配到多少个不一样的词语后结束匹配
	Greed   bool // 贪婪, 尽可能的多匹配词语. 如关键词定义 ["上海", "上海游玩"], 对于句子 "他到上海游玩". true 则匹配 ["上海", "上海游玩"], false 则只会匹配 ["上海"]
	Density bool // 密度, 匹配出词中词. 如关键词定义 ["到上海", "上海"], 对于句子 "他到上海游玩". true 则匹配 ["到上海", "上海"], false 则只会匹配 ["到上海"]
}

// Add 添加词语
func (t *Tire) Add(word string) *Tire {
	if word == "" {
		return t
	}

	var child *Tire = nil
	var current = t

	charList := strings.Split(word, "")
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
func (t *Tire) getChild(char string) *Tire {
	tire, ok := t.data[char]
	if ok {
		return tire
	}
	return nil
}

// 添加子节点
func (t *Tire) addChild(char string, child *Tire) {
	if t.data == nil {
		t.data = make(map[string]*Tire, 0)
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
		Limit:   1,
		Greed:   false,
		Density: false,
	})
	return len(word) != 0
}

// FindAll 匹配全部, 返回(匹配词 => 出现次数)的映射
func (t *Tire) FindAll(text string, opt Opt) map[string]int {
	var foundWordMap = make(map[string]int, 0)

	word := ""
	curNode := t
	charList := strings.Split(text, "")
	length := len(charList)
	for i := 0; i < length; i++ {
		word = ""
		curNode = t
		for j := i; j < length; j++ {
			char := charList[j]

			curNode = curNode.getChild(char)
			if curNode == nil {
				break
			}

			// 关键词是否是全量字母
			// 若关键词是全量字母且在它之前的字符是字母, 则该词无需被记录
			if !fun.IsASCIILetter(char) && i > 0 && fun.IsASCIILetter(charList[i-1]) {
				break
			}

			word += char

			if curNode.isEnd() {
				if _, ok := foundWordMap[word]; !ok {
					foundWordMap[word] = 1
				} else {
					foundWordMap[word] += 1
				}

				if opt.Limit > 0 && len(foundWordMap) >= opt.Limit {
					return foundWordMap
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

	return foundWordMap
}
