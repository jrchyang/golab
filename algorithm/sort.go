// 算法/main.go
package main

import "fmt"

// 快速排序算法示例
func main() {
	arr := []int{9, 7, 5, 11, 12, 2, 14, 3, 10, 6}
	fmt.Println("排序前：", arr)

	quickSort(arr, 0, len(arr)-1)
	fmt.Println("排序后：", arr)
}

func quickSort(arr []int, low, high int) {
	if low < high {
		pivot := partition(arr, low, high)
		quickSort(arr, low, pivot-1)
		quickSort(arr, pivot+1, high)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}
