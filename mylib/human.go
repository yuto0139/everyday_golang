package mylib

import "fmt"

// 61. PublicとPrivate

// Public ...
var Public = "Public"
var private = "private"

// Person ...
type Person struct {
	Name string
	Age  int
}

// Say ...他のpackageからの呼び出しを許可する場合、関数・変数などをcapitalで
func Say() {
	fmt.Println("Human!")
}
