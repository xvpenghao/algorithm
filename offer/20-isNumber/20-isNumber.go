package main

import (
	"fmt"
	"strings"
)

func main() {
	// 0.5
	// -1E-16
	fmt.Println(isNumber("0.5.."))
}

// 剑指 Offer 20. 表示数值的字符串
// https://leetcode-cn.com/problems/biao-shi-shu-zhi-de-zi-fu-chuan-lcof/
// 5e2 5的10^2次方
// -1E-16 -1的 10^16次方
// A[.[B]][e|E C] 或者 .B[e|EC]
func isNumber(s string) bool {
	s = strings.TrimSpace(s)
	index := 0
	isNumber := scanInteger(&index, s)
	// .123->0.123 , 123.->123.0, 123.123
	if index < len(s) && s[index] == '.' {
		index++
		// 扫描 小数点后面字符
		// 用 || 的原因
		// 123.->123.0 小数点后面没有数字 导致 scanUnsignedInteger 返回false 实则是正确的
		// .123->0.123 小数点前面没有数字 导致  scanInteger 返回false  实则是正确的
		// 所以用 || 两者一个为true则为true
		// 判断的位置 记得千万不要反了
		isNumber = scanUnsignedInteger(&index, s) || isNumber
	}

	if index < len(s) && (s[index] == 'e' || s[index] == 'E') {
		index++
		// e 或者E的 前后都是 数字才能匹配
		// e3,2e 都不能成功
		isNumber = scanInteger(&index, s) && isNumber
	}

	return isNumber && index >= len(s)
}

// 扫描有符号的
func scanInteger(index *int, s string) bool {
	if *index < len(s) && (s[*index] == '+' || s[*index] == '-') {
		*index++ // 顾虑到 符号
	}
	return scanUnsignedInteger(index, s)
}

// 扫描无符号的
func scanUnsignedInteger(index *int, s string) bool {
	// 这里面有没有数字
	before := *index
	for *index < len(s) && s[*index] >= '0' && s[*index] <= '9' {
		*index++
	}
	// index > before 说明里面确实确实数字
	return *index > before
}
