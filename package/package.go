package main

import (
	"context"
	"fmt"
	"regexp"
	"sort"
	"time"

	"../mylib"
	"../mylib/under"
)

const (
	c1 = iota
	c2
	c3
)

const (
	_ = iota
	// KB ...
	KB int = 1 << (10 * iota)
	// MB ...
	MB
	// GB ...
	GB
)

func longProcess(ctx context.Context, ch chan string) {
	fmt.Println("run")
	time.Sleep(2 * time.Second)
	fmt.Println("finish")
	ch <- "result"
}

// 60. package
func main() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(mylib.Average(s))

	mylib.Say()
	under.Hello()
	person := mylib.Person{Name: "Mike", Age: 20}
	fmt.Println(person)

	fmt.Println(mylib.Public)
	// fmt.Println(mylib.private)

	// 64. サードパーティーのpackageのインストール
	// spy, _ := quote.NewQuoteFromYahoo("spy", "2016-01-01", "2016-04-01", quote.Daily, true)
	// fmt.Print(spy.CSV())
	// rsi2 := talib.Rsi(spy.Close, 2)
	// fmt.Println(rsi2)

	// 67. time
	t := time.Now()
	fmt.Println(t)
	fmt.Println(t.Format(time.RFC3339))
	fmt.Println(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())

	// 68. regex
	match, _ := regexp.MatchString("a([a-z]+)e", "apple")
	fmt.Println(match)

	// 何度も同じ正規表現を使用する場合、一度変数代入して処理を高速に
	r := regexp.MustCompile("a([a-z]+)e")
	ms := r.MatchString("apple")
	fmt.Println(ms)

	// s := "/view/test"
	r2 := regexp.MustCompile("^/(edit|save|view)/([a-zA-z0-9]+)$")
	fs := r2.FindString("/view/test")
	fmt.Println(fs)
	// 個別に要素を取り出したい場合
	fss := r2.FindStringSubmatch("/view/test")
	fmt.Println(fss, fss[0], fss[1], fss[2])
	fss = r2.FindStringSubmatch("/edit/test")
	fmt.Println(fss, fss[0], fss[1], fss[2])
	fss = r2.FindStringSubmatch("/save/test")
	fmt.Println(fss, fss[0], fss[1], fss[2])

	// 69. sort
	i := []int{5, 3, 2, 8, 7}
	s2 := []string{"d", "a", "f"}
	p := []struct {
		Name string
		Age  int
	}{
		{"Nancy", 20},
		{"Vera", 40},
		{"Mike", 30},
		{"Bob", 50},
	}
	fmt.Println(i, s2, p)
	sort.Ints(i)
	sort.Strings(s2)
	// forループを2つの変数を使って廻している
	sort.Slice(p, func(i, j int) bool { return p[i].Name < p[j].Name })
	sort.Slice(p, func(i, j int) bool { return p[i].Age < p[j].Age })
	fmt.Println(i, s2, p)

	// 70. iota
	fmt.Println(c1, c2, c3)
	fmt.Println(KB, MB, GB)

	// 71. context
	ch := make(chan string)
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()
	// goroutineの処理が長くかかりすぎたときにキャンセルできる
	go longProcess(ctx, ch)
	// cancel()を使ってgoroutineをキャンセルする方法もある
	cancel()

CTXLOOP:
	for {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			break CTXLOOP
		case <-ch:
			fmt.Println("success")
			break CTXLOOP
		}
	}
	fmt.Println("###########")
}
