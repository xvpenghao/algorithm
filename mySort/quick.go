package mySort

/**
*@Author:徐鹏豪
*@Time:2019/8/5 0005
*@Description:快速排序
*/
func QuickSort(arrs []int, low,high int){
	if low <high{
		pivotkey:= partition(arrs,low,high)

		//算左边
		QuickSort(arrs,low,pivotkey-1)
		//算右边
		QuickSort(arrs,pivotkey+1,high)
	}
}

//划分，分开
func partition(arrs []int,low,high int)int{
	//选择一个作为中枢记录
	pivotal := arrs[low]

	//以中枢key进行划分，左边小于key的，右边为大于key的
	for low < high{
		//先走右边 ,如果不写相等，则high则会处于死循环
		//而则循环就是看 low 和high在变化的
		for low <high && arrs[high]>=pivotal{
			high--
		}
		//赋值
		arrs[low] = arrs[high]

		//看中枢左边的
		for low <high && arrs[low] <pivotal{
			low++
		}
		arrs[high] = arrs[low]
	}

	//算出中间的节点
	arrs[low] = pivotal
	return low
}