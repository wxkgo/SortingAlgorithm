package sort

import "fmt"

func InsertionSort(arr []int,n int) {
	for i := 1;i < n;i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j + 1] = arr[j]
			j--
		}
		arr[j + 1] = key
		fmt.Println(arr)
	}
}
