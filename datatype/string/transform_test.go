package main

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
	"testing"
	"unsafe"
)




//=============================================字符串转换============================================================================

func pp(format string, ptr interface{}) {
	p := reflect.ValueOf(ptr).Pointer()
	h := (*uintptr)(unsafe.Pointer(p))			//(*type)unsafe.Pointer(ptr),go是强语言类型，不同类型之间不能直接转换，通过指针的方式可以实现

	fmt.Printf(format, *h)
}

//修改字符串，须将其转换成可变类型（[]rune或[]byte）,待完成后再转换回来，但是不管如何转换，都需要重新分配内存，并复制数据
func TestTran_1(t *testing.T) {

	s := "hello, world!"
	pp("s: %x\n", &s)

	bs := []byte(s)
	s2 := string(bs)

	pp("string to []byte, bs: %x\n",&bs)
	pp("[]byte to string, s2: %x\n",&s2)

	rs := []rune(s)
	s3 := string(rs)

	pp("string to []rune, rs: %x\n",&rs)
	pp("[]rune to string, s3: %x\n",&s3)


}


//某些时候，转换操作会拖累算法性能，可尝试用“非安全”方法进行改善
func toString(bs []byte) string{
	return *(*string)(unsafe.Pointer(&bs))
}

func TestTran_2(t *testing.T) {
	bs := []byte("hello world")
	s:= toString(bs)

	fmt.Printf("bs: %x\n", &bs[0])
	fmt.Printf("s : %x\n", &s)
}


//=============================================性能问题============================================================================

//在用加法操作符拼接字符串时，每次都会重新分配内存，因此在构建超大字符串时，性能就会极差
func test() string {
	var s string
	for i := 0; i < 1000; i++{
		s += "a"
	}

	return s
}
//改进思路是预分配足够的空间，常用方法是用strings.Join函数，它会统计所有参数长度，并一次性完成内存分配操作
func test2() string {
	s := make([]string,1000)

	for i := 0; i < 1000; i++ {
		s[i] = "a"
	}
	return strings.Join(s,"")
}
//使用bytes.Buffer也能完成类似的操作
func test3() string {
	var b bytes.Buffer
	b.Grow(1000)				//事先准备足够的内存，避免中途扩展

	for i := 0; i < 1000; i++ {
		b.WriteString("a")
	}

	return b.String()
}

func BenchmarkTest(b *testing.B) {
	for i := 1; i < b.N; i++{
		//test2()
		//test()
		test3()
	}

}
