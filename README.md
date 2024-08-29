# wordFilter
go实现基于前缀数敏感词过滤

### 安装扩展

    go get github.com/765264087/wordFilter

####  构建敏感词库树
    trie := NewTrie()
    trie.Push("察象蚂")
    trie.Push("拆迁灭")
    trie.Push("车牌隐")
    trie.Push("成人电")
    trie.Push("成人卡通")

### 检测是否含有敏感词

    islegal = trie.Contains(str);

### 敏感词过滤

    // 敏感词替换为*为例（会替换为相同字符长度的*）
    filterContent = trie.Replace(str, '*');

### 标记敏感词
     markedContent =  trie.Mark(str, '<mark>', '</mark>');

### 获取文字中的敏感词

    // 获取内容中所有的敏感词
    sensitiveWordGroup = trie.GetBadWord(str);