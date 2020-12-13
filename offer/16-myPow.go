package main

import (
	"fmt"
	"math"
)

// 剑指 Offer 16. 数值的整数次方
// https://leetcode-cn.com/problems/shu-zhi-de-zheng-shu-ci-fang-lcof/
func main() {
	fmt.Print(myPow2(2, 6))
}

// 使用二分法来解决
func myPow2(x float64, n int) float64 {
	// 0 ^n =0
	if x == 0 {
		return 0
	}
	// x ^0 =1
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	if n < 0 {
		x = 1 / x
		n = -n
	}

	var res float64 = 1
	// 二分乘法, n=32, 其实就是 n^16 * n^16
	// 使用公式
	// 偶数 (x^2)^ (n^2)
	// 奇数 x * (x^2)^ (n^2)
	for n > 0 {
		// 说明是奇数
		if n&1 == 1 {
			res *= x
		}
		x *= x
		n = n >> 1
	}
	return res
}

func myPow(x float64, n int) float64 {
	// 0 ^n =0
	if x == 0 {
		return 0
	}
	// x ^0 =1
	if n == 0 {
		return 1
	}
	n2 := math.Abs(float64(n))
	res := x
	for i := 2; float64(i) <= n2; i++ {
		res *= x
	}
	// 判断符号
	if n < 0 {
		return float64(1) / res
	}
	return res
}
