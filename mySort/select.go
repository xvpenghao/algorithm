package mySort

/**
*@Author:徐鹏豪
*@Time:2019/7/31 0031
*@Description:选择排序
 */
func SelectSort(arrs []int) {
	var minIndex int
	//可以少做一次的循环判断
	for i := 0; i < len(arrs)-1; i++ {
		minIndex = i
		for j := i + 1; j < len(arrs); j++ {
			if arrs[j] < arrs[minIndex] {
				minIndex = j
			}
		}
		//如果不一样则交换
		if minIndex != i {
			arrs[minIndex], arrs[i] = arrs[i], arrs[minIndex]
		}
	}
}
