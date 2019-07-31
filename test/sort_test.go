package test

import (
	"algorithm/mySort"
	"testing"
)

//测试堆排序
func TestHeapSort(t *testing.T) {
	nums := []int{5, 3, 9, 1, 2, 6}
	mySort.HeapSort(nums)
	t.Log("排序后的数组为：", nums)
}

//测试冒泡排序
func TestBobbleSort(t *testing.T) {
	nums := []int{5, 3, 9, 1, 2, 6}
	mySort.BobbleSort(nums)
	t.Log("排序后的数组为：", nums)
}

//测试选择排序
func TestSelectSort(t *testing.T) {
	nums := []int{5, 3, 9, 1, 2, 6}
	mySort.SelectSort(nums)
	t.Log("排序后的数组为：", nums)
}
