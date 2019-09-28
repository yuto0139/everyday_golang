package main

import (
	"fmt"
)

func one(x *int) {
	*x = 1
}

func main() {
	// 34. ポインタ
	// var p *int = &nでの&nで、nのMemoryのAddressを出力
	// また、Addressを代入するのに、ポイント型 *intにする
	// 対して、fmt.Println(*p)での*pで、nのAddressのMemoryを出力
	var n int = 100
	one(&n)
	fmt.Println(n)

	// 35. newとmakeの違い
	var p *int = new(int)
	fmt.Println(*p)
	*p++
	fmt.Println(*p)
}
