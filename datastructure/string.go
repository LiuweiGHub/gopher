package main

/**
string 是 8 比特字节的集合， 通常但不一定是utf8编码

string可以为空（长度为0）， 但不会是nil  重要！！！ 所以不要判断一个string == nil ， 直接判断len即可判断它是否为空
string对象不能修改 ？ 怎么理解

字符串的构建是先构建stringStruct  再转换成string

零拷贝转换（源码）：
func gostringnocopy(str *byte) string {
   ss := stringstruct{str: unsafe.Pointer(str),  len: findnullstr()}   // 先构造struct
   s := *(*string)unsafe.Pointer(&ss)             //再将结构体转为string
   return s
}

string在runtime包中就是stringStruct , 对外呈现 string
*/

//  []byte 转 string
//  这种转换需要一次内存copy
func GetStringBySlice(s []byte) string {
	return string(s)
}
