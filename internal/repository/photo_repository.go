package repository

import (
	// "bufio"
	"fmt"
	"os"
	// "path/filepath"
	"photo_viewer_backend/internal/model"
	"strings"
)

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
	// 環境変数からNASのパスを取得
	// /mnt/photosにアクセス
	photoDir := "/mnt/photos/MyFavoritePhotos"
	fmt.Println("アクセスするパス:", photoDir)

	// ディレクトリを開く
	dir, err := os.Open(photoDir)
	fmt.Println(dir)
	if err != nil {
		return nil, fmt.Errorf("写真ディレクトリへのアクセスに失敗: %v", err)
	}
	defer dir.Close()

	// ディレクトリ内のファイル一覧を取得
	files, err := dir.Readdir(-1)
	fmt.Println(files)
	if err != nil {
		return nil, fmt.Errorf("ディレクトリの読み取りに失敗: %v", err)
	}

	fmt.Printf("見つかったファイル数: %d\n", len(files))
	for _, file := range files {
		fmt.Printf("ファイル名: %s\n", file.Name())
	}

	// var fileList []string
	// for fileScanner.Scan() {
	// 	fileList = append(fileList, fileScanner.Text())
	// }
	// fmt.Println(files)
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to read directory: %v", err)
	// }

	// var photos []model.Photo
	// for i, file := range fileList {
	// 	// 画像ファイルのみを処理
	// 	if !isImageFile(file) {
	// 		continue
	// 	}


	// 	// ファイルパスを構築
	// 	filePath := filepath.Join(nasPath, file)
	// 	fmt.Println(filePath)
	// 	photo := model.Photo{
	// 		ID:          fmt.Sprintf("%x%x", i+1, len(file)*17),
	// 		Title:       file,
	// 		URL:         fmt.Sprintf("file://%s", filePath),
	// 		Description: "", // ファイルの説明は今後必要に応じて追加
	// 	}
	// 	fmt.Println(photo)
	// 	photos = append(photos, photo)
	// }

	return nil, nil
}

func isImageFile(filename string) bool {
	return strings.HasSuffix(strings.ToLower(filename), ".jpg") || strings.HasSuffix(strings.ToLower(filename), ".jpeg")
}
