package wordFilter

import (
	"bufio"
	"os"
	"testing"
)

func Init() *Trie {
	trie := NewTrie()
	file, err := os.Open("./test/words.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		trie.Push(line)
	}
	return trie
}

func TestContains(t *testing.T) {
	trie := Init()

	if !trie.Contains("法大东京热大电视") {
		t.Errorf("期望：%t,实际结果：%t", true, false)
	}

	if trie.Contains("ssssh东热ossss") {
		t.Errorf("期望：%t,实际结果：%t", false, true)
	}
}

func TestGetBadWord(t *testing.T) {
	trie := Init()

	w := []string{"东京热"}
	res := trie.GetBadWord("法大东京热大电视")

	if !Equal(w, res) {
		t.Errorf("期望：%v,实际结果：%v", w, res)
	}

	w = []string{"东京热", "法伦功"}
	res = trie.GetBadWord("法大东京热大电法伦功视")

	if !Equal(w, res) {
		t.Errorf("期望：%v,实际结果：%v", w, res)
	}
}

func TestReplace(t *testing.T) {
	trie := Init()

	w := "法大***大电视"
	res := trie.Replace("法大东京热大电视", "*")

	if w != res {
		t.Errorf("期望：%v,实际结果：%v", w, res)
	}

	w = "法大***大电***视"
	res = trie.Replace("法大东京热大电法伦功视", "*")

	if w != res {
		t.Errorf("期望：%v,实际结果：%v", w, res)
	}
}

func TestMark(t *testing.T) {
	trie := Init()

	w := "法大<span>东京热</span>大电视"
	res := trie.Mark("法大东京热大电视", "<span>", "</span>")

	if w != res {
		t.Errorf("期望：%v,实际结果：%v", w, res)
	}

	w = "sssshellossssworld"
	res = trie.Mark("sssshellossssworld", "<span>", "</span>")

	if w != res {
		t.Errorf("期望：%v,实际结果：%v", w, res)
	}
}

func Equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
