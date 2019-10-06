package main

import (
	"fmt"
)

func one(x *int) {
	*x = 1
}

// Vertex ...
type Vertex struct {
	X, Y int
	S    string
}

func changeVertex(v Vertex) {
	v.X = 1000
}

// structの場合、参照先もポインタ型にしてくれる (*v.X)
func changeVertex2(v *Vertex) {
	v.X = 1000
}

func main() {
	// 34. ポインタ
	// var p *int = &nでの&nで、nのMemoryのAddressを出力
	// また、Addressを代入するのに、ポイント型 *intにする
	// 対して、fmt.Println(*p)での*pで、nのAddressのMemoryを出力
	var n = 100
	one(&n)
	fmt.Println(n)

	// 35. newとmakeの違い
	// new(int)でメモリの領域を確保 → アドレスを返す
	var p = new(int)
	fmt.Println(*p)
	*p++
	fmt.Println(*p)

	// メモリを確保していないから、nilを返す
	var p2 *int
	fmt.Println(p2)

	// ポインタを返すか返さないかでnewとmakeを使い分け
	s := make([]int, 0)
	fmt.Printf("%T\n", s)

	m := make(map[string]int)
	fmt.Printf("%T\n", m)

	ch := make(chan int)
	fmt.Printf("%T\n", ch)

	var p3 = new(int)
	fmt.Printf("%T\n", p3)

	var st = new(struct{})
	fmt.Printf("%T\n", st)

	// 36. struct
	v := Vertex{X: 1, Y: 2}
	fmt.Println(v)
	fmt.Println(v.X, v.Y)
	v.X = 1000
	fmt.Println(v.X, v.Y)

	// integerおよびstringnの初期値は0
	v2 := Vertex{X: 1}
	fmt.Println(v2)

	// わざわざ宣言しなくても、順番通りに引数を渡して出力可能
	v3 := Vertex{1, 2, "test"}
	fmt.Println(v3)

	v4 := Vertex{}
	fmt.Printf("%T %v\n", v4, v4)

	// sliceやmapはnilを返すが、vertexはnilを返さない
	var v5 Vertex
	fmt.Printf("%T %v\n", v5, v5)

	// Vertexを使ってポインタ型を返す
	v6 := new(Vertex)
	fmt.Printf("%T %v\n", v6, v6)

	// v6の書き方を変更、どちらかというとこちらをよく使う
	// sliceやmapを作成する際は、makeを明示的に使う場合が多い
	v7 := &Vertex{}
	fmt.Printf("%T %v\n", v7, v7)

	// 値渡しをしているため、v8の値は変わらない
	v8 := Vertex{1, 2, "test"}
	changeVertex(v8)
	fmt.Println(v8)

	v9 := &Vertex{1, 2, "test"}
	changeVertex2(v9)
	fmt.Println(v9)
}
