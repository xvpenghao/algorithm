package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(cuttingRope2(10))
}

// 剑指 Offer 14- I. 剪绳子
// https://leetcode-cn.com/problems/jian-sheng-zi-lcof/
func cuttingRope(n int) int {
	if n <= 3 {
		return n - 1
	}
	count := n / 3
	if n%3 == 0 {
		return int(math.Pow(3, float64(count)))
	} else if n%3 == 1 {
		return int(math.Pow(3, float64(count-1))) * 4
	}

	return int(math.Pow(3, float64(count))) * 2
}

// 动态规划
func cuttingRope2(n int) int {
	if n <= 3 {
		return n - 1
	}

	nums := make([]int, n+1)
	nums[0], nums[1], nums[2], nums[3] = 0, 1, 2, 3

	// f(n) = max(f(i)*f(n-i))
	// i 表示每一段绳子的长度
	for i := 4; i <= n; i++ {
		max := 0
		// 对绳子进行切割，计算每一段绳子的最大值
		for j := 1; j <= i/2; j++ {
			if t := nums[j] * nums[i-j]; t > max {
				max = t
			}
		}
		nums[i] = max
	}
	return nums[n]
}
