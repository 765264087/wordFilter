package wordFilter

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

type TrieNode struct {
	children map[rune]*TrieNode // 存储子节点
	isEnd    bool               // 标记该节点是否是某个单词的结尾
}

func NewTrieNode() *TrieNode {
	return &TrieNode{children: make(map[rune]*TrieNode)}
}

type Trie struct {
	root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{root: NewTrieNode()}
}

// Push 插入单词
func (t *Trie) Push(word string) {
	currentNode := t.root
	for _, ch := range word {
		if _, ok := currentNode.children[ch]; !ok {
			currentNode.children[ch] = NewTrieNode()
		}
		currentNode = currentNode.children[ch]
	}
	currentNode.isEnd = true
}

// 检查并返回敏感词长度
func (t *Trie) checkSensitiveWord(str string) (matchFlag int) {
	currentNode := t.root
	flag := false
	for _, ch := range str {
		node, ok := currentNode.children[ch]
		if ok {
			matchFlag += len(string(ch))

			if node.isEnd {
				flag = true
			}
		} else {
			break
		}
		currentNode = node
	}
	if matchFlag < 2 || !flag {
		matchFlag = 0
	}

	return
}

// Contains 是否有敏感词
func (t *Trie) Contains(str string) bool {
	if str == "" {
		return false
	}
	for i := 0; i < len(str); i++ {
		if t.checkSensitiveWord(str[i:]) > 0 {
			return true
		}
	}
	return false
}

// GetBadWord 获取敏感词
func (t *Trie) GetBadWord(str string) (w []string) {
	if str == "" {
		return
	}

	for i := 0; i < len(str); i++ {
		len := t.checkSensitiveWord(str[i:])

		if len > 0 {
			w = append(w, string([]rune(str[i:i+len])))
			i = i + len - 1
		}
	}
	return
}

func (t *Trie) Replace(str, replaceChar string) string {
	if str == "" {
		return ""
	}
	badWordList := t.GetBadWord(str)

	for _, v := range badWordList {
		displace := strings.Repeat(replaceChar, utf8.RuneCountInString(v))

		str = strings.Replace(str, v, displace, 1)
	}
	return str
}

func (t *Trie) Mark(str, sTag, eTag string) string {
	if str == "" {
		return ""
	}
	badWordList := unique(t.GetBadWord(str))

	for _, v := range badWordList {
		displace := fmt.Sprintf("%s%s%s", sTag, v, eTag)
		str = strings.Replace(str, v, displace, -1)
	}
	return str
}

func unique(str []string) []string {
	var newStr []string

	for _, s := range str {
		if !in(s, newStr) {
			newStr = append(newStr, s)
		}
	}
	return newStr
}

func in(target string, arr []string) bool {
	for _, value := range arr {
		if value == target {
			return true
		}
	}
	return false
}
