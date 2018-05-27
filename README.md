# docker_webapp
JSON API on Docker

## Description
Docker上でGo言語を用いてJSON APIを作成しました。  
JSON形式で"Hello World!!"を返します。

ディレクトリ構成
<pre>
HeWd/
├─ README.md
├─ docker-compose.yml
└─ webapp/
    ├─ Dockerfile
    └─ webapp.go
</pre>

## Requirement
<pre>
#依存ソフトウェア
docker  
docker-compose

＃依存コンテナ
golang:latest
</pre>

## Usage
サーバー側
<pre>
#コンテナ起動
$ docker-compose up -d

#コンテナ終了
$ docker-compose down
</pre>

クライアント側
<pre>
#httpリクエスト
$ curl -XGET -H 'Content-Type:application/json' http://localhost:8080/
</pre>

レスポンス(HTTP ステータスコード 200)
<pre>
{
    "message": "Hello World!!"
}
</pre>
