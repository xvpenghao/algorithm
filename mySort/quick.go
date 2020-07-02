package mySort

/**
*@Author:徐鹏豪
*@Time:2019/8/5 0005
*@Description:快速排序 挖坑+分治
* 参考链接 https://blog.csdn.net/MoreWindows/article/details/6684558
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
		//下面的两部分循环操作，就相当于是在做值的交换
		for low <high && arrs[high]>=pivotal{
			high--
		}
		//这里可以在优化与一下
		if low < high{
			arrs[low] = arrs[high]
			//这样就可以减少循环
			low++
		}

		//看中枢左边的
		for low <high && arrs[low] <pivotal{
			low++
		}
		if low < high{
			arrs[high] = arrs[low]
			high--
		}
	}
	//算出中间的节点
	arrs[low] = pivotal
	return low
}