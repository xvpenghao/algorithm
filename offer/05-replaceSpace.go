package main

import "unicode"

// 减少了内存的消耗，s+=""，这样会频繁的创建字符串对象，分配内存
func replaceSpace2(s string) string {
	// %20
	strRunes := make([]rune, len(s)*3)
	var index int
	for _, v := range s {
		if string(v) != " " {
			strRunes[index] = v
			index++
			continue
		}
		strRunes[index] = '%'
		index++
		strRunes[index] = '2'
		index++
		strRunes[index] = '0'
		index++
	}

	return string(strRunes[:index])
}

// 剑指 Offer 05. 替换空格
// https://leetcode-cn.com/problems/ti-huan-kong-ge-lcof/
func replaceSpace(s string) string {
	var newStr string
	for _, v := range s {
		if unicode.IsSpace(v) {
			newStr += "%20"
			continue
		}
		newStr += string(v)
	}
	return newStr
}
