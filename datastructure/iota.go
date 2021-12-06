package main

import "fmt"

/**
iota常用于const表达式中， 且是从0开始， const声明块中每增加一行 就自增 1

iota特点：
   第一行必须有一个表达式， 后续的常量如果没有表达式，则继承上面的表达式
   可以使用下划线 _ 跳过不想要的值

iota是go中常量计数器，只能用于常量中
*/

const (
	A int    = 1 << iota //  << 将 0 的二进制左移 1 位， 2^0  1
	B                    // 2^1 2
	C                    // 2^2 4
	D                    // 2^3 8
	E = iota             // 4
	F                    // 5
	G = 1e6              // 1000000
	H                    // 7  错了！！ 继承上面的表达式 1e6 也是表达式  所以应该等于1e6
)

/**
位掩码表达式
1 << iota  // 1 << 0   二进制 00000001  十进制 1
1
*/
const (
	aa = 1 << iota // iota = 0 ， 00000001 左移 0 位， 1
	ab             // iota = 1  00000001 左移1位， 00000010 2
	ac             // iota = 2  00000001 左移2位   00000100 4
	ad             // 8

	ae = 1 >> iota // iota = 4, 00000001 右移 4 位， 高位补 0 ， 00000000   0
	af
	ag
)

const (
	a = iota << 1 //0/2  0
	b = 1 >> iota //
	c = iota << 1 // 2^2 4
	d = 1 << iota //
	e
	f = "ha"
	g
	h = iota // 7
	_
	_
	_
	i // 11

)

const (
	lorcked = 1 << iota
	workend
	workend1
	workend2
)

const (
	ba, bb = 1 << iota, 2 << iota // 1   00000010     2
	bc, bd                        // 2 4
	be, bf                        // 4 8
	_, _
	bg, bh // 16,32                     // 00000010  00100000
)

func main() {
	fmt.Println(A, B, C, D, E, F, G, H)
	fmt.Println(a, b, c, d, e, f, g, h, i)
	fmt.Println(aa, ab, ac, ad, ae, af, ag)
	fmt.Println(ba, bb, bc, bd, be, bf, bg, bh)
	fmt.Println(lorcked, workend, workend1, workend2)
}
