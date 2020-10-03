package dictionary

import (
	"fmt"
	"sync"
	"testing"
	"time"
	"unsafe"
)

//============================基本操作=============================
//访问不存在的键值，默认会返回零值，不会引发错误，可以使用ok-idiom模式
func TestBaseOperation(t *testing.T){

	m := map[string]int{
		"a": 1,
		"b": 2,
	}

	m["a"] = 10									//修改
	m["c"] = 30									//新增


	if v,ok := m["d"]; ok {						// 使用ok-idiom 判断 key 是否存在，返回值
		println(v)
	}

	delete(m, "d")							//删除键值对，不存在时，不会出错
}


//========================================安全========================================
//在迭代期间删除或新增键值是安全的
func TestSafe1(t *testing.T){
	m := make(map[int]int)

	for i := 0; i < 10; i++ {
		m[i] = i + 10
	}

	for k := range m {				//在对字典进行遍历时，遍历的顺序是不固定的


		if k == 5 {
			m[100] = 1000
		}

		delete(m, k)
		fmt.Println(k, m)
	}
}

//运行时会对字典并发操作做出检测。如果某个任务正在对字典进行写操作，那么其他任务就不能对该字典执行并发操作，否则会导致进程崩溃
//可使用sync.RWMutex实现同步，避免读写操作同时进行
func TestSafe2(t *testing.T) {
	var lock sync.RWMutex 				//使用读写锁，以获得最佳性能

	m := make(map[string]int)

	go func() {
		lock.Lock()						//锁的粒度
		m["a"] += 1
		lock.Unlock()					//不能使用defer

		time.Sleep(time.Microsecond)
	}()

	go func() {
		for{
			lock.RLock()
			_ = m["b"]
			lock.RUnlock()

			time.Sleep(time.Microsecond)
		}
	}()

	select {

	}
}

//==================================================性能============================================
//字典对象本身就是指针包装，传参时无须再次取地址
//在创建时预先准备足够的空间有助于提升性能，减少扩张时的内存分配和重新哈希操作
func test(x map[string]int) {
	fmt.Printf("x: %p \n", x)
}

func TestAttr(t *testing.T) {
	m := make(map[string]int)
	test(m)
	fmt.Printf("m: %p, %d \n", m, unsafe.Sizeof(m))
}