package main

import (
	"fmt"
)

// 剑指 Offer 10- I. 斐波那契数列
// https://leetcode-cn.com/problems/fei-bo-na-qi-shu-lie-lcof/

func main() {
	var res uint64 = 1000000008
	fmt.Println(res % uint64(1e9+7))
	fmt.Println(fib3(8))
}

// 1 1 2 3 5 8 13 21 34 55
// 递推
func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	// 利用切片,n+1,用n当做下标，来获取值
	nums := make([]int, n+1)
	nums[0] = 0
	nums[1] = 1
	for i := 2; i <= n; i++ {
		nums[i] = (nums[i-1] + nums[i-2]) % 1000000007
	}
	return nums[n]
}

func fib2(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	var m1, m2 int
	m1, m2 = 1, 2
	for n > 3 {
		m2, m1 = m1+m2, m2
		m1, m2 = m1%1000000007, m2%1000000007
		n--
	}

	return m2
}

// 递归-超时
func fib3(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	if n == 2 {
		return 1
	}
	return fib3(n-1) + fib3(n-2)
}
