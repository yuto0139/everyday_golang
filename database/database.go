package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// DbConnection ...
var DbConnection *sql.DB

// Person ...
type Person struct {
	Name string
	Age  int
}

func main() {
	// 84. データベース操作
	// データベースのコネクションをオープン
	DbConnection, _ := sql.Open("sqlite3", "./example.sql")
	defer DbConnection.Close()
	cmd := `CREATE TABLE IF NOT EXISTS person(
              name STRING,
              age INT)`
	// 作成したコマンドの実行
	// データベースの情報を返す必要がないので、`_`へ代入
	_, err := DbConnection.Exec(cmd)
	if err != nil {
		log.Fatalln(err)
	}

	cmd = "INSERT INTO person (name, age) VALUES (?, ?)"
	// 引数 "Mike", 24がコマンドの?へ代入
	_, err = DbConnection.Exec(cmd, "Mike", 24)
	if err != nil {
		log.Fatalln(err)
	}

	// '?'でsqlインジェクション防止 (不要な値はエスケープしてくれる)
	cmd = "UPDATE person SET age = ? WHERE name = ?"
	_, err = DbConnection.Exec(cmd, 30, "Mike")
	if err != nil {
		log.Fatalln(err)
	}

	// マルチセレクト
	cmd = "SELECT * FROM person"
	// レコードを取得するときはdb.Query(...)
	rows, _ := DbConnection.Query(cmd)
	// 処理が終わったらカーソルを閉じる
	defer rows.Close()
	var pp []Person
	for rows.Next() {
		var p Person
		// scanがstructの値にデータを挿入
		err := rows.Scan(&p.Name, &p.Age)
		// 一つずつエラーハンドリング
		if err != nil {
			log.Println(err)
		}
		// 一つ一つのインスタンスをpp(スライス)へ挿入
		pp = append(pp, p)
	}
	// 一括でエラーハンドリングも可能
	err = rows.Err()
	if err != nil {
		log.Fatalln(err)
	}
	for _, p := range pp {
		fmt.Println(p.Name, p.Age)
	}

	// シングルセレクト
	cmd = "SELECT * FROM person where age = ?"
	// １件取得するときはdb.QueryRow(...)
	row := DbConnection.QueryRow(cmd, 20)
	var p Person
	err = row.Scan(&p.Name, &p.Age)
	if err == sql.ErrNoRows {
		// データがnilの場合
		log.Println("No row")
	} else {
		// 例外発生の場合
		log.Println(err)
	}
	fmt.Println(p.Name, p.Age, "シングルセレクトの結果")

	cmd = "DELETE FROM person WHERE name = ?"
	_, err = DbConnection.Exec(cmd, "Mike")
	if err != nil {
		log.Fatalln(err)
	}
}
