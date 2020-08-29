# cleanarchitecture-fes
CleanArchitectureのサンプル

## 環境準備
必要なもの
- VSCodeでGoが動かせる環境
- docker-compose

## 参考: Winows10 環境セットアップ
(8/29/2020現在で実施手環境構築した手順です。）
- VisualStudioCodeのインストール

https://azure.microsoft.com/ja-jp/products/visual-studio-code/

VSCodeUserSetup-x64-1.48.2.exeをダウンロードしてインストール

- Docker, Docker composeの導入

https://docs.docker.com/docker-for-windows/install/

からDocker Desktop Installer.exeをダウンロードしてインストール
ターミナルにて確認
```
>docker --version
Docker version 19.03.12, build 48a66213fe
>docker-compose --version
docker-compose version 1.26.2, build eefe0d31
```
- Goのインストール

https://golang.org/dl/

go1.15.windows-amd64.msi　をダウンロード＆インストール
ターミナルにて確認
```
>go version
go version go1.15 windows/amd64
```

- gitのインストール
以下からGit-2.28.0-64-bit.exe　をダウンロード＆インストール

https://gitforwindows.org/
ターミナルにて確認
```
>git --version
git version 2.28.0.windows.1
```

## DBの初期化・マイグレーション
macの場合
```
export MYSQL_ARGS=cafes:cafes99@(localhost:3306)/cleanarchitecture_fes
docker-compose up -d
go run maintain/initdb.go
```
Windowsの場合
```
$env:MYSQL_ARGS="cafes:cafes99@(localhost:3306)/cleanarchitecture_fes"
docker-compose up -d
go run maintain/initdb.go
```

## 起動
(先にDBの初期化マイグレーションを実施してください)
macの場合
```
go run src/entrypoint/*.go
```
Windowsの場合
```
go run .\src\entrypoint\main.go .\src\entrypoint\graphql_resolver.go
```

## 動作確認
ブラウザにてPlaygroundが立ち上がります。

http://localhost:19001　にアクセスしplaygroundで動作確認します。

### イベントを登録するユースケースAPI
```
mutation saveFesEvent{
  SaveFesEvent(input:{
    title: "CleanArchitectureでGo!"
    speaker: "Kenichi Suzuki"
  })
}
```
### イベントを参照するユースケースAPI
```
query getFesEvents{
  GetFesEvent {
    fesEvents{
      id
      title
      speaker
    }
  }
}
```

- Databaseの内容はAdminorで確認するこも可能です。

以下にアクセスします

https://localhost:9088

以下を入力しログインできます。
```
データベース種類:MySQL
サーバー:rdb
ユーザ名:cafes
パスワード:cafes99
```

## 終了
Ctrl+Cでgoプロセスを停止

Docker(MySQLサーバーコンテナ、Adminorコンテナ）を停止
```
docker-compose down
```

### GraphQL codegen

```
go run github.com/99designs/gqlgen --config graphql_schema/gqlgen.yml
```

#### lint

```
golint ./...
```
(golintない場合以下コマンドでインストール）
```
go get golang.org/x/lint/golint 
```
