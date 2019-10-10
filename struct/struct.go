package main

import "fmt"

// Vertex ...
type Vertex struct {
	x, y int
}

// Area ... 値レシーバー シンプルに値を渡す
func (v Vertex) Area() int {
	return v.x * v.y
}

// Area ... こちらでも使えるが、上記の関数の方がオブジェクトぽく使える
// func Area(v Vertex) int {
// 	return v.x * v.y
// }

// Scale ...ポインタレシーバー structの中身を変更可能
func (v *Vertex) Scale(i int) {
	v.x = v.x * i
	v.y = v.y * i
}

// Vertex3D ... Vertexを埋め込みしているため、x, yを使用可能
type Vertex3D struct {
	Vertex
	z int
}

// Area3D ...
func (v Vertex3D) Area3D() int {
	return v.x * v.y * v.z
}

// Scale3D ...
func (v *Vertex3D) Scale3D(i int) {
	v.x = v.x * i
	v.y = v.y * i
	v.z = v.z * i
}

// New ... 初期化
func New(x, y, z int) *Vertex3D {
	return &Vertex3D{Vertex{x, y}, z}
}

// MyInt ...
type MyInt int

// Double ...
func (i MyInt) Double() int {
	fmt.Printf("%T %v\n", i, i)
	fmt.Printf("%T %v\n", 1, 1)
	return int(i * 2)
}

// Human ...
type Human interface {
	Say() string
}

// Person ...
type Person struct {
	Name string
}

// Say ...
func (p *Person) Say() string {
	p.Name = "Mr." + p.Name
	fmt.Println(p.Name)
	return p.Name
}

// DriveCar ...
// Humanインターフェースで指定したメソッドsay()を実装する必要がある -> 型に近い
func DriveCar(human Human) {
	if human.Say() == "Mr.Mike" {
		fmt.Println("Run")
	} else {
		fmt.Println("Get out")
	}
}
func main() {
	// 39. メソッドとポインタレシーバーと値レシーバー
	// v := Vertex{3, 4}
	// fmt.Println(Area(v))

	// 40. コンストラクタ
	v := New(3, 4, 5)
	v.Scale(10)
	fmt.Println(v.Area())
	// 41. Embedded
	fmt.Println(v.Area3D())

	// 42. non-struct
	// 自分が作ったtypeに対して、Double()といった関数をもたせることができる
	myInt := MyInt(10)
	fmt.Println(myInt.Double())

	// 43. インターフェースとダッグタイピング
	var mike Human = &Person{"Mike"}
	var x Human = &Person{"X"}
	DriveCar(mike)
	DriveCar(x)
}
