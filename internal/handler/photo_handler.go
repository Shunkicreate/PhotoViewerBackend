package handler

import (
    "encoding/json"
    "net/http"
    "photo_viewer_backend/internal/service"
    "strconv"
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
            http.Error(w, "無効なパラメータ", http.StatusBadRequest)
            return
        }
    }

    // Get the 'width' parameter from the query string
    widthStr := r.URL.Query().Get("width")
    width := 0 // デフォルト値
    if widthStr != "" {
        var err error
        width, err = strconv.Atoi(widthStr)
        if err != nil || width <= 0 {
            http.Error(w, "無効なパラメータ", http.StatusBadRequest)
            return
        }
    }

    // Get the 'height' parameter from the query string
    heightStr := r.URL.Query().Get("height")
    height := 0 // デフォルト値
    if heightStr != "" {
        var err error
        height, err = strconv.Atoi(heightStr)
        if err != nil || height <= 0 {
            http.Error(w, "無効なパラメータ", http.StatusBadRequest)
            return
        }
    }

    photos, err := h.service.GetTopPhotos(count, width, height)
    if err != nil {
        http.Error(w, "内部サーバーエラー", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(photos)
}
