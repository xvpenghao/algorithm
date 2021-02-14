package main

import (
	"fmt"
)

/**
*@Author:徐鹏豪
*@Time: 2021/2/14 08:55
*@Description:
 */
func main() {
	pushed := []int{1, 2, 3, 4, 5} // 我怎样才能知道左边是 3呢
	popped := []int{4, 3, 5, 1, 2}
	// [4,3,5,1,2]
	// key是唯一的， 找到pop在push的的左边，右边的是一堆，需要移出 pushed中的元素
	fmt.Println(validateStackSequences3(pushed, popped))
}

// 剑指 Offer 31. 栈的压入、弹出序列
// https://leetcode-cn.com/problems/zhan-de-ya-ru-dan-chu-xu-lie-lcof/
func validateStackSequences3(pushed []int, popped []int) bool {
	i := 0
	var stack []int
	for _, ele := range pushed {
		stack = append(stack, ele)
		// 判断栈顶元素 是否和 pop相等
		for len(stack) > 0 && stack[len(stack)-1] == popped[i] {
			stack = stack[:len(stack)-1]
			i++ // 这里应该不用担心小标越级， 应为他的 stack和 push 和pop的长度都一致
		}
	}
	return len(stack) == 0
}
func validateStackSequences2(pushed []int, popped []int) bool {
	var stack []int
	eleMap := make(map[int]bool)
	startPushIndex := 0
	for _, popVal := range popped {
		// 从哪里开始 没有则 push
		// 有则pop 然后看值是否相等
		_, ok := eleMap[popVal]
		for startPushIndex < len(pushed) && !ok {
			stack = append(stack, pushed[startPushIndex])
			eleMap[pushed[startPushIndex]] = true
			// 根pop的值一样则 break
			if pushed[startPushIndex] == popVal {
				startPushIndex++
				break
			}
			startPushIndex++
		}
		// pop中的值 和stack 弹窗的值不相等
		if len(stack) > 0 && stack[len(stack)-1] != popVal {
			return false
		}
		stack = stack[:len(stack)-1]
	}
	return true
}

// push 1,2,3,4,5,

// pop 4,5,3,2,1 ok
// pop 4,3,5,[1],2 error 弹5后 可以弹 2
// pop 3,[5],4,2,1 error 弹3后 可以弹 2,4
// 观察分析，被弹数的 左右一位可以被弹，不能跨弹
// 确定pop中的元素，是否可以在 push中 弹出（获取弹出元素的 左或右一位）是否在 pop弹窗元素的位置中
// 可以弹窗，则将弹出的元素从push中踢出，
func validateStackSequences(pushed []int, popped []int) bool {
	if len(pushed) == 0 && len(popped) > 0 {
		return false
	}
	if len(popped) == 0 && len(pushed) > 0 {
		return false
	}

	// 数字均不相等
	leftEle, rightEle := -1, make(map[int]bool)
	for i := 0; i < len(popped); i++ {
		// 找出被弹元素在push中的位置
		popEleInPushIndex := -1
		for j := 0; j < len(pushed); j++ {
			if popped[i] == pushed[j] {
				popEleInPushIndex = j
			}
		}
		// 这个应该是能找到的
		if popEleInPushIndex < 0 {
			return false
		}
		// 得到被弹元素 在 pushed的 左边和 右边的位置
		leftVal, rightVal := -1, make(map[int]bool)
		if popEleInPushIndex > 0 {
			leftVal = pushed[popEleInPushIndex-1]
		}
		if popEleInPushIndex+1 < len(pushed) {
			for _, val := range pushed[popEleInPushIndex+1:] {
				rightVal[val] = true
			}
		}
		// 弹出的第一个 元素，就不用比了，一定能弹出的
		if i > 0 {
			if !(leftEle == popped[i] || rightEle[popped[i]]) {
				return false
			}
		}
		leftEle, rightEle = leftVal, rightVal
		// 移出弹窗元素
		var tmp []int
		if popEleInPushIndex+1 < len(pushed) {
			tmp = pushed[popEleInPushIndex+1:]
		}
		pushed = pushed[0:popEleInPushIndex]
		if len(tmp) > 0 {
			pushed = append(pushed, tmp...)
		}
	}
	return true
}
