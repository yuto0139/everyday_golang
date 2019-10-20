package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	// 0600でWebサーバーを立ち上げたユーザーへファイルの読み書き権限付与
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, _ := template.ParseFiles("web_app/" + tmpl + ".html")
	t.Execute(w, p)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	// アクセスしたURLがrに格納 (ここでは、/view/test)
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title)
	// fmt.Fprintf(w, "<h1>%s</h1>", p.Title)
	// 87. Web Applications 3 - templateとhttp.ResponseWriterとhttp.Request
	// 引数が長くなってしまうので、テンプレート化
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	// アクセスしたURLがrに格納 (ここでは、/view/test)
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		// 新規ページを作成
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func main() {
	// 85. Web Applications 1 - ioutil
	p1 := &Page{Title: "test", Body: []byte("This is a sample Page.")}
	p1.save()

	p2, _ := loadPage(p1.Title)
	fmt.Println(string(p2.Body))

	// 86. Web Applications 2 - http.ListenAndServer
	// URL(handler)の登録
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	// サーバー起動 (第2引数がnilの場合、デフォルトのpage not foundが返り値)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
