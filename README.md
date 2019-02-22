# docker_webapp
JSON API on Docker

## Description
Docker上でGo言語を用いてJSON APIを作成しました。  
Postgresと連結して名前とメールアドレスを管理します。  
**20180702: [Review](/Review.txt)を元に改善を行いました**  
**20190223: [Docker multi-stage builds](https://docs.docker.com/develop/develop-images/multistage-build/)に変更しました。ローカルのgolang環境が必要なくなりました**

ディレクトリ構成
<pre>
.
├── README.md
├── docker-compose.yml   # docker-compose設定ファイル
└── webapp               # WebAPIコンテナ
    ├── CRUD             # PostgresのCRUD操作用Go言語パッケージ
    │   └── CRUD.go
    ├── Dockerfile
    ├── Makefile
    └── webapp.go
</pre>

## Requirement
<pre>
#依存ソフトウェア
docker  
docker-compose

#依存コンテナ
alpine:latest
golang:alpine
postgres:alpine
</pre>

## Usage
サーバー側
<pre>
#ビルド
$ docker-compose build

#コンテナ起動
$ docker-compose up -d

#コンテナ終了
$ docker-compose down
</pre>

#### Hello World!!
httpリクエスト
<pre>
$ curl -XGET -H 'Content-Type:application/json' http://localhost:8080/
</pre>

レスポンス(HTTP ステータスコード 200)
<pre>
{
    "message": "Hello World!!"
}
</pre>

#### Create
httpリクエスト
<pre>
$ curl -XPOST -H 'Content-Type:application/json' http://localhost:8080/users -d '{"name": "test", "email": "hoge@example.com" }'
</pre>

レスポンス(HTTP ステータスコード 200)
<pre>
{
    "id": 1,
    "name": "test",
    "email": "hoge@example.com",
    "created_at": "2018-05-27T23:30:46.061325+09:00",
    "updated_at": "2018-05-27T23:30:46.061325+09:00"
}
</pre>

#### Update
httpリクエスト
<pre>
$ curl -XPUT -H 'Content-Type:application/json' http://localhost:8080/users/1 -d '{"name": "koudaiii", "email": "hoge@example.com" }'
</pre>

レスポンス(HTTP ステータスコード 200)
<pre>
{
    "id": 1,
    "name": "koudaiii",
    "email": "hoge@example.com",
    "created_at": "2018-05-27T23:30:46.061325+09:00",
    "updated_at": "2018-05-27T23:31:14.140414+09:00"
}
</pre>


#### Read
httpリクエスト
<pre>
$ curl -XGET -H 'Content-Type:application/json' http://localhost:8080/users/1
</pre>

レスポンス(HTTP ステータスコード 200)
<pre>
{
    "id": 1,
    "name": "koudaiii",
    "email": "hoge@example.com",
    "created_at": "2018-05-27T23:30:46.061325+09:00",
    "updated_at": "2018-05-27T23:31:14.140414+09:00"
}
</pre>


#### Read (All)
httpリクエスト
<pre>
$ curl -XGET -H 'Content-Type:application/json' http://localhost:8080/users
</pre>

レスポンス(HTTP ステータスコード 200)
<pre>
[
    {
        "id": 1,
        "name": "koudaiii",
        "email": "hoge@example.com",
        "created_at": "2018-05-27T23:30:46.061325+09:00",
        "updated_at": "2018-05-27T23:31:14.140414+09:00"
    }
]
</pre>


#### Delete
httpリクエスト
<pre>
$ curl -XDELETE -H 'Content-Type:application/json' http://localhost:8080/users/1
</pre>

レスポンス(HTTP ステータスコード 200)
<pre>

</pre>
