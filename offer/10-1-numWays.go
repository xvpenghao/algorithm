package main

import "fmt"

func main() {
	fmt.Println(numWays(7))
}

//剑指 Offer 10- II. 青蛙跳台阶问题
//https://leetcode-cn.com/problems/qing-wa-tiao-tai-jie-wen-ti-lcof/
func numWays(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	m1, m2 := 1, 2
	for n > 2 {
		m1, m2 = m2%1000000007, (m1+m2)%1000000007
		n--
	}
	return m2
}
