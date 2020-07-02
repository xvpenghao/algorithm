package mySort

import "fmt"

/**
*@Author:徐鹏豪
*@Time:2019/7/28 0028
*@Description:堆排序->大堆排序
* 参考链接 https://www.cnblogs.com/chengxiao/p/6129630.html
 */
func HeapSort(nums []int) {
	if len(nums) == 0 {
		return
	}
	//根据双亲节点，来构造一个大顶堆
	for i := len(nums)/2 - 1; i >= 0; i-- {
		buildHeap(nums, i, len(nums))
	}

	//从这打印的结果就能看出，构建的大顶堆还是小顶堆
	//fmt.Println(nums)

	//使这个数组有序，数组中最有一个依次和堆顶元素互换
	//然后在重新构造
	for i := len(nums) - 1; i > 0; i-- {
		//交换完之后，最后一个就是最大的了
		nums[0], nums[i] = nums[i], nums[0]
		//提出最后一个，在重新构造
		buildHeap(nums, 0, i)
	}
}

func buildHeap(nums []int, parentIndex, length int) {
	//临时保存数据
	key := nums[parentIndex]
	//用这个双亲节点的左右孩子和这个key比较
	for i := parentIndex*2 + 1; i < length; i = i*2 + 1 {
		//左节点是否小于右节点
		if (i+1) <= length && nums[i] < nums[i+1] {
			i++
		}
		//说明这个双亲节点最大了
		if key >= nums[i] {
			break
		}
		//将这个大值复制给这个双亲节点
		nums[parentIndex] = nums[i]
		parentIndex = i
	}
	nums[parentIndex] = key
}


//构建大顶堆
//升序适合 大顶堆
//降序适合 小顶堆
func HeapSort2() {
	nums := []int{3, 2, 5, 7, 9, 23, 13, 6, 8, 1}

	//从下到上，从左到右，从非叶结点开始比较
	//等得到一个复合规范的大顶堆
	for i := len(nums)/2 - 1; i >= 0; i-- {
		//从每个非叶子节点开始比较，构建一个复合大顶堆的树
		buildHeap2(nums, i, len(nums))
	}

	//每次比完提出最后一个，使其构建一个有序的升序数值
	for j := len(nums) - 1; j > 0; j-- {
		//头尾交换，踢出尾部
		nums[0], nums[j] = nums[j], nums[0]
		buildHeap2(nums, 0, j)
	}
	fmt.Println(nums)
}

func buildHeap2(nums []int, nodeIndex, length int) {
	temp := nums[nodeIndex]

	for i := nodeIndex*2 + 1; i < length; i = i*2 + 1 {
		//比较左右节点得出一个最大值
		if i+1 < length && nums[i] < nums[i+1] {
			//移动到右节点
			i++
		}

		//再次比较得出最大值
		if nums[i] > temp {
			//交换
			nums[nodeIndex] = nums[i]
			nodeIndex = i
		} else {
			break
		}
	}
	//值的替换
	nums[nodeIndex] = temp
}


//构建小顶堆
func buildLittleHeap(nums []int, nodeIndex, length int) {
	keyVal := nums[nodeIndex]
	for i := nodeIndex*2 + 1; i < length; i = i*2 + 1 {
		if i+1 < length && nums[i] > nums[i+1] {
			i++
		}

		if nums[i] < keyVal {
			nums[nodeIndex] = nums[i]
			nodeIndex = i
		} else {
			break
		}
	}
	nums[nodeIndex] = keyVal
}