package mySort

import "fmt"

/**
*@Author:徐鹏豪
*@Time:2019/8/1 0001
*@Description:插入排序
 */
func InsertSort(arrs []int) {
	for i := 1; i < len(arrs); i++ {
		j := i-1
		temp := arrs[i]
		//每次都是从一个有序的列表中进行插入
		for j > 0 && arrs[j] > temp {
			//开始移动位置,找到可以插入位置
			arrs[j+1] = arrs[j]
			j--
		}
		arrs[j+1] = temp
	}
}

//插入排序-简单插入排序
func InsertSort2 (nums []int64){
	//从无序劣列表中，取值每次插入有序的列表中,目标构建有序序列
	for i:=0;i<len(nums);i++{
		curNum := nums[i]
		j :=i
		//对于已经排序的序列中，从后往前扫描，找到相应的位置并插入
		//这里的j--就是移动后的位置
		for ;j>0 && nums[j-1] > curNum;j--{
			//移动位置
			nums[j] = nums[j-1]
		}
		//插入该值
		nums[j] = curNum
	}
	fmt.Println(nums)
}