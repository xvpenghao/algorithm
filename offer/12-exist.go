package main

import (
	"fmt"
)

// 剑指 Offer 12. 矩阵中的路径
// https://leetcode-cn.com/problems/ju-zhen-zhong-de-lu-jing-lcof/

func main() {
	board := [][]byte{
		{'C', 'A', 'A'},
		{'A', 'A', 'A'},
		{'B', 'C', 'D'},
	}

	bol := exist(board, "AAB")
	fmt.Println(bol)

}

// 使用深度 优先算法,
func exist(board [][]byte, word string) bool {
	// 题中说了，可以从 数组的任意位置开始
	for rowIndex, row := range board {
		for colIndex, _ := range row {
			if dfs(board, word, rowIndex, colIndex, 0) {
				return true
			}
		}
	}
	return false
}

// 不能走重复的路，需要到时候标记一下
func dfs(board [][]byte, word string, rowIndex, colIndex int, keyIndex int) bool {
	// 失败的递归的终止条件
	// 越界 or 未命中
	if rowIndex < 0 || rowIndex >= len(board) || colIndex < 0 ||
		colIndex >= len(board[0]) ||
		word[keyIndex] != board[rowIndex][colIndex] {
		return false
	}
	// 成功的递归终止条件keyIndex 找完了
	if keyIndex == len(word)-1 {
		return true
	}
	// 避免走重复路径
	tmp := board[rowIndex][colIndex]
	board[rowIndex][colIndex] = '/'
	// 开始找
	res := dfs(board, word, rowIndex+1, colIndex, keyIndex+1) || dfs(board, word, rowIndex-1, colIndex, keyIndex+1) ||
		dfs(board, word, rowIndex, colIndex+1, keyIndex+1) || dfs(board, word, rowIndex, colIndex-1, keyIndex+1)
	board[rowIndex][colIndex] = tmp
	return res
}
