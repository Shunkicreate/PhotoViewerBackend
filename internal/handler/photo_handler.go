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
    count, err := strconv.Atoi(countStr)
    if err != nil || count <= 0 {
        http.Error(w, "無効なパラメータ", http.StatusBadRequest)
        return
    }

    photos, err := h.service.GetTopPhotos(count)
    if err != nil {
        http.Error(w, "内部サーバーエラー", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(photos)
}
