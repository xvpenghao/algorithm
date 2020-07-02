package mySort

/**
*@Author:徐鹏豪
*@Time:2019/7/31 0031
*@Description:冒泡排序
 */
func BobbleSort(nums []int) {
	//定义标志变量
	var varFlag bool
	//对比的次数
	for i := 1; i < len(nums); i++ {
		varFlag = true
		for j := len(nums) - 2; j >= i-1; j-- {
			//则两者就交换
			if nums[j+1] < nums[j] {
				nums[j+1], nums[j] = nums[j], nums[j+1]
				//说明进行排序了
				varFlag = false
			}
		}
		//我排序了，现在还不能退出
		if varFlag {
			break
		}
	}
}


func BobbleSort2 (nums []int64){
	//冒泡排序
	for i:=1;i<len(nums);i++{
		for j:=0;j<len(nums)-i;j++{
			if nums[j]>nums[j+1]{
				nums[j],nums[j+1] = nums[j+1],nums[j]
			}
		}
	}
}


func BobbleSort3 (nums []int64){
	var isSwap bool
	for i:=1;i<len(nums);i++{
		isSwap = false
		//len(nums)-i;
		for j:=0;j<len(nums)-i;j++{
			if nums[j]>nums[j+1]{
				//是否做了位置的交换
				isSwap = true
				nums[j],nums[j+1] = nums[j+1],nums[j]
			}
		}
		if !isSwap{
			break
		}
	}
}
