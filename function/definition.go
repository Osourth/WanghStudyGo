package main

import "fmt"

//===============================================================================
//receiver    前置实例接收参数
//receiver的类型自然可以是基础类型或者指针类型，这会关系到调用时对象实例是否被复制
//可以使用实例值或者指针类型调用方法，编译器会根据方法receiver类型自动在基础类型和指针类型之间转换
//不能使用多级指针调用方法
type N int

func (n N) value() {							//相当于 func value(n N)
	n++
	fmt.Printf("v: %p, %v\n", &n, n)
}

func (n *N) pointer() {							//相当于 func pointer(n *N)
	(*n)++
	fmt.Printf("v: %p, %v\n", n, *n)
}

func main() {

	var a N = 25

	a.value()
	a.pointer()									//方法前指代的是指针，但调用时不需要使用指针

	fmt.Printf("a,: %p, %v\n", &a, a)


}
