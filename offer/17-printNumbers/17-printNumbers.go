package main

import (
	"fmt"
	"math"
	"strings"
)

// 剑指 Offer 17. 打印从1到最大的n位数
// https://leetcode-cn.com/problems/da-yin-cong-1dao-zui-da-de-nwei-shu-lcof/
func main() {
	// res := printNumbers(3)
	// fmt.Println(res)
	printMaxNumbers(2)
}

// 1 max 9
// 2 max 99
// 3 max 999
// n max (10^n)-1
// 1,2,3,4,5,6,7,8,9,10
// 11,12,13,14,15,16,17,18,19,20
// 21,22,23,24,25,26,27,28,29,30
func printNumbers(n int) []int {
	bases := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	if n == 1 {
		return bases[:len(bases)-1]
	}

	num := math.Pow(10, float64(n)) / 10
	for i := 2; i <= int(num); i++ {
		length := len(bases)
		tmp := []int{
			bases[length-10] + 10,
			bases[length-9] + 10,
			bases[length-8] + 10,
			bases[length-7] + 10,
			bases[length-6] + 10,
			bases[length-5] + 10,
			bases[length-4] + 10,
			bases[length-3] + 10,
			bases[length-2] + 10,
			bases[length-1] + 10,
		}
		bases = append(bases, tmp...)
	}
	return bases[:len(bases)-1]
}

// 考虑大数的问题
// https://learnku.com/articles/48282
// 代码在优化一下
func printMaxNumbers(n int) {
	// 固定高位，然后低位
	loops := []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	nums := make([]byte, n) // 定义长度为n 的字符列表
	sb := new(strings.Builder)
	// 消除 多于0   n=3 009 099 999
	start := n - 1 // 截取的小标
	nine := 0      // 用于记录9的个数
	dfs2(n, 0, &start, &nine, sb, loops, nums)
	// 去掉 最后一位的，
	res := strings.TrimSuffix(sb.String(), ",")
	fmt.Println(res)
}

// 深度优先的查找
func dfs2(n, x int, start, nine *int, sb *strings.Builder, loops, nums []byte) {
	if n == x {
		s := string(nums[*start:]) // 根据 start 截取合适的 位置 009 099 999
		if !strings.EqualFold(s, "0") {
			sb.WriteString(s + ",")
		}
		// n=3，009 099， 999
		// n-start 为9的个数
		if n-(*start) == *nine {
			*start--
		}
		return
	}

	// 给固定位置添加数据
	for _, val := range loops {
		nums[x] = val // 固定第 x 位 0表示第一位, 固定位置从 高位-低位
		// 外层9 到 内层9的个数
		if val == '9' {
			*nine++
		}
		dfs2(n, x+1, start, nine, sb, loops, nums)
	}
	// 内层执行完毕后，在复原
	*nine--
}
