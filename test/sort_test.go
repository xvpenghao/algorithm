package test

import (
	"algorithm/mySort"
	"testing"
)

func TestHeapSort(t *testing.T){
	nums := []int{5,3,9,1,2,6}
	mySort.HeapSort(nums)
	t.Log("排序后的数组为：",nums)
}
