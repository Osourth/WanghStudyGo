package _struct

import (
	"fmt"
	"testing"
)

//==============================================匿名内部类============================================
func TestAnonymous(t *testing.T) {

	u := struct {						//直接定义匿名结构变量
		name string
		age byte
	}{
		name: "tom",
		age: 23,
	}

	type file struct {
		name string
		attr struct {					//定义匿名结构类型字段
			owner int
			perm int
		}
	}

	f := file{
		name : "test",
		//attr:  {						//会报错
		//	owner 21
		//	perm  23
		//}
	}

	f.attr.owner = 21
	f.attr.perm =23

	fmt.Println(u, f)
}

//struct中只有在所有字段类型全部支持，才可以做相等操作
//可使用指针直接操作结构字段，但不能时多级指针
func TestPtr(t *testing.T) {
	type user struct {
		name string
		age int
	}

	p := &user{
		name:"Tom",
		age: 20,
	}

	p.name = "mary"
	p.age++

	//p2 := &p											//报错，不能使用多级指针
	//*p2.name = "Jack"
}

