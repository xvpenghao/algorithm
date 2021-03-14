package main

import (
	"fmt"
	"sort"
)

/**
*@Author:徐鹏豪
*@Time: 2021/3/13 21:21
*@Description:
 */
func main() {
	nums := []int{1, 2, 3, 2, 2, 2, 5, 4, 2}
	fmt.Println(majorityElement3(nums))
}

// 剑指 Offer 39. 数组中出现次数超过一半的数字
// https://leetcode-cn.com/problems/shu-zu-zhong-chu-xian-ci-shu-chao-guo-yi-ban-de-shu-zi-lcof/
// 摩尔投票法
func majorityElement3(nums []int) int {
	var votes, x int
	for _, num := range nums {
		if votes == 0 {
			x = num
		}
		// 使用相互抵消的法则，你用 你需要的 数 是超过 数组长度的一半
		if x == num {
			votes += 1
		} else {
			votes += -1
		}
	}
	return x
}

// 排序法，去众数
func majorityElement2(nums []int) int {
	sort.SliceStable(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	halfNumsLen := len(nums) / 2
	return nums[halfNumsLen]
}

// map方法
func majorityElement(nums []int) int {
	countMap := make(map[int]int)
	// 计算每个数字查询的次数
	for _, n := range nums {
		countMap[n] = countMap[n] + 1
	}
	halfNumsLen := len(nums) / 2

	for n, count := range countMap {
		if count > halfNumsLen {
			return n
		}
	}
	return -1
}

// 1, 2, 3, 2, 2, 2, 5, 4, 2
// 1, 2, 2, 2, 2, 2 ,3,4, 5,
// 1 3 3 3
