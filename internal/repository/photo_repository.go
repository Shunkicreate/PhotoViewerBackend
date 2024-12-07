package repository

import (
    "fmt"
    "math/rand"
    "os"
    "photo_viewer_backend/internal/model"
    "strings"
    "sync"
    "time"
)

type PhotoRepository interface {
    GetTopPhotos(count int) ([]model.Photo, error)
}

type photoRepository struct {
    fileCache      map[string][]byte
    cacheMutex     sync.Mutex
    cacheTime      map[string]time.Time
    cacheTTL       time.Duration
}

func NewPhotoRepository() PhotoRepository {
    return &photoRepository{
        fileCache:  make(map[string][]byte),
        cacheTime:  make(map[string]time.Time),
        cacheTTL:   10 * time.Minute, // キャッシュの有効期限を10分に設定
    }
}

func (r *photoRepository) GetTopPhotos(count int) ([]model.Photo, error) {
    // 環境変数からNASのパスを取得
    photoDir := os.Getenv("NAS_PATH")

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

    // 画像ファイルのみをフィルタリング
    var imageFiles []os.FileInfo
    for _, file := range files {
        if isImageFile(file.Name()) {
            imageFiles = append(imageFiles, file)
        }
    }

    // ランダムにファイルを選択
    if len(imageFiles) > count {
        rand.Shuffle(len(imageFiles), func(i, j int) {
            imageFiles[i], imageFiles[j] = imageFiles[j], imageFiles[i]
        })
        imageFiles = imageFiles[:count]
    }

    var photos []model.Photo
    for i, file := range imageFiles {
        photoPath := fmt.Sprintf("%s/%s", photoDir, file.Name())

        // キャッシュを確認
        r.cacheMutex.Lock()
        fileData, cached := r.fileCache[photoPath]
        if cached && time.Since(r.cacheTime[photoPath]) < r.cacheTTL {
            r.cacheMutex.Unlock()
        } else {
            // キャッシュがないか期限切れの場合、ファイルを読み込む
            fileData, err = os.ReadFile(photoPath)
            if err != nil {
                r.cacheMutex.Unlock()
                return nil, fmt.Errorf("ファイルの読み取りに失敗: %v", err)
            }
            // キャッシュを更新
            r.fileCache[photoPath] = fileData
            r.cacheTime[photoPath] = time.Now()
            r.cacheMutex.Unlock()
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
    return strings.HasSuffix(strings.ToLower(filename), ".jpg") || strings.HasSuffix(strings.ToLower(filename), ".jpeg") || strings.HasSuffix(strings.ToLower(filename), ".png")
}
