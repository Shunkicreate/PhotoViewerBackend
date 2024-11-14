package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"photo_viewer_backend/internal/handler"
	"photo_viewer_backend/internal/repository"
	"photo_viewer_backend/internal/service"
)

func main() {
	// 依存関係を初期化
	photoRepo := repository.NewPhotoRepository()
	photoService := service.NewPhotoService(photoRepo)
	photoHandler := handler.NewPhotoHandler(photoService)

	// ルーターを設定
	r := chi.NewRouter()

	// エンドポイントを設定
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})
	
	r.Get("/top-photos", photoHandler.GetTopPhotos)

	// サーバーを起動
	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
} 