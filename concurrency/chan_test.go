package concurrency

import (
	"sync"
	"testing"
)

//===========================================收发==============================================
//对于循环接受数据，range模式更简洁一些
func TestChan1(t *testing.T) {

	done := make(chan struct{})												//用作同步通知
	c := make(chan int)														//用作数据收发

	go func() {
		defer close(done)													//发送信号

		for x := range c {
			println(x)
		}
	}()

	c <- 1
	c <- 2
	c <- 3
	close(c)																//发送信号，range循环终止

	<-done
}

//==============================================单向=====================================================
//通道默认是双向的，我们可以设置收发操作的方向
func TestChan2(t *testing.T) {

	var wg sync.WaitGroup
	wg.Add(2)

	c := make(chan int)

	var send chan<- int = c
	var recv <-chan int = c

	go func() {
		defer wg.Done()

		for x := range recv{
			println(x)

		}
	}()

	go func() {
		defer wg.Done()
		defer close(c)

		for i := 0; i < 3; i++ {
			send <- i
		}
	}()

	wg.Wait()

}

//=======================================================select==========================================
//如果需要同时处理多个通道，可以选用select语句，他会随机选择一个可用通道做收发操作
//如要去全部通道处理结束，可将已完成通道设置为nil,这样它就会被阻塞，不再被select选中
//即便是同一通道，也会随机选择case
//当所有通道都不可用的时，select会执行default语句，如此磕避免select阻塞
func TestChan3(t *testing.T) {

	var wg sync.WaitGroup
	wg.Add(3)

	a, b := make(chan int), make(chan int)


	go func() {
		defer wg.Done()

		for true {
			select {
			case x, ok := <-a:
				if !ok {											//如果通道关闭，则设置为nil，阻塞
					a = nil
					break
				}

				println("a", x)

			case x, ok := <-b:
				if !ok {
					b = nil
					break
				}

				println("b", x)

			default:
				if a == nil && b == nil {
					return
				}
			}



		}
	}()

	go func() {

		defer wg.Done()
		defer close(a)

		for i := 0; i < 3; i++ {
			a <- i
		}

	}()


	go func() {
		defer wg.Done()
		defer close(b)

		for i := 0; i < 5; i++ {
			b <- i * 10

		}

	}()

	wg.Wait()

}