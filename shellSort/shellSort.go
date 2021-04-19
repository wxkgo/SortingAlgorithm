package sort

func ShellSort(arr []int,n int)  {
	var i,j,inc,key int
	for inc = n / 2;inc > 0;inc /= 2 {  //控制增量，控制趟数
		for i = inc;i < n;i++ {
			key = arr[i]
			for j = i;j >= inc && key < arr[j - inc];j -= inc {
				arr[j] = arr[j - inc]
			}
			arr[j] = key
		}
	}
}
