package main

import "fmt"

/**

多个defer满足LIFO原则， 后进先出

编码习惯养成：申请资源后，立即使用defer关闭资源

*/

func main() {
	deferFuncParameter()
	deferPrintArr()
	a := deferFuncReturn()
	fmt.Println(a)
}

func deferFuncParameter() {
	var i = 1

	defer fmt.Println(i) // 1  延迟函数的参数在defer语句出现时就已经确立下来了

	i = 2
	return
}

func printArr(array *[3]int) {
	for i := range array {
		fmt.Println(array[i])
	}
}

func deferPrintArr() {
	var arr = [3]int{1, 2, 3}
	defer printArr(&arr)

	arr[0] = 10
	return
}

func deferFuncReturn() (result int) {
	var i = 1
	defer func() {
		result++
	}()
	return i
}
