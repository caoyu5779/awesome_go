package ladderLength

func LadderLength(beginWord string, endWord string, words []string) int {
	dict := make(map[string]bool, len(words))

	for i := 0; i < len(words); i++ {
		dict[words[i]] = true
	}

	delete(dict, beginWord)
	queue := make([]string, 0, len(words))

	//这里可以抽象出去，当然也可以不
	var trans func(string) bool

	trans = func(word string) bool {
		//将字符串转为byte
		bytes := []byte(word)

		for i := 0; i < len(bytes); i++ {
			//每个字幕都比较
			diffLetter := bytes[i]
			// 26个字幕逐一比较
			for j := 0; j < 26; j++ {
				b := 'a' + byte(j)
				//相同就跳出
				if b == diffLetter {
					continue
				}
				//不同的字符
				bytes[i] = b
				if dict[string(bytes)] {
					//检查下 这个不同的单词，是不是终点
					if string(bytes) == endWord {
						return true
					}
					//不是就吧这个单词放进 队列里面
					queue = append(queue, string(bytes))
					//然后把这个单词从字典中删除
					delete(dict, string(bytes))
				}
			}

			bytes[i] = diffLetter
		}
		return false
	}

	queue = append(queue, beginWord)
	dist := 1

	for len(queue) > 0 {
		qLen := len(queue)
		//逻辑从这里开始
		for i := 0; i < qLen; i++ {
			word := queue[0]
			queue = queue[1:]
			//调用的函数
			if trans(word) {
				return dist + 1
			}
		}

		dist++
	}
	return 0
}
