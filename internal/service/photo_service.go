package service

import (
    "photo_viewer_backend/internal/model"
    "photo_viewer_backend/internal/repository"
    "net/http"
)

type PhotoService struct {
    repo repository.PhotoRepository
}

func NewPhotoService(repo repository.PhotoRepository) *PhotoService {
    return &PhotoService{repo: repo}
}

func (s *PhotoService) GetTopPhotos(count, width, height int) ([]model.ImageFile, error) {
    return s.repo.GetTopPhotos(count, width, height)
}

func (s *PhotoService) GetPhoto(path string, width, height int) (*http.Response, error) {
    return s.repo.GetPhoto(path, width, height)
}
