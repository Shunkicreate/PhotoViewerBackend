package main

import (
    "fmt"
    "log"
    "net/http"
    "os"

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
    port := "8080"
    if envPort := os.Getenv("BACKEND_PORT"); envPort != "" {
        port = envPort
    }

    fmt.Printf("Starting server on :%s\n", port)
    if err := http.ListenAndServe(":"+port, r); err != nil {
        log.Fatal(err)
    }
}