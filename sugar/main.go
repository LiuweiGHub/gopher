package main

import "fmt"

func main() {
	i := 0
	i, j := 1, 2
	fmt.Printf("i = %d, j = %d \n", i, j)

	func3()
}

// 编译报错
//func getInt(i int) {
//	i := 3
//}

// 作用域不同
func func3() {
	i, j := 0, 0

	if true {
		j, k := 1, 2
		fmt.Printf("j = %d, k = %d \n", j, k)
	}

	fmt.Printf("i = %d, j = %d \n", i, j)

	//f3(1,2,"3", "ssss",4)

	s := []int{1, 2, 3, 4}
	f3(s...)
}

// 语法糖  :=
/**
  a,b := func(0, 1)
  c,b := func(0, 2)

  1.当左侧有新变量时，b相当于被重新赋值 ， 如果左侧变量没变 则会报错 no new variable
  2.只能在函数中使用 := , 无法声明全局变量
  3.

*/

// 可变参数（0个或多个）  参数类型前面加...
/**
  1.可变参数必须在参数最后一位
  2.可变参数在函数内部作为切片使用
  3.可变参数如果不传为nil切片
  4.可变参数必须相同类型，如果需要不同类型 可以定义为...interface{}
  5.可变参数可以直接传递一个切片， 切片部分需使用 slice...表示， 但是必须明确类型

*/

func f3(nums ...int) {
	fmt.Println(nums)
}
