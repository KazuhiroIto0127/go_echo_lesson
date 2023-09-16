# go_echo_lesson

goのechoフレームワークを使ってwebサーバーを作れます。
dbはsqlite3を使います。

## 開発環境（build用）

```
docker compose up -d dev
```

devコンテナ内では下記を実行します。

ホットリロード対応
```
$ air
```

sqlite3クライアントがあります。
```
$ sqlite3 example.sql

> .tables
> .schema user
> .exit
```

## 本番環境

```
docker compose up -d prod
```

参考にさせていただきました。

- https://zenn.dev/def_yuisato/articles/echo-get-started
- https://iketechblog.com/database-sql-go-sqlite3/
