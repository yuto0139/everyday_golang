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

func producer2(first chan int) {
	defer close(first)
	for i := 0; i < 10; i++ {
		first <- i
	}
}

// 引数でchannelの出力・入力を明示的に示す
func multi2(first <-chan int, second chan<- int) {
	defer close(second)
	for i := range first {
		second <- i * 2
	}
}

func multi4(second chan int, third chan int) {
	defer close(third)
	for i := range second {
		third <- i * 4
	}
}

func goroutine4(ch chan string) {
	for i := 0; i < 5; i++ {
		ch <- "packet from 1"
		time.Sleep(1 * time.Second)
	}
}

func goroutine5(ch chan string) {
	for i := 0; i < 5; i++ {
		ch <- "packet from 2"
		time.Sleep(1 * time.Second)
	}
}

// Counter ...
type Counter struct {
	v   map[string]int
	mux sync.Mutex
}

// Inc ... 片方のgoroutineが書き込んでいる最中、他のgoroutineが干渉できないように
func (c *Counter) Inc(key string) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.v[key]++
}

// Value ... 上記の処理の出力用
func (c *Counter) Value(key string) int {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.v[key]
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

	// 54. fan-out fan-in
	first := make(chan int)
	second := make(chan int)
	third := make(chan int)

	go producer2(first)
	go multi2(first, second)
	go multi4(second, third)
	for result := range third {
		fmt.Println(result)
	}

	// 55. channelとselect
	c1 := make(chan string)
	c2 := make(chan string)
	go goroutine4(c1)
	go goroutine5(c2)

	//	複数のgoroutineが同時にmain()に対してchannelを渡す
	for i := 0; i < 10; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}
	// 56. Default Selection と for break
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	// 任意の名前でbreak可能
OuterLoop:
	for {
		select {
		case <-tick:
			fmt.Println(".tick")
		case <-boom:
			fmt.Println("BOOM!")
			break OuterLoop
		// ２つのchannel以外だと、以下の処理を走らせる
		default:
			fmt.Println(".")
			time.Sleep(50 * time.Millisecond)
		}
	}
	fmt.Println("#####################")

	// 57. sync.Mutex
	// channelを使わずして、複数のgoroutineから値を取得するケース
	c3 := Counter{v: make(map[string]int)}
	go func() {
		for i := 0; i < 10; i++ {
			// c3["key"]++
			c3.Inc("Key")
		}
	}()
	go func() {
		for i := 0; i < 10; i++ {
			// c3["key"]++
			c3.Inc("Key")
		}
	}()
	time.Sleep(1 * time.Second)
	fmt.Println(c3, c3.Value("Key"))
}
