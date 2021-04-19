package sort

//arr 存储堆的数组
//n 数组长度
//i 待维护元素的下标
func heapify(arr []int,n int,i int)  {
	max := i
	lson := 2 * i + 1
	rson := 2 * i + 2
	if lson < n && arr[lson] > arr[max] {
		max = lson
	}
	if rson < n && arr[rson] > arr[max] {
		max = rson
	}
	if max != i {
		arr[max],arr[i] = arr[i],arr[max]
		heapify(arr,n,max)
	}
}

func HeapSort(arr []int,n int)  {
	//建堆
	for i := n / 2 - 1;i >= 0;i-- {
		heapify(arr,n,i)
	}
	//排序
	for i := n - 1;i > 0;i-- {
		arr[0],arr[i] = arr[i],arr[0]
		heapify(arr,i,0)
	}
}
