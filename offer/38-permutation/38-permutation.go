package main

import "fmt"

/**
*@Author:徐鹏豪
*@Time: 2021/3/7 16:18
*@Description:
 */
func main() {
	result := permutation("1234")
	fmt.Println(result, len(result))
}

// 剑指 Offer 38. 字符串的排列
/*
abc
acb
bac
-----
bca
cab
cba


*/
// https://leetcode-cn.com/problems/zi-fu-chuan-de-pai-lie-lcof/
func permutation(s string) []string {
	if s == "" {
		return nil
	}
	dataByte := []byte(s)

	var dfs func(x int)
	var result []string
	dfs = func(x int) {
		if x == len(dataByte)-1 {
			result = append(result, string(dataByte))
			return
		}
		set := make(map[byte]bool)
		for i := x; i < len(dataByte); i++ {
			if _, ok := set[dataByte[i]]; ok {
				continue
			}
			set[dataByte[i]] = true
			// swap
			dataByte[i], dataByte[x] = dataByte[x], dataByte[i] // 交换
			dfs(x + 1)
			dataByte[i], dataByte[x] = dataByte[x], dataByte[i] // 恢复交换
		}
	}

	dfs(0)
	return result
}
