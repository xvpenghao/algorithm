package main

func findRepeatNumber3(nums []int) int {
	// 避免了内存空间的多次分配和 hash冲突处理
	h := make([]bool, len(nums))
	for _, val := range nums {
		if h[val] {
			return val
		}
		h[val] = true
	}
	return 0
}

// 数组中的数字的在0~len(num)-1 范围内
// 所谓的冲突就是 多个值可能共用了下标
func findRepeatNumber2(nums []int) int {
	for index, val := range nums {
		// 则不用交换
		if index == val {
			continue
		}
		// 说明冲突要交换
		if val == nums[val] {
			return val
		}
		nums[index], nums[val] = nums[val], nums[index]
	}
	return 0
}

// 剑指 Offer 03. 数组中重复的数字
// https://leetcode-cn.com/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof/
func findRepeatNumber(nums []int) int {
	numMap := make(map[int]bool)
	for _, n := range nums {
		if _, ok := numMap[n]; ok {
			return n
		}
		numMap[n] = true
	}
	return 0
}
