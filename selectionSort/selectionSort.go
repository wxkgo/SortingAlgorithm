package sort


func SelecttionSort(arr []int,n int) {
	for i := 0; i < n; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		if i != minIndex {
			arr[i],arr[minIndex] = arr[minIndex],arr[i]
		}
	}
}
