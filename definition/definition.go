package main

// 必要なライブラリ
import (
	"fmt"
	"os/user"
	"strconv"
	"strings"
	"time"
)

// 初期の実行
func init() {
	fmt.Println("Init!")
}

func bazz() {
	fmt.Println("Bazz")
}

func add(x, y int) (int, int) {
	return x + y, x - y
}

// named return values 名前付き戻り値
// メリットは、どんな戻り値なのか変数名ですぐわかる
func cal(price, item int) (result int) {
	result = price * item
	return
}

// 関数mainから実行が開始されると定められている -> エントリーポイント
func main() {
	// 8.変数宣言
	bazz()
	fmt.Println("Hello world!", time.Now())
	fmt.Println(user.Current())
	// 型推論が行われ、型指定が不要
	var (
		i    = 1
		f64  = 1.2
		s    = "test"
		t, f = true, false
	)
	fmt.Println(i, f64, s, t, f)

	// 関数定義の外部に定義された変数は、同一パッケージであればどこでも参照可 -> パッケージ変数
	xi := 1
	xi = 2
	xf64 := 1.2
	xs := "test"
	xt, xf := true, false
	fmt.Println(xi, xf64, xs, xt, xf)
	fmt.Printf("%T\n", xf64)
	fmt.Printf("%T\n", xi)

	// 9.定数宣言
	const Pi = 3.14
	const (
		// 型なし定数
		Username = "test_user"
		// 型あり定数
		Password string = "test_pass"
	)

	fmt.Println(Pi, Username, Password)

	// 10.数値型
	var (
		u8 uint8 = 255
	)
	fmt.Printf("type=%T value=%v", u8, u8)
	x := 0
	fmt.Println(x)
	x++
	fmt.Println(x)

	// 11.文字列型
	fmt.Println("Hello World!")

	var ss = "Hello World!"
	ss = strings.Replace(ss, "H", "X", 1)
	fmt.Println(ss)

	fmt.Println(strings.Contains(ss, "World!"))

	fmt.Println(`Test
		Test
									Test`)
	// 12.論理値型
	tt, ff := true, false
	fmt.Printf("%T %v\n", tt, tt)
	fmt.Printf("%T %v\new", ff, ff)

	// 13.型変換
	var xx = 1
	xxx := float64(xx)
	fmt.Printf("%T %v %f\n", xxx, xxx, xxx)

	// strcov 文字列の型変更のライブラリ
	// Atoi ascii to integer
	var sss = "14"
	i, _ = strconv.Atoi(sss)
	fmt.Printf("%T %v\n", i, i)

	// 14.配列
	var a [2]int
	a[0] = 100
	a[1] = 200
	fmt.Println(a)

	var b = [2]int{100, 200}
	// b = append(b, 300)
	fmt.Println(b)

	// 15.スライス
	n := []int{1, 2, 3, 4}
	fmt.Println(n)

	n[2] = 100
	fmt.Println(n)

	n = append(n, 100, 200, 300)
	fmt.Println(n)

	var board = [][]int{
		[]int{0, 1, 2},
		[]int{3, 4, 5},
		[]int{6, 7, 8},
	}
	fmt.Println(board)

	// 16.スライスのmake, cap
	// makeでlengthとcapacityを指定
	nn := make([]int, 3, 5)
	fmt.Printf("len=%d cap=%d value=%v\n", len(nn), cap(nn), nn)

	var c []int
	// c = make([]int, 5)
	c = make([]int, 0, 5)
	for i := 0; i < 5; i++ {
		c = append(c, i)
		fmt.Println(c)
	}
	fmt.Println(c)

	// 17.map
	m := map[string]int{"apple": 100, "banana": 200}
	fmt.Println(m)
	fmt.Println(m["apple"])
	m["banana"] = 300
	fmt.Println(m)

	fmt.Println(m["nothing"])

	v, ok := m["apple"]
	fmt.Println(v, ok)

	v2, ok2 := m["nothing"]
	fmt.Println(v2, ok2)

	m2 := make(map[string]int)
	m2["pc"] = 5000
	fmt.Println(m2)

	// varで値を入れずにスライスまたはmapを宣言 -> 値はnil
	var nothing []int
	if nothing == nil {
		fmt.Println("nil")
	}

	// 19.関数型
	r1, r2 := add(10, 20)
	fmt.Println(r1, r2)

	r3 := cal(100, 2)
	fmt.Println(r3)

	function := func(x int) {
		fmt.Println("inner func", x)
	}
	function(1)
	func(x int) {
		fmt.Println("inner func", x)
	}(2)
}
