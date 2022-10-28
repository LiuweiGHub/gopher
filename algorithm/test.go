package main

import (
	"bytes"
	"fmt"
	"time"
)

type BinaryAdder interface {
	Add(int, int) int
}

type MyAddFunc func(int, int) int

func (f MyAddFunc) Add(x, y int) int {
	return f(x, y)
}

func MyAdd(x, y int) int {
	return x + y
}

func times(x, y int) int {
	return x * y
}

func partialTimes(x int) func(int) int {
	return func(y int) int {
		return times(x, y)
	}
}

func Max(n, m int, f func(y int)) {
	if n > m {
		f(n)
	} else {
		f(m)
	}
}

func makeSlice(n int) []byte {
	defer func() {
		if recover() != nil {
			panic(bytes.ErrTooLarge)
		}
	}()
	return make([]byte, n)
}

func foo1() {
	for i := 0; i <= 3; i++ {
		defer fmt.Println(i)
	}
}

func foo2() {
	for i := 0; i <= 3; i++ {
		defer func(n int) {
			fmt.Println(n)
		}(i)
	}
}

func foo3() {
	for i := 0; i <= 3; i++ {
		defer func() {
			fmt.Println(i)
		}()
	}
}

type T struct {
	a int
}

func (t T) M1() {
	t.a = 10
}

func (t *T) M2() {
	t.a = 11
}

type field struct {
	name string
}

func (p field) print() {
	fmt.Println(p.name)
}

func main() {
	var i BinaryAdder = MyAddFunc(MyAdd)
	fmt.Println(i.Add(5, 6))

	// 函数柯里化：把多个参数变换成单个参数
	times2 := partialTimes(2)
	times3 := partialTimes(3)
	times4 := partialTimes(4)
	fmt.Println(times2(5))
	fmt.Println(times3(5))
	fmt.Println(times4(5))

	// 函子 有点难 没看懂
	// 延续性递式 CPS  ， 函数不允许有返回值
	Max(7, 6, func(y int) {
		fmt.Printf("%d\n", y)
	})

	ss := makeSlice(7)
	fmt.Println(ss)

	foo1()
	foo2()
	foo3()

	// 方法
	var t T
	println(t.a) //0
	t.M1()
	println(t.a) //0
	t.M2()
	println(t.a) //11

	data1 := []*field{{"one"}, {"two"}, {"three"}}

	for _, v := range data1 {
		go v.print()
	}

	data2 := []field{{"four"}, {"five"}, {"six"}}

	for _, v := range data2 {
		go v.print()
	}

	time.Sleep(3 * time.Second)

}
