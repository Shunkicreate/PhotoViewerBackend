package repository

import "photo_viewer_backend/internal/model"

type PhotoRepository interface {
	GetTopPhotos() ([]model.Photo, error)
}

type photoRepository struct {
	// 後でDBの設定を追加する場合はここに追加
}

func NewPhotoRepository() PhotoRepository {
	return &photoRepository{}
}

func (r *photoRepository) GetTopPhotos() ([]model.Photo, error) {
	// とりあえずモックデータを返す
	photos := []model.Photo{
		{
			ID:          "1",
			Title:       "富士山の写真",
			URL:         "https://example.com/photo1.jpg",
			Description: "富士山の美しい風景",
		},
		{
			ID:          "2",
			Title:       "桜の写真",
			URL:         "https://example.com/photo2.jpg",
			Description: "満開の桜",
		},
	}
	return photos, nil
}
