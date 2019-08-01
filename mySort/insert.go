package mySort

/**
*@Author:徐鹏豪
*@Time:2019/8/1 0001
*@Description:插入排序
 */
func InsertSort(arrs []int) {
	for i := 1; i < len(arrs); i++ {
		j := i
		temp := arrs[i]
		//每次都是从一个有序的列表中进行插入
		for j > 0 && arrs[j-1] > temp {
			//开始移动位置,找到可以插入位置
			arrs[j] = arrs[j-1]
			j--
		}
		if j != i {
			arrs[j] = temp
		}
	}
}
