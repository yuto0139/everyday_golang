package main

import (
	"fmt"
	"sync"
	"time"
)

// 49. goroutineとsync.WaitGroup
// wgのポインタを宣言 wg *sync.WaitGroup
func goroutine(s string, wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		// time.Sleep(300 * time.Millisecond)
		fmt.Println(s)
	}
	// 処理が終わったことを示すwg.Done()
	wg.Done()
}

func normal(s string) {
	for i := 0; i < 5; i++ {
		// time.Sleep(300 * time.Millisecond)
		fmt.Println(s)
	}
}

func goroutine1(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func goroutine2(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func goroutine3(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
		c <- sum
	}
	close(c)
}

func producer(ch chan int, i int) {
	// something
	ch <- i * 2
}

func consumer(ch chan int, wg *sync.WaitGroup) {
	for i := range ch {
		fmt.Println("process", i*1000)
		wg.Done()
	}
	fmt.Println("###################")
}

func main() {
	// goroutineの処理が完了するのを待つsync.WaitGroup
	var wg sync.WaitGroup
	// 一つの並列処理があることをwgに伝える
	wg.Add(1)
	// goroutineにwgにアドレスを引数として渡す
	go goroutine("world", &wg)
	normal("hello")
	//	wgでの並列処理が終了するまで待ってもらう
	wg.Wait()

	// 50. channel
	s := []int{1, 2, 3, 4, 5}
	c := make(chan int, len(s))
	go goroutine1(s, c)
	go goroutine2(s, c)
	x := <-c
	fmt.Println(x)
	y := <-c
	fmt.Println(y)

	// 51. Buffered Channels
	// 第2引数でbufferの数を指定
	ch := make(chan int, 2)
	ch <- 100
	fmt.Println(len(ch))
	ch <- 200
	fmt.Println(len(ch))
	close(ch)

	// rangeでループ処理する際、closeでchannelの終了を教えてあげる
	for c := range ch {
		fmt.Println(c)
	}

	// 52. channelのrangeとclose
	//	goroutine3からmain()へ繰り返し値を送信＆終わったらclose
	go goroutine3(s, c)
	for i := range c {
		fmt.Println(i)
	}

	// 53. producerとconsumer
	var wg2 sync.WaitGroup
	ch2 := make(chan int)

	// producer
	for i := 0; i < 10; i++ {
		wg2.Add(1)
		go producer(ch2, i)
	}

	//	consumer
	go consumer(ch2, &wg2)
	wg2.Wait()
	close(ch2)
	time.Sleep(2 * time.Second)
	fmt.Println("Done")
}
