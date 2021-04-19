package sort

func MergeSort(arr []int,n int)  {
	tmp := make([]int,n)
	msort(arr,tmp,0,n - 1)
}

func msort(arr []int,tmp []int,left int,right int)  {
	if left < right {
		mid := (left + right) / 2
		msort(arr,tmp,left,mid)
		msort(arr,tmp,mid + 1,right)
		merge(arr,tmp,left,mid,right)
	}
}

func merge(arr []int,tmp []int,left int,mid int,right int)  {
	//标记左半区第一个未排序的元素
	l := left
	//标记右半区第一个未排序的元素
	r := mid + 1
	//临时数组元素下标
	pos := left
	//合并
	for l <= mid && r <= right {
		if arr[l] < arr[r] {
			tmp[pos] = arr[l]
			pos++
			l++
		} else {
			tmp[pos] = arr[r]
			pos++
			r++
		}
	}
	//合并左半区剩余元素
	for l <= mid {
		tmp[pos] = arr[l]
		pos++
		l++
	}
	//合并右半区剩余元素
	for r <= right {
		tmp[pos] = arr[r]
		pos++
		r++
	}
	//把临时数组中合并后的元素复制会原数组
	for left <= right {
		arr[left] = tmp[left]
		left++
	}
}
