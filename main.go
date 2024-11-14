package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	// 新しいルーターを作成
	r := chi.NewRouter()

	// ルートにハンドラを設定
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	// /helloエンドポイントを追加
	r.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "こんにちは！")
	})

	// サーバーを起動
	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
