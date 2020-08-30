package leetcode

import "bytes"

/*
557. 反转字符串中的单词 III
给定一个字符串，你需要反转字符串中每个单词的字符顺序，同时仍保留空格和单词的初始顺序。



示例：

输入："Let's take LeetCode contest"
输出："s'teL ekat edoCteeL tsetnoc"


提示：

在字符串中，每个单词由单个空格分隔，并且字符串中不会有任何额外的空格。
*/
func reverseWords(s string) string {
	size := len(s)
	b := bytes.NewBuffer(nil)
	sta := 0
	var reverseWordsRevere = func(end int) {
		for end >= sta {
			b.WriteByte(s[end])
			end--
		}
	}

	for i := range s {
		if s[i] == ' ' {
			reverseWordsRevere(i - 1)
			b.WriteByte(' ')
			sta = i + 1
		}
	}
	reverseWordsRevere(size - 1)
	return b.String()
}