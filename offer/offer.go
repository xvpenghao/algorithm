package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(replaceSpace2("We are happy."))
}

//减少了内存的消耗，s+=""，这样会频繁的创建字符串对象，分配内存
func replaceSpace2(s string) string {
	//%20
	strRunes := make([]rune, len(s)*3)
	var index int
	for _, v := range s {
		if string(v) != " " {
			strRunes[index] = v
			index++
			continue
		}
		strRunes[index] = '%'
		index++
		strRunes[index] = '2'
		index++
		strRunes[index] = '0'
		index++
	}

	return string(strRunes[:index])
}

//剑指 Offer 05. 替换空格
//https://leetcode-cn.com/problems/ti-huan-kong-ge-lcof/
func replaceSpace(s string) string {
	var newStr string
	for _, v := range s {
		if unicode.IsSpace(v) {
			newStr += "%20"
			continue
		}
		newStr += string(v)
	}
	return newStr
}

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
	//与关键字比较，来踢行和踢列
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

//剑指 Offer 04. 二维数组中的查找
//https://leetcode-cn.com/problems/er-wei-shu-zu-zhong-de-cha-zhao-lcof/
func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	//开始竖着找
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

func findRepeatNumber3(nums []int) int {
	//避免了内存空间的多次分配和 hash冲突处理
	h := make([]bool, len(nums))
	for _, val := range nums {
		if h[val] {
			return val
		}
		h[val] = true
	}
	return 0
}

//数组中的数字的在0~len(num)-1 范围内
//所谓的冲突就是 多个值可能共用了下标
func findRepeatNumber2(nums []int) int {
	for index, val := range nums {
		//则不用交换
		if index == val {
			continue
		}
		//说明冲突要交换
		if val == nums[val] {
			return val
		}
		nums[index], nums[val] = nums[val], nums[index]
	}
	return 0
}

//剑指 Offer 03. 数组中重复的数字
//https://leetcode-cn.com/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof/
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
