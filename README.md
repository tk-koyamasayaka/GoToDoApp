# Go Todo API

GinとGORMを使用したシンプルなTodo APIです。クリーンアーキテクチャの原則に従い、保守性と拡張性を考慮して設計されています。

## 機能

- Todo の作成
- Todo 一覧の取得
- 特定の Todo の取得
- Todo の更新
- Todo の削除

## 技術スタック

- [Gin Web Framework](https://github.com/gin-gonic/gin)
- [GORM](https://gorm.io/)
- SQLite3

## プロジェクト構造

```
.
├── controllers/     # HTTPリクエストの処理
├── models/         # データモデル
├── repositories/   # データアクセス層
├── dto/           # データ転送オブジェクト
├── validators/    # バリデーションロジック
├── responses/     # レスポンスフォーマット
├── mappers/       # DTOとモデル間の変換
├── test/          # テストコード
│   ├── controllers/  # コントローラーのテスト
│   └── mocks/        # モックオブジェクト
└── main.go        # アプリケーションのエントリーポイント
```

## セットアップ

1. リポジトリのクローン:
```bash
git clone https://github.com/yourusername/GoToDoApp.git
cd GoToDoApp
```

2. 依存関係のインストール:
```bash
go mod tidy
```

3. アプリケーションの起動:
```bash
go run main.go
```

サーバーは `http://localhost:8080` で起動します。

## API エンドポイント

### Todo の作成
```bash
POST /api/v1/todos

リクエストボディ:
{
    "title": "買い物",
    "description": "牛乳を買う",
    "completed": false
}
```

### Todo 一覧の取得
```bash
GET /api/v1/todos
```

### 特定の Todo の取得
```bash
GET /api/v1/todos/:id
```

### Todo の更新
```bash
PUT /api/v1/todos/:id

リクエストボディ:
{
    "title": "買い物",
    "description": "牛乳と卵を買う",
    "completed": true
}
```

### Todo の削除
```bash
DELETE /api/v1/todos/:id
```

## レスポンス形式

すべてのAPIレスポンスは以下の形式で返されます：

```json
{
    "success": true,
    "message": "処理が成功しました",
    "data": {
        // レスポンスデータ
    },
    "errors": null
}
```

エラー時:
```json
{
    "success": false,
    "message": "エラーが発生しました",
    "data": null,
    "errors": {
        // エラー詳細
    }
}
```

## テスト

テストの実行:
```bash
go test -v ./test/...
```