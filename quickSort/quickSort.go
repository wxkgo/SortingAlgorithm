package sort

func getMid(arr []int,left int,right int) int {
	pivot := arr[left]
	for left < right {
		for left < right && arr[right] >= pivot {
			right--
		}
		arr[left] = arr[right]
		for left < right && arr[left] <= pivot {
			left++
		}
		arr[right] = arr[left]
	}
	arr[left] = pivot
	return left
}

func QuickSort(arr []int,left int,right int)  {
	mid := getMid(arr,left,right)
	if left < mid - 1 {
		QuickSort(arr,left,mid - 1)

	}
	if right > mid + 1 {
		QuickSort(arr,mid + 1,right)
	}
}
