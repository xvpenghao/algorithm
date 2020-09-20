package main

import "fmt"

func main() {
	// fmt.Println(getDigitSum(10))
	//
	// 10 6
	fmt.Println(movingCount(38, 15, 9))
	/*
	 */
}

// 剑指 Offer 13. 机器人的运动范围
// https://leetcode-cn.com/problems/ji-qi-ren-de-yun-dong-fan-wei-lcof/
func movingCount(m int, n int, k int) int {
	if k == 0 {
		return 1
	}

	count := 1
	vis := getVisBool(m, n)
	vis[0][0] = true
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// vis[0][0] 已经算了，所以这里要略掉
			if (i == 0 && j == 0) || (getDigitSum(i)+getDigitSum(j) > k) {
				continue
			}

			// 计算是否连续
			if (i - 1) >= 0 {
				vis[i][j] = vis[i][j] || vis[i-1][j] // 是否是连续的走过
			}

			if (j - 1) >= 0 {
				vis[i][j] = vis[i][j] || vis[i][j-1] // 是否是连续的走过
			}

			if vis[i][j] {
				count++
			}
		}
	}
	return count
}

func getVisBool(m, n int) [][]bool {
	var res [][]bool
	for i := 0; i < m; i++ {
		res = append(res, make([]bool, n))
	}
	return res
}

func getDigitSum(x int) int {
	sum := 0
	for x > 0 {
		sum += x % 10
		x /= 10
	}
	return sum
}
