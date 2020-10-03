package main

import (
	"fmt"
	"sync"
	"testing"
)

//可以像访问匿名字段成员那样调用其方法，由编译器负责查找
type data struct {
	sync.Mutex
	buf [1024]byte
}

func TestAnonymousFunc(t *testing.T) {
	d := data{}

	d.Lock()

	defer d.Unlock()
}

//结构体中的方法会有同名的问题，利用这种特性，可以实现类似覆盖（override）的操作
type user struct {}

type manage struct {
	user
}

func (user) toString() string {				//注意此处不存在传入实例对象，是否类似于静态方法
	return "user"
}

func (m manage) toString() string {

	return m.user.toString() + "manager"
}

func TestOverride(t *testing.T) {
	var m manage

	fmt.Println(m.toString())
//	fmt.Println(user.toString())			//此处并不能直接调用
	fmt.Println(m.user.toString())
}


//==================================================方法集=======================================================
//类型 T 方法集包含所有receiver T　方法
//类型 *T 方法集包含所有receiver T + *T 的方法
//匿名嵌入S, T方法集包含所有receiver S的方法
//匿名嵌入 *S ， T 方法集包含所有receiver S + *S 方法
//匿名嵌入S 或 *S, *T方法集包含所有receiver S + *S方法
