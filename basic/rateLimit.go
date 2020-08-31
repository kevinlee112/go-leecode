package basic

import (
	"fmt"
	"sync"
	"time"
)
var (
	working chan int //goroutine计数器 用于限制最大并发数
	wg      sync.WaitGroup
)

func main() {
	jobList := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} //要工作的任务清单

	//每秒3个请求 最大同时4个请求
	duration := time.Second
	concurrency := 3
	concurrencyMax := 4

	ticker := time.NewTicker(duration)
	working = make(chan int, concurrencyMax)

	//通过限定1个周期派发3个任务来实现限制请求次数
	k := 0 //用于控制周期内发送次数
	for c, job := range jobList {
		working <- c //计数器+1 可能会发生阻塞

		wg.Add(1)
		go work(job)

		k++
		if k == concurrency {
			<-ticker.C //等待一个周期 可能会白等
			k = 0
		}
	}
	wg.Wait()
}

func work(j int64) {
	defer wg.Done()

	fmt.Println("doing work#", j)
	<-time.After(5 * time.Second) //假设5秒完成

	//工作完成后计数器减1
	<-working
}