package mySort

//二分查找算法 非递归
func binarySearch(nums []int, key int) int {
	start, end := 0, len(nums)-1
	var mid int
	//查找的的key有可能是 数组的最后一个，所以 stat >=end
	for start <= end {
		//防止溢出的优化
		mid = (end-start)>>1 + start
		if nums[mid] == key {
			return mid
		} else if key > nums[mid] {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	//找不到就返回-1
	return -1
}

//二分查找递归
func binarySearch2(nums []int, start, end, key int) int {
	mid := (end-start)>>1 + start
	//注意这里有个小细节， start> end ,查找的key 可能是改数组最后一个，所以不能过早的退出
	if start > end {
		return -1
	}

	//nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	if nums[mid] == key {
		return mid
	} else if key > nums[mid] {
		return binarySearch2(nums, mid+1, end, key)
	}
	return binarySearch2(nums, start, mid-1, key)
}