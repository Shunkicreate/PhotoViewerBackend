Dockerfileを使ってビルドして、デプロイする

## ローカルでビルドする
```
docker build -t photo_viewer_backend .
```

## ローカルで実行する
```
docker run -p 8080:8080 photo_viewer_backend
```
