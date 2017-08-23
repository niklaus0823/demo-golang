/**
 * select
 * -------------
select 语句使得一个 goroutine 在多个通讯操作上等待。

select 会阻塞，直到条件分支中的某个可以继续执行，这时就会执行那个条件分支。当多个都准备好的时候，会随机选择一个。

select 就是监听 IO 操作，当 IO 操作发生时，触发相应的动作，

在执行select语句的时候，运行时系统会自上而下地判断每个case中的发送或接收操作是否可以被立即执行(立即执行：意思是当前Goroutine不会因此操作而被阻塞)

每个case语句里必须是一个IO操作，确切的说，应该是一个面向channel的IO操作。
*/
package main

import (
	"fmt"
	"time"
)

func f1(ch chan int) {
	time.Sleep(time.Second * 5)
	ch <- 1
}
func f2(ch chan int) {
	time.Sleep(time.Second * 10)
	ch <- 1
}

func main() {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)

	go f1(ch1)
	go f2(ch2)

	select {  // 阻塞进程，直到 5 秒后，ch1 通道收到数据。打印并退出
	case <-ch1:
		fmt.Println("The first case is selected.")
	case <-ch2:
		fmt.Println("The second case is selected.")
	}
}
