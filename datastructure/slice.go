package main

import "fmt"

/*
又称动态数组
slice copy过程不会发生扩容
slice copy两个切片时，取两个切片长度的最小值， 如 copy 长度为 10的切片到 长度为5的切片时， 只copy前五个元素
扩容算法：
  容量小于1024 就翻倍
  容量大于1024 就翻1.25倍
*/

func main() {
	var array [10]int
	var slice = array[5:7] // 左闭右包

	fmt.Println(len(slice))             // 1
	fmt.Println(cap(slice))             // 5 为啥是5? 原来如此，这种方式为使用数组创建切片， 方式 arr := [10]int  slice:=arr[5:7] , 切片长度为数组小标5开始 包含5，到小标7结束，不包含7， 所以长度为2， *数组后面的内容都作为切片的预留内存，所以cap为5， 即起始下标之后都作为slice的cap*
	fmt.Println(&slice[0] == &array[5]) // true

	var s []int
	s = append(s, 1, 2, 3)

	newS := addElement(s, 4)

	fmt.Println(&s[0] == &newS[0])

	ss := make([]int, 5, 10)
	ss1 := ss[0:5]
	ss2 := ss[0:5:5]
	ss3 := ss[:5:10] // slice(start:end:cap)  slice也可以创建新的切片

	fmt.Println(cap(ss1), cap(ss2), cap(ss3))

}

func addElement(s []int, e int) []int {
	return append(s, e)
}
