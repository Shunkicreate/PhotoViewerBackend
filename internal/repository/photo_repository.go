package repository

import (
    "fmt"
    "os"
	"math/rand"
    "photo_viewer_backend/internal/model"
    "strings"
)

type PhotoRepository interface {
    GetTopPhotos(count int) ([]model.Photo, error)
}

type photoRepository struct {
    // 後でDBの設定を追加する場合はここに追加
}

func NewPhotoRepository() PhotoRepository {
    return &photoRepository{}
}

func (r *photoRepository) GetTopPhotos(count int) ([]model.Photo, error) {
	// 環境変数からNASのパスを取得
	photoDir := os.Getenv("NAS_PATH")
	fmt.Println("アクセスするパス:", photoDir)

	// ディレクトリを開く
	dir, err := os.Open(photoDir)
	if err != nil {
		return nil, fmt.Errorf("写真ディレクトリへのアクセスに失敗: %v", err)
	}
	defer dir.Close()

	// ディレクトリ内のファイル一覧を取得
	files, err := dir.Readdir(-1)
	if err != nil {
		return nil, fmt.Errorf("ディレクトリの読み取りに失敗: %v", err)
	}

	// ランダムにファイルを選択
	if len(files) > count {
		rand.Shuffle(len(files), func(i, j int) {
			files[i], files[j] = files[j], files[i]
		})
		files = files[:count]
	}

	var photos []model.Photo
	for i, file := range files {
		if !isImageFile(file.Name()) {
			continue
		}

		photoPath := fmt.Sprintf("%s/%s", photoDir, file.Name())
		fileData, err := os.ReadFile(photoPath)
		if err != nil {
			return nil, fmt.Errorf("ファイルの読み取りに失敗: %v", err)
		}

		photo := model.Photo{
			ID:          fmt.Sprintf("%x%x", i+1, len(file.Name())*17),
			Title:       file.Name(),
			URL:         fmt.Sprintf("file://%s", photoPath),
			Description: "", // ファイルの説明は今後必要に応じて追加
			ImageData:   fileData,
		}
		photos = append(photos, photo)
	}

	return photos, nil
}

func isImageFile(filename string) bool {
    return strings.HasSuffix(strings.ToLower(filename), ".jpg") || strings.HasSuffix(strings.ToLower(filename), ".jpeg")
}