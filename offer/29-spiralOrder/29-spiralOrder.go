package main

import "fmt"

/**
*@Author:徐鹏豪
*@Time: 2021/2/11 14:26
*@Description: 剑指 Offer 29. 顺时针打印矩阵
 */
func main() {
	matrix := [][]int{
		/*{1, 2, 3, 4, 5},
		{14, 15, 16, 17, 6},
		{13, 20, 19, 18, 7},
		{12, 11, 10, 9, 8},*/
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}
	// result 顺时针打印 123 6987 45
	// r d l u
	// 遍历方向 结束条件（都不能走了才退出） 循环条件（true）
	// 遍历过的如何不在遍历
	// 		【标记已经遍历过的矩阵，比如999】
	//		算出循环的轮次，然后 通过 cur+循环轮次 来确定时候已经遍历过了
	// 		使用单独的一个 map来存储那些下标已经遍历过了
	fmt.Println(spiralOrder2(matrix))
}

// 剑指 Offer 29. 顺时针打印矩阵
// https://leetcode-cn.com/problems/shun-shi-zhen-da-yin-ju-zhen-lcof/

func spiralOrder2(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}
	var result []int
	topRow, rightCol := 0, len(matrix[0])-1
	downRow, leftCol := len(matrix)-1, 0
	for {
		// 从左到右 列动
		for i := leftCol; i <= rightCol; i++ {
			result = append(result, matrix[topRow][i])
		}
		// 看上面这一行是否遍历完毕
		if topRow++; topRow > downRow {
			break
		}
		// 从上-下 行动
		for i := topRow; i <= downRow; i++ {
			result = append(result, matrix[i][rightCol])
		}
		// 看 右边这列是否遍历完毕
		if rightCol--; rightCol < leftCol {
			break
		}
		// 从右到左 列动
		for i := rightCol; i >= leftCol; i-- {
			result = append(result, matrix[downRow][i])
		}
		// 看下边这一行 是否边完毕
		if downRow--; downRow < topRow {
			break
		}
		// 从下到上 行动
		for i := downRow; i >= topRow; i-- {
			result = append(result, matrix[i][leftCol])
		}
		// 看左边这一列是否遍历完毕
		if leftCol++; leftCol > rightCol {
			break
		}
	}

	return result
}

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}

	row, col := 0, 0
	const left, right, up, down = 1, 2, 3, 4
	direct := right
	result := []int{matrix[0][0]}
	rightUpRow, rightUpCol := 0, len(matrix[0])-1
	leftDownRow, leftDownCol := len(matrix)-1, 0
	for {
		if direct == right && col+1 <= rightUpCol { // 右边
			col++
			result = append(result, matrix[row][col])
			continue
		} else if direct == down && row+1 <= leftDownRow { // 下边
			row++
			result = append(result, matrix[row][col])
			continue
		} else if direct == left && col-1 >= leftDownCol { // 左边
			col--
			result = append(result, matrix[row][col])
			continue
		} else if direct == up && row-1 > rightUpRow { // 上边
			row--
			result = append(result, matrix[row][col])
			continue
		}

		// 边界值的判断 不太正确 应该是每次 走完一个方向，然后就开始
		// 判断边界值的case
		switch direct {
		case right:
			direct = down
		case down:
			direct = left
		case left:
			direct = up
		case up:
			direct = right
			rightUpRow, rightUpCol = rightUpRow+1, rightUpCol-1
			leftDownRow, leftDownCol = leftDownRow-1, leftDownCol+1
		}

		// 他们交叉了说明 已经遍历完毕
		//  len(result) == (len(matrix)*len(matrix[0]) 如果不增加这个判断 可能会导致 多添加内容
		if rightUpRow > leftDownRow || len(result) == (len(matrix)*len(matrix[0])) {
			break
		}
	}

	return result
}
