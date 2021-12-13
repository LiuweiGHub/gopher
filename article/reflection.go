package main

import (
	"fmt"
	"reflect"
)

/**
- 反射提供一种让程序检查自身结构的能力
- 反射是困惑的源泉
*/

/**
反射三律：
1.反射可以将interface类型变量转换成反 射对象
2.反射可以将反射对象还原成interface对 象
3.反射对象可修改，value值必须是可设置 的
*/

func main() {
	var x float64 = 3.4
	t := reflect.TypeOf(x)

	fmt.Println(t)

	v := reflect.ValueOf(x)
	//v.SetFloat(7.3)  会panic ， 因为如果修改x的value是无效订单

	fmt.Println(v)

	var y float64 = v.Interface().(float64)

	fmt.Println(y)

	v1 := reflect.ValueOf(&x) //获取指针
	v1.Elem().SetFloat(7.1)   //reflect.Value 提供了 Elem() 方法，可以获得指针向指向的 value

	fmt.Println(x)
}
