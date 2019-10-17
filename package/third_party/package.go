package main

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
	"golang.org/x/sync/semaphore"
	"gopkg.in/ini.v1"
)

var s *semaphore.Weighted = semaphore.NewWeighted(1)

// ConfigList ...
type ConfigList struct {
	Port      int
	DbName    string
	SQLDriver string
}

// JsonRPC2 ...
type JsonRPC2 struct {
	Version string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
	Result  interface{} `json:"result,omitempty"`
	Id      *int        `json:"id,omitempty"`
}

// SubscribeParams ...
type SubscribeParams struct {
	Channel string `json:"channel"`
}

// Config ...
var Config ConfigList

func init() {
	cfg, _ := ini.Load("config.ini")
	Config = ConfigList{
		Port: cfg.Section("web").Key("port").MustInt(),
		// MustString()で引数にデフォルト値を設定
		DbName: cfg.Section("db").Key("name").MustString("example.sql"),
		// String()で値がなければ、デフォルト値はnull
		SQLDriver: cfg.Section("db").Key("driver").String(),
	}
}

// longProcess ...
func longProcess(ctx context.Context) {
	// 1つ目のgoroutineを確保してlock
	// if err := s.Acquire(ctx, 1); err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// あるgoroutineの処理途中に他のgoroutineはキャンセル
	isAcquire := s.TryAcquire(1)
	if !isAcquire {
		fmt.Println("Could not get lock")
		return
	}
	// Releaseされて初めて他のgoroutineの処理がスタート
	defer s.Release(1)
	fmt.Println("Wait...")
	time.Sleep(1 * time.Second)
	fmt.Println("Done")
}

func main() {
	// 76. Semaphore
	// goroutineが走っている数を限定
	// TODO() 使うか使わないかわからないけどnilを渡したくない時、空のcontextを作成
	ctx := context.TODO()
	go longProcess(ctx)
	go longProcess(ctx)
	go longProcess(ctx)
	time.Sleep(2 * time.Second)
	go longProcess(ctx)
	time.Sleep(5 * time.Second)

	// 77. iniでConfigの設定ファイルを読み込む
	// iniにより、goではないconfigファイルのデータの読み込みが可能
	fmt.Printf("%T %v\n", Config.Port, Config.Port)
	fmt.Printf("%T %v\n", Config.DbName, Config.DbName)
	fmt.Printf("%T %v\n", Config.SQLDriver, Config.SQLDriver)

	// 79. pubnubでBitcoinの価格をリアルタイムに取得する
	u := url.URL{Scheme: "wss", Host: "ws.lightstream.bitflyer.com", Path: "/json-rpc"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	if err := c.WriteJSON(&JsonRPC2{Version: "2.0", Method: "subscribe", Params: &SubscribeParams{"lightning_ticker_BTC_JPY"}}); err != nil {
		log.Fatal("subscribe:", err)
		return
	}

	for {
		message := new(JsonRPC2)
		if err := c.ReadJSON(message); err != nil {
			log.Println("read:", err)
			return
		}

		if message.Method == "channelMessage" {
			log.Println(message.Params)
		}
	}
}
