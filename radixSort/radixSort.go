package sort


func getMax(arr []int) int {
	max := arr[0]
	for i := 1;i < len(arr);i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}

func radixSort(arr []int,n int)  {
	var max,base int
	base = 1
	max = getMax(arr)
	tmp := make([]int,n)
	for max / base > 0 {
		bucket := make([]int,10)
		for i := 0;i < n;i++ {
			bucket[arr[i] / base % 10]++
		}
		for i := 1;i < 10;i++ {
			bucket[i] += bucket[i - 1]
		}
		for i := n - 1;i >= 0;i-- {
			tmp[bucket[arr[i] / base % 10] - 1] = arr[i]
			bucket[arr[i] / base % 10]--
		}
		for i := 0;i < n;i++ {
			arr[i] = tmp[i]
		}
		base *= 10
	}
}
