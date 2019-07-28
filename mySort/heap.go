package mySort

/**
*@Author:徐鹏豪
*@Time:2019/7/28 0028
*@Description:堆排序->大堆排序
*/
func HeapSort(nums []int){
	if len(nums) ==0{
		return
	}
	//根据双亲节点，来构造一个大顶堆
	for i:= len(nums)/2-1;i>=0;i--{
		buildHead(nums,i,len(nums)-1)
	}
	//使这个数组有序，数组中最有一个依次和堆顶元素互换
	//然后在重新构造
	for i:=len(nums)-1;i>0;i--{
		//交换完之后，最后一个就是最大的了
		nums[0],nums[i] = nums[i],nums[0]
		//提出最后一个，在重新构造
		buildHead(nums,0,i-1)
	}
}

func buildHead(nums []int,parentIndex,length int){
	//临时保存数据
	key := nums[parentIndex]
	//用这个双亲节点的左右孩子和这个key比较
	for i:=parentIndex*2+1;i<=length;i =i*2+1{
		//左节点是否小于右节点
		if (i+1) <= length && nums[i] < nums[i+1]{
			i++
		}
		//说明这个双亲节点最大了
		if key >= nums[i]{
			break
		}
		//将这个大值复制给这个双亲节点
		nums[parentIndex] = nums[i]
		parentIndex  = i
	}
	nums[parentIndex] = key
}
