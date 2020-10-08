package main

import "math"

func main() {

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
