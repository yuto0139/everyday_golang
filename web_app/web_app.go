package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
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

var templates = template.Must(template.ParseFiles("web_app/edit.html", "web_app/view.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	// t, _ := template.ParseFiles("web_app/" + tmpl + ".html")
	// t.Execute(w, p)
	// 上記の代わりにキャッシュを利用
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	// アクセスしたURLがrに格納 (ここでは、/view/test)
	// title := r.URL.Path[len("/view/"):]
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	// fmt.Fprintf(w, "<h1>%s</h1>", p.Title)
	// 87. Web Applications 3 - templateとhttp.ResponseWriterとhttp.Request
	// 引数が長くなってしまうので、テンプレート化
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		// 新規ページを作成
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

// 89. Web Applications 5 - templateのキャッシュとハンドラー
// titleを取得の上、各handlerへの処理へ行く
func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

func main() {
	// 85. Web Applications 1 - ioutil
	p1 := &Page{Title: "test", Body: []byte("This is a sample Page.")}
	p1.save()

	p2, _ := loadPage(p1.Title)
	fmt.Println(string(p2.Body))

	// 86. Web Applications 2 - http.ListenAndServer
	// URL(handler)の登録
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
	// サーバー起動 (第2引数がnilの場合、デフォルトのpage not foundが返り値)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
