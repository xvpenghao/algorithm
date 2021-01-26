package main

import "fmt"

func main() {
	fmt.Println(exchange2([]int{1, 6, 3, 4}))
}

func exchange2(nums []int) []int {
	// 类似 快速排序的交换
	frontIndex, lastIndex := 0, len(nums)-1
	numsLen := len(nums)
	for frontIndex < lastIndex {
		// 左边是奇数
		for frontIndex < numsLen-1 && nums[frontIndex]&0x1 != 0 {
			frontIndex++
		}

		// 右边是偶数
		for lastIndex >= 0 && nums[lastIndex]&0x1 == 0 {
			lastIndex--
		}
		// 如果左边和 右边的位置相等则不用交换
		if frontIndex < lastIndex {
			// 交换
			nums[frontIndex], nums[lastIndex] = nums[lastIndex], nums[frontIndex]
		}
	}
	return nums
}

// https://leetcode-cn.com/problems/diao-zheng-shu-zu-shun-xu-shi-qi-shu-wei-yu-ou-shu-qian-mian-lcof/
// 剑指 Offer 21. 调整数组顺序使奇数位于偶数前面
func exchange(nums []int) []int {
	// 数组中可能都是无序的
	nums2 := make([]int, len(nums))
	// 1 2 3 4
	frontIndex, tailIndex := 0, len(nums)-1
	for _, val := range nums {
		// 不是偶数
		if val%2 != 0 {
			nums2[frontIndex] = val
			frontIndex++
			continue
		}
		nums2[tailIndex] = val
		tailIndex--
	}

	return nums2
}
