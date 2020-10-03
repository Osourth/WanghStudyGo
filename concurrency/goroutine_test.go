package concurrency

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

//========================================================并发与并行==========================================================
//并发： 逻辑上具有同时处理多个任务的能力
//并行：物理上在同一时刻执行多个并发任务
//关键字 go 并非执行并发操作，而是创建一个并发任务单元、新建任务被放置在系统队列中，等待调度器安排合适系统进程去获取执行权
//与defer一样，goroutine也会因 “延迟执行” 而立即计算并复制执行参数

var c int

func counter() int {
	c++
	return c
}

func TestGoroutine(t *testing.T) {

	a := 100

	go func(x , y int) {
		time.Sleep(time.Second)
		println("go: ", x, y)
	}(a, counter())										//立即计算并复制参数

	a += 100
	println("main: ", a, counter())

	time.Sleep(time.Second * 3)							//等待goroutine运行结束


}


//=============================================================进程退出时，不会等待并发任务结束，可用通道（channel）阻塞，然后发出推出信号

func TestWait1(t *testing.T) {

	exit := make(chan struct{})							//创建通道，因为仅是通知，数据并没有实际意义

	go func() {
		time.Sleep(time.Second)
		println("goroutine done")

		close(exit)										//关闭通道，发出信号。除关闭通道外，写入数据也可以解除阻塞
	}()

	println("main ...")
	<-exit												//如通道关闭，立即接触阻塞
	println("main exit ")
}

//如果等待多个任务结束，推荐使用sync.WaitGroup。通过设定计数器，让每个goroutine在退出前递减，直至归零时接触阻塞
func TestWait2(t *testing.T) {

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {

		wg.Add(1)									//累加计数，推荐在 go 之外进行设置

		go func(id int) {
			defer wg.Done()								//递减计数
			time.Sleep(time.Second)
			println("goroutine", id, "done")
		}(i)

	}

	println("main ...")
	wg.Wait()											//阻塞，直到计数归零
	println("main exit")
}

//可多处使用Wait阻塞，他们都能收到消息
func TestWait3(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		wg.Wait()										//等待归零
		println("wait exit!")

	}()

	go func() {
		time.Sleep(time.Second)
		println("done")

		wg.Done()										//递减计数
	}()

	wg.Wait()
	println("main.exit")


}

//===============================================================Local Storage========================================
//设置局部存储，如果使用map作为局部存储容器，注意做同步处理，因为运行时会对其做并发处理
func TestTLS(t *testing.T) {
	var wg sync.WaitGroup

	var gs [5]struct{									//用于实现类似TLS功能
		id 	int
		result int
	}

	for i := 0; i < len(gs); i++ {
		wg.Add(1)

		go func(id int) {								//使用参数避免闭包延迟求值
			defer wg.Done()

			gs[id].id = id
			gs[id].result = (id + 1) * 100

		}(i)

	}

	wg.Wait()
	fmt.Printf("%+v\n", gs)

}


//====================================================暂停和终止=======================================================
//Gosched()		暂停，释放线程去执行其他任务。当前任务被放回队列，等待下次调度时恢复执行
func TestGosched(t *testing.T) {

	runtime.GOMAXPROCS(1)								//设置多少个线程参与并发任务执行，默认与处理器核数相同
	exit := make(chan struct{})

	go func() {

		defer close(exit)

		go func() {
			println("b")
		}()

		for i := 0; i < 4; i++ {
			println("a:", i)

			if i == 1 {										//让出当前线程，调度执行b
				runtime.Gosched()
			//	runtime.Goexit()							//立即终止整个调用堆栈，如果在main.main里调用Goexit，他会等待其他任务结束，然后让进程直接崩溃
			}
		}

	}()

	<-exit
}
