package repository

import (
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "photo_viewer_backend/internal/model"
    "sync"
    "time"
    "os"
)

type PhotoRepository interface {
    GetTopPhotos(count int, width int, height int) ([]model.ImageFile, error)
    GetPhoto(path string, width int, height int) (*http.Response, error)
}

type photoRepository struct {
    fileCache      map[string][]byte
    cacheMutex     sync.Mutex
    cacheTime      map[string]time.Time
    cacheTTL       time.Duration
    apiBaseURL     string
}

func NewPhotoRepository() PhotoRepository {
    return &photoRepository{
        fileCache:   make(map[string][]byte),
        cacheTime:   make(map[string]time.Time),
        cacheTTL:    10 * time.Minute,
        apiBaseURL:  fmt.Sprintf("http://%s:8090/api", os.Getenv("NAS_SERVER_PATH")),
    }
}

func (r *photoRepository) GetTopPhotos(count, width, height int) ([]model.ImageFile, error) {
    // APIエンドポイントを構築
    url := fmt.Sprintf("%s/files/random?count=%d&width=%d&height=%d", r.apiBaseURL, count, width, height)
    
    // HTTPリクエストを送信
    resp, err := http.Get(url)
    if err != nil {
        return nil, fmt.Errorf("APIリクエストに失敗: %v", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("APIから不正なレスポンス: %s", resp.Status)
    }

    // レスポンスボディを読み取り
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, fmt.Errorf("レスポンスの読み取りに失敗: %v", err)
    }

    // ImageFileの配列にデコード
    var imageFiles []model.ImageFile
    if err := json.Unmarshal(body, &imageFiles); err != nil {
        return nil, fmt.Errorf("JSONのデコードに失敗: %v", err)
    }

    // キャッシュの更新
    for _, imgFile := range imageFiles {
        r.cacheMutex.Lock()
        if !r.isCacheValid(imgFile.Path) {
            r.fileCache[imgFile.Path] = imgFile.Data
            r.cacheTime[imgFile.Path] = time.Now()
        }
        r.cacheMutex.Unlock()
    }

    return imageFiles, nil
}

func (r *photoRepository) GetPhoto(path string, width, height int) (*http.Response, error) {
    // APIエンドポイントを構築
    url := fmt.Sprintf("%s/files/image/nas/%s", r.apiBaseURL, path)
    fmt.Println(url)

    // HTTPリクエストを送信
    resp, err := http.Get(url)
    if err != nil {
        return nil, fmt.Errorf("APIリクエストに失敗: %v", err)
    }

    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("APIから不正なレスポンス: %s", resp.Status)
    }

    return resp, nil
}

func (r *photoRepository) isCacheValid(path string) bool {
    _, exists := r.fileCache[path]
    if !exists {
        return false
    }
    return time.Since(r.cacheTime[path]) < r.cacheTTL
}
