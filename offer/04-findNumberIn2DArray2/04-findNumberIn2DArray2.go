package _4_findNumberIn2DArray2

/*
* 正解
思路：深入分析题，挖出题中隐藏的知识点
1、行 顺序递增，列顺序递增
2、矩形角通过和关键字比较来 排除行，和排除列
3、最终得到结果, 矩形角，行的终点，列的起点
*/
func findNumberIn2DArray2(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	row, col := 0, len(matrix[0])-1
	// 与关键字比较，来踢行和踢列
	for row < len(matrix) && col >= 0 {
		if matrix[row][col] == target {
			return true
		} else if matrix[row][col] > target {
			col--
		} else {
			row++
		}
	}

	return false
}

// 剑指 Offer 04. 二维数组中的查找
// https://leetcode-cn.com/problems/er-wei-shu-zu-zhong-de-cha-zhao-lcof/
func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	// 开始竖着找
	for i := 0; i < len(matrix); i++ {
		if _, ok := findInsertLoc(matrix[i], target); ok {
			return true
		}
	}
	return false
}

func findInsertLoc(nums []int, target int) (int, bool) {
	left, right := 0, len(nums)-1
	var mid int
	for left <= right {
		mid = left + (right-left)>>1
		if nums[mid] == target {
			return mid, true
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		}
	}

	return left, false
}
