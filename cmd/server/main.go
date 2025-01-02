package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
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
	r.Get("/photo", photoHandler.GetPhoto)

	// サーバーを起動
	port := "8080"
	if envPort := os.Getenv("BACKEND_PORT"); envPort != "" {
		port = envPort
	}
	fmt.Printf("Starting server on :%s\n", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
}
