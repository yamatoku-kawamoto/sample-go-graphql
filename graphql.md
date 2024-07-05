# Readme


## Go GraphQL


Github
https://github.com/99designs/gqlgen


公式リファレンス
https://gqlgen.com/getting-started/

## 概要

GraphQLの細かい処理を自分で作るのは困難。
また、多くの処理は共通である。

開発者が必要なことに注力できるように、gqlgenというライブラリを使用し、開発利便性を高めることにする。

開発者は、schemaのモデリングやデータの抽出、作成などの具体的な処理に集中することができる。

## 構造

### bin

goはコンパイル言語なため、OSが実行可能なバイナリを出力する。
binディレクトリは実行可能なバイナリを格納する。

### cmd

goのエントリーポイントとなるソースファイルを格納する。
goはディレクトリの命名規則がなく、任意に設定可能であるが、慣習的にcmdとする。

### graph

GraphQLの実装をここに書く。
多くはライブラリによって自動生成されるソースファイルであり、開発者が気にするファイルは、 **schema.graphqls** 、**schema.resolvers.go** の2つである。

なお、**values.go** は、サンプルのために作成したソースファイルであり、GraphQLの実装には直接関係しない。

### server

GraphQLをAPIとして提供するためのソースファイルを格納する。
GraphQLの実装には直接関係しない。

### store

データベース関係の処理を格納する。
ORMとして利用する形とするが、GraphQLの実装とは直接関係しない。
実際にGraphQLを開発する場合、データベースとの関係が発生するため、実装することとなる。

## Getting Started

goをインストール

バージョンは **1.22.1** 以上を推奨

makeがインストールされている場合
```
$ make go/run
```

makeがインストールされていない場合
```
go run cmd/server.go
```

ブラウザで http://localhost:9980/ にアクセス

クエリを実行してみる

### シンプルなクエリを試してみる

```
query {
  todos {
    id
  	text
	}
}
```
このような出力が得られる
```
{
  "data": {
    "todos": [
      {
        "id": "0",
        "text": "mail check"
      },
      {
        "id": "1",
        "text": "MTG"
      }
    ]
  }
}
```

### 変数を使ったクエリを試してみる

```
query ($todos: [getTodo]!){
  getTodos(input: $todos) {
    id
    text
  }
}
```
```Variables
{
  "todos": [{"id":"0"},{"id":"1"}]
}
```
出力はさっきと同じだろう
```
{
  "data": {
    "todos": [
      {
        "id": "0",
        "text": "mail check"
      },
      {
        "id": "1",
        "text": "MTG"
      }
    ]
  }
}
```

試しに、変数のtodos配列にあるオブジェクトを一つ消す。
思った通りに動くだろう。

