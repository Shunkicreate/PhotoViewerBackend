package handler

import (
	"encoding/json"
	"net/http"
	"photo_viewer_backend/internal/service"
)

type PhotoHandler struct {
	service service.PhotoService
}

func NewPhotoHandler(service service.PhotoService) *PhotoHandler {
	return &PhotoHandler{service: service}
}

func (h *PhotoHandler) GetTopPhotos(w http.ResponseWriter, r *http.Request) {
	photos, err := h.service.GetTopPhotos()
	if err != nil {
		http.Error(w, "内部サーバーエラー", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(photos)
}
