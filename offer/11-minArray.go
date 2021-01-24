package main

import "fmt"

func main() {
	// [-9,-9,-9,-8,-8,-7,-7,-7,-7,-6,-6,-6,-6,-6,-6,-6,-6,-6,-5,-5,-5,-5,-5,-4,-4,-4,-3,-3,-3,-3,-3,-3,-2,-2,-2,-2,-2,-2,-2,-2,-2,-2,-1,-1,-1,-1,-1,-1,0,0,0,1,1,2,2,2,2,2,2,2,3,3,3,3,4,4,4,4,4,5,5,5,5,5,5,5,6,6,6,7,7,7,7,7,8,9,9,9,10,10,10,10,10,10,10,-10,-9,-9,-9,-9]
	fmt.Println(minArray([]int{-9, -9, -9, -8, -8, -7, -7, -7, -7, -6, -6, -6, -6, -6, -6, -6, -6, -6, -5, -5, -5, -5, -5, -4, -4, -4, -3, -3, -3, -3, -3, -3, -2, -2, -2, -2, -2, -2, -2, -2, -2, -2, -1, -1, -1, -1, -1, -1, 0, 0, 0, 1, 1, 2, 2, 2, 2, 2, 2, 2, 3, 3, 3, 3, 4, 4, 4, 4, 4, 5, 5, 5, 5, 5, 5, 5, 6, 6, 6, 7, 7, 7, 7, 7, 8, 9, 9, 9, 10, 10, 10, 10, 10, 10, 10, -10, -9, -9, -9, -9}))
}

// 剑指 Offer 11. 旋转数组的最小数字
// https://leetcode-cn.com/problems/xuan-zhuan-shu-zu-de-zui-xiao-shu-zi-lcof/
/*func minArray(numbers []int) int {

}*/

func minArray(numbers []int) int {
	if len(numbers) == 0 {
		return -1
	}
	low := 0
	high := len(numbers) - 1
	for low < high {
		// 算出中枢节点
		pivot := low + (high-low)/2
		// 445 123   //结果1

		// 说明 中枢节点 在最小值的右侧，则右侧可以不判断
		if numbers[pivot] < numbers[high] {
			high = pivot
		} else if numbers[pivot] > numbers[high] { // 说明了 中枢节点在 最小值的左侧
			low = pivot + 1
		} else {
			// 否则，两个节点相同,因为数组的中存在重复节点值，所以不能立即确定
			high--
		}
	}
	return numbers[low]
}

// 类似暴力破解了
func minArray2(numbers []int) int {
	if len(numbers) == 0 {
		return -1
	}
	if len(numbers) == 1 {
		return numbers[0]
	}

	for i := 1; i < len(numbers); i++ {
		if numbers[i-1] > numbers[i] {
			return numbers[i]
		}
	}
	// 说明了不是旋转数组，则直接返回 numbers[0]
	return numbers[0]
}
