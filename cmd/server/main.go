package main

import (
	"fmt"
	"log"
	"net/http"
	"photo_viewer_backend/internal/handler"
	"photo_viewer_backend/internal/repository"
	"photo_viewer_backend/internal/service"
	"github.com/go-chi/chi/v5"
)

func main() {
	// ルーターの作成
	r := chi.NewRouter()

	// ハンドラーの初期化
	photoHandler := handler.NewPhotoHandler(service.NewPhotoService(repository.NewPhotoRepository()))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	// エンドポイントを設定
	r.Get("/top-photos", photoHandler.GetTopPhotos)

	// サーバーを起動
	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
