package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// T ...
type T struct{}

// Person ...
type Person struct {
	// json: "hogehoge" でmarchalするときの名前を設定可能、型も変更可能
	// json: "-" もしくは、omitemptyで非表示
	Name      string   `json: "-"`
	Age       int      `json: "age,omitempty"`
	Nicknames []string `json: "nicknames"`
	T         *T       `json: "T, omitempty"`
}

// UnmarshalJSON ... Unmarshalのカスタマイズ
func (p *Person) UnmarshalJSON(b []byte) error {
	type Person2 struct {
		Name string
	}
	var p2 Person2
	err := json.Unmarshal(b, &p2)
	if err != nil {
		fmt.Println(err)
	}
	p.Name = p2.Name + "!"
	return err
}

// MarshalJSON ... Marshalのカスタマイズ
func (p Person) MarshalJSON() ([]byte, error) {
	// a := &struct {Name string}{Name: "test"}
	v, err := json.Marshal(&struct {
		Name string
	}{
		Name: "Mr." + p.Name,
	})
	return v, err
}

// DB ...
var DB = map[string]string{
	"User1Key": "User1Secret",
	"User2Key": "User2Secret",
}

// Server ...
func Server(apiKey, sign string, data []byte) {
	apiSecret := DB[apiKey]
	h := hmac.New(sha256.New, []byte(apiSecret))
	h.Write(data)
	expectedHMAC := hex.EncodeToString(h.Sum(nil))
	// クライアントからのsignとサーバーで生成したexpectedHMACが一致するかどうか
	fmt.Println(sign == expectedHMAC)
}

func main() {
	resp, _ := http.Get("http://example.com")
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

	// Parseを使って、正しいURLかどうか認証
	base, _ := url.Parse("https://example.com/")
	reference, _ := url.Parse("/test?a=1&b=2")

	// ResolveReferenceを使って、エンドポイントを含むURLを生成
	endpoint := base.ResolveReference(reference).String()
	fmt.Println(base)
	fmt.Println(endpoint)

	// req, _ := http.NewRequest("Get", endpoint, nil)
	req, _ := http.NewRequest("POST", endpoint, bytes.NewBuffer([]byte("password")))
	req.Header.Add("If-None-Match", `W/"xyz"`)
	q := req.URL.Query()
	q.Add("c", "3&%")
	fmt.Println(q)
	fmt.Println(q.Encode())
	req.URL.RawQuery = q.Encode()

	var client *http.Client = &http.Client{}
	resp2, _ := client.Do(req)
	body2, _ := ioutil.ReadAll(resp2.Body)
	fmt.Println(string(body2))

	// 74. json.UnmarshalとMarshalとエンコード
	b := []byte(`{"name": "mike", "age": 0, "nicknames": ["a", "b", "c"]}`)
	var p Person
	// Unmarshalを使ってjsonからstructのkeyに合った値を入力
	if err := json.Unmarshal(b, &p); err != nil {
		fmt.Println(err)
	}
	fmt.Println(p.Name, p.Age, p.Nicknames)

	// jsonに変換して、Network越しへ出力
	v, _ := json.Marshal(p)
	fmt.Println(string(v))

	// 75. hmacでAPI認証
	// クライアントから送られたデータが正しいクライアントからかどうかをサーバー側で認証
	const apiKey = "User1Key"
	const apiSecret = "User1Secret"

	data := []byte("important_data")
	h := hmac.New(sha256.New, []byte(apiSecret))
	h.Write(data)
	sign := hex.EncodeToString(h.Sum(nil))

	fmt.Println(sign)

	// 3つの情報をサーバー側へ送信
	Server(apiKey, sign, data)
}
