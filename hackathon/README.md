# leaQ

## 開発環境の立ち上げ

```sh
task up # またはdocker-compose up -d

# バックエンドの起動（backendディレクトリ内）
go run main.go

# フロントエンドの起動（frontendディレクトリ内）
npm install # 初回のみ
npm run dev

# 終了
task down # またはdocker-compose down
```

## 使っているポート番号

- 3306: MariaDB
- [5173](http://localhost:5173): フロントエンド
- 8080: バックエンド
- 9000: MinIO（画像の保存用）
- [9001](http://localhost:9001): MinIOの管理画面（ユーザー名、パスワード共に`minioadmin`）
