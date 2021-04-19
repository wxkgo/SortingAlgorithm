package sort

func BubbleSort(arr []int,n int) {
	for i := n; i > 1; i-- {
		for j := 1;j < i;j++ {
			if arr[j] < arr[j - 1] {
				arr[j],arr[j - 1] = arr[j - 1],arr[j]
			}
		}
	}
}
