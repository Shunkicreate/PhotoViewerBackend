package handler

import (
    "encoding/json"
    "net/http"
    "photo_viewer_backend/internal/service"
    "strconv"
    "io"
    "fmt"
    "log"
)

type PhotoHandler struct {
    service *service.PhotoService
}

func NewPhotoHandler(service *service.PhotoService) *PhotoHandler {
    return &PhotoHandler{service: service}
}

func (h *PhotoHandler) GetTopPhotos(w http.ResponseWriter, r *http.Request) {
    // Get the 'count' parameter from the query string
    countStr := r.URL.Query().Get("count")
    count := 6 // デフォルト値
    if countStr != "" {
        var err error
        count, err = strconv.Atoi(countStr)
        if err != nil || count <= 0 {
            log.Printf("Invalid 'count' parameter: %v", err)
            http.Error(w, "無効なパラメータ", http.StatusBadRequest)
            return
        }
    }

    // Get the 'width' parameter from the query string
    widthStr := r.URL.Query().Get("width")
    width := 640
    if widthStr != "" {
        var err error
        width, err = strconv.Atoi(widthStr)
        if err != nil || width <= 0 {
            log.Printf("Invalid 'width' parameter: %v", err)
            http.Error(w, "無効なパラメータ", http.StatusBadRequest)
            return
        }
    }

    // Get the 'height' parameter from the query string
    heightStr := r.URL.Query().Get("height")
    height := 420
    if heightStr != "" {
        var err error
        height, err = strconv.Atoi(heightStr)
        if err != nil || height <= 0 {
            log.Printf("Invalid 'height' parameter: %v", err)
            http.Error(w, "無効なパラメータ", http.StatusBadRequest)
            return
        }
    }

    photos, err := h.service.GetTopPhotos(count, width, height)
    if err != nil {
        log.Printf("Failed to get top photos: %v", err)
        http.Error(w, "内部サーバーエラー", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(photos)
}

func (h *PhotoHandler) GetPhoto(w http.ResponseWriter, r *http.Request) {
    path := r.URL.Query().Get("path")

    widthStr := r.URL.Query().Get("width")
    width := 0 // デフォルト値
    if widthStr != "" {
        var err error
        width, err = strconv.Atoi(widthStr)
        if err != nil || width <= 0 {
            log.Printf("Invalid 'width' parameter: %v", err)
            http.Error(w, "無効なパラメータ", http.StatusBadRequest)
            return
        }
    }

    heightStr := r.URL.Query().Get("height")
    height := 0 // デフォルト値
    if heightStr != "" {
        var err error
        height, err = strconv.Atoi(heightStr)
        if err != nil || height <= 0 {
            log.Printf("Invalid 'height' parameter: %v", err)
            http.Error(w, "無効なパラメータ", http.StatusBadRequest)
            return
        }
    }

    resp, err := h.service.GetPhoto(path, width, height)
    if err != nil {
        log.Printf("Failed to get photo: %v", err)
        fmt.Printf("写真の取得に失敗: %v\n", err)
        http.Error(w, "写真が見つかりません", http.StatusNotFound)
        return
    }
    defer resp.Body.Close()

    w.Header().Set("Content-Type", "image/jpeg")
    w.WriteHeader(resp.StatusCode)
    io.Copy(w, resp.Body)
}
