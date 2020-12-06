package main

import "fmt"

func main() {
	var num uint32 = 00000000000000000000000000001011
	fmt.Println(hammingWeight(num), num)
}

// 剑指 Offer 15. 二进制中1的个数
// https://leetcode-cn.com/problems/er-jin-zhi-zhong-1de-ge-shu-lcof/
func hammingWeight(num uint32) int {
	count := 0
	for num != 0 {
		count++
		// 二级制 1011
		// 1011-1 = 1100
		// 1100 & 1011 = 1000 操作完毕后每次 二进制中最右边的1都会变为0
		num = (num - 1) & num
	}

	return count
}
