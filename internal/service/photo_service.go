package service

import (
    "photo_viewer_backend/internal/model"
    "photo_viewer_backend/internal/repository"
)

type PhotoService struct {
    repo repository.PhotoRepository
}

func NewPhotoService(repo repository.PhotoRepository) *PhotoService {
    return &PhotoService{repo: repo}
}

func (s *PhotoService) GetTopPhotos(count int) ([]model.Photo, error) {
    return s.repo.GetTopPhotos(count)
}
