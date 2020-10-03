package slice

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

//==================================================定义=================================================
//前者仅定义了一个[]int类型变量，并未执行初始化，而后者则用初始化表达式完成了全部创建过程
// a== nil,仅表示她是一个未初始化的切片对象，切片本身依然会分配所需的内存
func TestDefinition(t *testing.T) {
	var a []int
	b := []int{}

	println(a == nil, b == nil)

	fmt.Printf("a: %#v\n", (*reflect.SliceHeader)(unsafe.Pointer(&a)))
	fmt.Printf("b: %#v\n", (*reflect.SliceHeader)(unsafe.Pointer(&b)))

	fmt.Printf("a size: %d\n", unsafe.Sizeof(&a))
	fmt.Printf("b size: %d\n", unsafe.Sizeof(&b))
}



//===============================================reslice======================================================
func TestStack(t *testing.T) {

	//栈的最大内存是5
	stack := make([]int, 0 , 5)

	//入栈
	push := func(x int) error {
		n := len(stack)
		if n == cap(stack) {
			return errors.New("stack is full")
		}
		stack = stack[:n+1]
		stack[n] = x

		return nil
	}

	//出栈
	pop := func() (int, error){
		n := len(stack)
		if n == 0 {
			return 0, errors.New("stack is empty")
		}
		x := stack[n-1]
		stack = stack[:n-1]

		return x,nil
	}

	//入站测试
	for i := 0; i < 7; i++ {
		fmt.Printf("push %d : %v, %v \n", i, push(i), stack)

	}

	//出栈测试

	for i := 0; i < 7; i++ {
		x, err := pop()
		fmt.Printf("pop: %d, %v, %v\n", x, err, stack)
	}

}


//========================================================copy=================================================
//在两个切片对象间复制数据，允许指向同一底层数组，允许目标区间重叠。

func test(s []int){

	fmt.Printf("x: %p \n",s)
	s[0] = 10020
}


func TestCopt(t *testing.T)  {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}


	s1 := s[3:6]
	n := copy(s[2:], s1)		//在同一底层数组的不同区间复制
	fmt.Println(n, s)

	s2 := make([]int, 6)		//在不同数组间复制
	n = copy(s2, s)
	fmt.Println(n, s2)

	b := make([]byte, 3)
	n = copy(b, "abcde")	//可以直接从字符串中复制数据到[]byte
	fmt.Println(n, b)

	fmt.Printf("%p \n", s)
	fmt.Println(s)
	test(s)
	fmt.Println(s)
}