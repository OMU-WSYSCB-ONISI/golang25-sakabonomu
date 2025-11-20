package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	// 既存のハンドラ
	http.HandleFunc("/hello", hellohandler)

	// 課題追加分
	http.HandleFunc("/webfortune", fortuneHandler)
	http.HandleFunc("/info", infoHandler)

	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}

// 既存のハンドラ
func hellohandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "こんにちは from Codespace !")
}

// 課題3: Webおみくじ
func fortuneHandler(w http.ResponseWriter, r *http.Request) {
	// 運勢のスライス
	fortunes := []string{"大吉", "中吉", "小吉", "吉", "凶"}

	// 乱数の初期化
	seed := time.Now().UnixNano()
	rng := rand.New(rand.NewSource(seed))

	// ランダムにインデックスを決定
	idx := rng.Intn(len(fortunes))

	fmt.Fprintf(w, "今の運勢は%sです", fortunes[idx])
}

// 課題4: 時刻とブラウザ情報
func infoHandler(w http.ResponseWriter, r *http.Request) {
	// 時刻の取得 (JST)
	jst, _ := time.LoadLocation("Asia/Tokyo")
	curTime := time.Now().In(jst).Format("15:04:05")

	// ブラウザ情報 (User-Agent) の取得
	// r.Headerは map[string][]string 型である
	userAgent := r.Header["User-Agent"][0]

	fmt.Fprintf(w, "今の時刻は%sで，利用しているブラウザは%sですね", curTime, userAgent)
}
