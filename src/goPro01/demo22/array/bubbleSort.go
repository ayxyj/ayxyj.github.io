package main

import (
	"fmt"
)

/**
1、冒泡排序
2、选择排序
*/
func main() {
	fmt.Println("arr : ")
	arr := [10]int{10, 9, 5, 8, 7, 1, 6, 2, 3, 4}
	fmt.Println(arr)
	BubbleSort(&arr, 10)
	fmt.Println(arr)
	SelectSort(&arr, 10)
	fmt.Println(arr)
}

func SelectSort(arr *[10]int, len int) {
	for i := 0; i < len; i++ {
		for j := i + 1; j < len; j++ {
			//大的在前
			if arr[i] < arr[j] {
				Swap(&arr[i], &arr[j])
			}
		}
	}
}

func BubbleSort(arr *[10]int, len int) {
	for i := 0; i < len; i++ {
		for j := 0; j < len-i-1; j++ {
			//大的在后  冒泡大个子的
			if arr[j] > arr[j+1] {
				//arr[j] , arr[j+1] = arr[j+1] , arr[j]
				Swap(&arr[j], &arr[j+1])
			}
		}
	}
}

func Swap(a, b *int) {
	*a, *b = *b, *a
}
