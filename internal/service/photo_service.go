package service

import (
	"photo_viewer_backend/internal/model"
	"photo_viewer_backend/internal/repository"
)

type PhotoService interface {
	GetTopPhotos() ([]model.Photo, error)
}

type photoService struct {
	repo repository.PhotoRepository
}

func NewPhotoService(repo repository.PhotoRepository) PhotoService {
	return &photoService{repo: repo}
}

func (s *photoService) GetTopPhotos() ([]model.Photo, error) {
	return s.repo.GetTopPhotos()
}
