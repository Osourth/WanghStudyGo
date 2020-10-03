package array

import (
	"fmt"
	"testing"
)

//===================================================复制=====================================
//Go数组是值类型，复制和传参操作都会复制整个数组数据
func test1(x [2]int) {
	fmt.Printf("x: %p, %v\n",&x,x)
}

func TestCopy1(t *testing.T) {
	a := [2]int{10,20}

	var b [2]int

	b = a

	fmt.Printf("a: %p, %v\n", &a, a)
	fmt.Printf("b: %p, %v\n", &b, b)

	test1(a)
}
//如果需要可以通过指针或切片，以此避免数据复制
func test2(x *[2]int) {
	fmt.Printf("x: %p, %v\n",x, *x)

	x[1] += 100
}

func TestCopy2(t *testing.T) {
	a := [2]int{12,23}

	test2(&a)

	fmt.Printf("a: %p, %v\n",&a, a)
}