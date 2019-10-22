package main

// 必要なライブラリ
import (
	"fmt"
	"log"
	"os"
	"time"
)

func by2(num int) string {
	if num%2 == 0 {
		return "ok"
	}
	return "no"
}

func getOsName() string {
	return "mac"
}

func main() {
	// 24. if文
	result := by2(10)
	if result == "ok" {
		fmt.Println("great")
	}
	// 1行で書く
	if result2 := by2(10); result2 == "ok" {
		fmt.Println("great2")
	}

	num := 9
	if num%2 == 0 {
		fmt.Println("by 2")
	} else if num%3 == 0 {
		fmt.Println("by 3")
	} else {
		fmt.Println("else")
	}

	x, y := 11, 12
	if x == 10 && y == 10 {
		fmt.Println("&&")
	}

	if x == 10 || y == 10 {
		fmt.Println("||")
	}
	// 25. for文
	for i := 0; i < 10; i++ {
		if i == 3 {
			fmt.Println("continue")
			continue
		}
		if i > 5 {
			fmt.Println("break")
			break
		}
		fmt.Println(i)
	}
	// for文の省略記法
	sum := 1
	for sum < 10 {
		sum += sum
		fmt.Println(sum)
	}
	// 26. range
	l := []string{"python", "go", "java"}

	for i := 0; i < len(l); i++ {
		fmt.Println(i, l[i])
	}
	// 要素分だけ何らかの処理を繰り返し実行
	for i, v := range l {
		fmt.Println(i, v)
	}
	// i(インデックス)が不要な時
	for _, v := range l {
		fmt.Println(v)
	}

	// mapを使う場合
	m := map[string]int{"apple": 100, "banana": 200}
	// forと同様に展開できる
	for k, v := range m {
		fmt.Println(k, v)
	}
	for k := range m {
		fmt.Println(k)
	}
	for _, v := range m {
		fmt.Println(v)
	}

	// 27. switch文
	// 省略記法 (変数osをswitch文の外へ使わない場合)
	switch os := getOsName(); os {
	case "mac":
		fmt.Println("Mac!!")
	case "windows":
		fmt.Println("Windows!!")
	default:
		fmt.Println("Default!!")
	}

	// 条件を書かなくてもOkな場合もある
	t := time.Now()
	fmt.Println(t)
	switch {
	case t.Hour() < 12:
		fmt.Println("Morning!!")
	case t.Hour() < 17:
		fmt.Println("Afternoon!!")
	}

	// 28. defer
	// 特定の処理を遅らせることができる
	// 他のディレクトリを読み込む時に、closeし忘れないようにdeferを使うとか

	// 29. log
	// _, err := os.Open("dfhdfhaafk")
	// if err != nil {
	// 	log.Fatalln("Exit", err)
	// }

	log.Println("logging!")
	log.Printf("%T %v", "test", "test")

	// log.Fatalでコンパイルが終了する
	// log.Fatalf("%T %v", "test", "test")
	// log.Fatalf("error!!")

	// 30. error handling
	file, err := os.Open("./definition.go")
	if err != nil {
		log.Fatalln("Error!")
	}
	defer file.Close()
	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		log.Fatalln("Error!")
	}
	fmt.Println(count, string(data))

	// 31. panicとrecover
	// 実行を継続できないランタイムエラー（配列の範囲外アクセスなど）が発生した場合には、パニック (panic) を発生させる仕組み
	// パニックが発生すると、デフォルトではプログラム全体が終了

}
