// study_list01.go
package main

import (
	"fmt"
)

func printArray(arr [5]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func printArrayPoint(arr *[5]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func main() {
	//定义数组方法
	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{2, 4, 6, 8, 10}
	//多维数组
	var grid [4][5]bool

	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)
	//数组遍历方法
	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}
	for i := range arr3 {
		fmt.Println(arr3[i])

	}
	for i, v := range arr3 {
		fmt.Println(i, v)
	}
	for _, v := range arr3 {
		fmt.Println(v)
	}
	//求数组中的最大值及index
	numbers := [...]int{9, 5, 8, 6, 13, 3, 15, 30}
	maxi := -1
	maxValue := -1
	for i, v := range numbers {
		if v > maxValue {
			maxi, maxValue = i, v
		}
	}
	fmt.Println(maxi, maxValue)

	//求数组之和
	//可以通过_省略变量，不仅range，任何地方都可以通过_省略变量
	//如果只要i(下标)，可写成for i := range numbers
	//java/python 只能for each value，不能同时获取i，v
	sum := 0
	for _, v := range numbers {
		sum += v
	}
	fmt.Println(sum)

	//数组是值类型
	//[10]int 和[20]int 是不同类型
	//调用func f(arr [10]int)会拷贝数组
	//在go语言中一般不直接使用数组，使用切片
	fmt.Println("printArray(arr1)")
	printArray(arr1)
	//cannot use arr2 (type [3]int) as type [5]int in argument to printArray
	//printArray(arr2)
	fmt.Println("printArray(arr3)")
	printArray(arr3)
	fmt.Println("arr1 and arr3")
	fmt.Println(arr1, arr3)

	//指针传递
	printArrayPoint(&arr1)
	printArrayPoint(&arr3)

	fmt.Println(arr1, arr3)

}
