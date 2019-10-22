package main

import "fmt"

// Vertex ... 構造体を定義
type Vertex struct {
	// フィールド
	x, y int
}

// Area ... レシーバーで構造体とメソッドを紐付け シンプルに値を渡す
func (v Vertex) Area() int {
	return v.x * v.y
}

// Area ... こちらでも使えるが、上記の関数の方がオブジェクトぽく使える
// func Area(v Vertex) int {
// 	return v.x * v.y
// }

// Scale ...ポインタレシーバー structの中身を変更可能
// デフォルトではVertexには値渡し(コピー)でメソッドに渡される
// *をつけることで参照渡しになる
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
	Age  int
}

// Say ...
func (p *Person) Say() string {
	p.Name = "Mr." + p.Name
	fmt.Println(p.Name)
	return p.Name
}

// DriveCar ...
// Humanインターフェースで指定したメソッドsay()を実装する必要あり
func DriveCar(human Human) {
	if human.Say() == "Mr.Mike" {
		fmt.Println("Run")
	} else {
		fmt.Println("Get out")
	}
}

func do1(i interface{}) {
	// inteface型は全ての型と互換性を持つ
	// 使用するには、型の上書きをする必要がある -> Type Assertion
	ii := i.(int)
	ii *= 2
	fmt.Println(ii)
}

// 様々な型の呼び出しに対応
func do2(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println(v * 2)
	case string:
		fmt.Println(v + "!")
	default:
		fmt.Printf("I don't know %T\n", v)
	}
}

// structを使用した出力結果を自由に変更できる
func (p Person) String() string {
	// return "My name is" + p.Nameと等しい
	return fmt.Sprintf("My name is %v.", p.Name)
}

// UserNotFound ...
type UserNotFound struct {
	Username string
}

func (e *UserNotFound) Error() string {
	return fmt.Sprintf("User not found: %v", e.Username)
}

func myFunc() error {
	// Something wrong
	ok := false
	if ok {
		return nil
	}
	return &UserNotFound{Username: "mike"}
}

func main() {
	// 39. メソッドとポインタレシーバーと値レシーバー
	// v := Vertex{3, 4}
	// fmt.Println(Area(v))

	// 40. コンストラクタ
	// 他のパッケージからstructのインスタンスを生成するときは、Newを使用
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
	var mike Human = &Person{"Mike", 22}
	var x Human = &Person{"X", 22}
	DriveCar(mike)
	DriveCar(x)

	// 44. Type Assertionとswitch
	var i interface{} = 10
	do1(i)
	do1(10)
	do2(10)
	do2("Mike")
	do2(true)

	// 45. Stringer
	// 出力方法(文字列)を変更可能 mark, 22 -> "My name is Mike
	mark := Person{"Mark", 22}
	fmt.Println(mark)

	// 46. カスタムエラー
	if err := myFunc(); err != nil {
		fmt.Println(err)
	}
}
