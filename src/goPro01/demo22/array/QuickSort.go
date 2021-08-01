package main

import (
	 "fmt"
)
/**
1、快速排序
*/
var arr = [10]int{10,9,5,8,7,1,6,2,3,4}

func main()  {
	fmt.Println("arr:")
	fmt.Println(arr)
	QuickSort(0 , 9)
	fmt.Println("sorted arr : ")
	fmt.Println(arr)
	a1 := []int{10,9,5,8,7,1,6,2,3,4}
	a := [10]int{10,9,5,8,7,1,6,2,3,4}
	fmt.Println("a1 :" , a)
	QuickSort1(a1, 0 , 9)
	fmt.Println("sort a1 :" ,a)
	fmt.Println("a :" , a)
	QuickSort2(&a, 0 , 9)
	fmt.Println("sort a :" ,a)
}

func QuickSort(left , right int) int  {
	var i , j ,temp int
	if left > right {
		return 0
	}
	temp = arr[left] 
	i = left 
	j = right

	for i!=j {
		for arr[j] >= temp && i<j {j--}  //当右侧的数大于基准，则移动，进而停下的位置是右边找小于基数的数的位置 
		for arr[i] <= temp && i<j {i++}	 //当左侧的数小于基准，则移动，进而停下的位置是左边找大于基数的数的位置
		fmt.Println( i , j )
		if i < j {
			arr[i] , arr[j] = arr[j] , arr[i]  
		}
	}
	arr[left] , arr[i] = arr[i] , arr[left]
	fmt.Println("---",arr)
	QuickSort(left , i-1)
	QuickSort(i+1 , right)
	return 0
}


func QuickSort1(a []int , left , right int ) int {

	if left > right{
		return 0
	}
	i := left 
	j := right 
	temp := a[left] 
	for i!=j{
		for a[j] >= temp && i<j{j--}
		for a[i] <= temp && i<j{i++}
		if i<j{
			a[i] , a[j] = a[j] , a[i]
		} 
	}
	a[left] , a[i] = a[i] , a[left]
	QuickSort1(a , left , i-1)
	QuickSort1(a , i+1 , right)
	return 0
}

func QuickSort2(a *[10]int , left , right int ) int {

	if left > right{
		return 0
	}
	i := left 
	j := right 
	temp := a[left] 
	for i!=j{
		for a[j] >= temp && i<j{j--}
		for a[i] <= temp && i<j{i++}
		if i<j{
			a[i] , a[j] = a[j] , a[i]
		} 
	}
	a[left] , a[i] = a[i] , a[left]
	QuickSort2(a , left , i-1)
	QuickSort2(a , i+1 , right)
	return 0
}
