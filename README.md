# TimeUS-api

reyくんと作り始めたAPIのリボジトリですわあ

## 使用したもの
- [Golang](https://golang.org/)
- [Echo](https://echo.labstack.com/) (Web Application Framework)
- [gorm](http://gorm.io/) (ORM)
- [JWT](https://jwt.io/)

## ユーザー認証API

### ユーザーの登録

`POST /signup`

### ユーザーの認証

`POST /login`


## post内操作api

JWTをリクエストに含めないとアクセスできないようにする(認証済である必要があるようにする)

### 新たなTodoの登録

`POST /api/posts`

### 指定されたIDのpost削除

`DELETE /api/posts/:id`

### 指定されたIDのpost状態の変更

`PUT /api/posts/:id/config`

## ディレクトリ構成

```
TimeUS/
├── db
│   └── ぱあ
├── handler
│   ├── auth.go
│   └── handler.go
├── model
│   ├── db.go
│   ├── post.go
│   └── user.go
├── route
│   └── router.go
├── .gitignore
├── Dockerfile
├── docker-compose.yaml
├── README.md
├── go.mod
├── go.sum
└── main.go
```
