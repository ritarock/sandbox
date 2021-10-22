まずは project の作成.

```bash
$ go mod init prisma-todo
```

clientのインストール
```bash
$ go get github.com/prisma/prisma-client-go
```

スキーマの作成
```bash
npx prisma init
```

生成されたファイルを編集
```.env
# Environment variables declared in this file are automatically made available to Prisma.
# See the documentation for more detail: https://pris.ly/d/prisma-schema#using-environment-variables

# Prisma supports the native connection string format for PostgreSQL, MySQL, SQLite, SQL Server and MongoDB (Preview).
# See the documentation for all the connection string options: https://pris.ly/d/connection-strings

DATABASE_URL="mysql://user:password@localhost:3306/app"
```

```schema.prisma
// This is your Prisma schema file,
// learn more about it in the docs: https://pris.ly/d/prisma-schema

datasource db {
  provider = "mysql"
  url      = env("DATABASE_URL")
}

generator client {
  provider = "go run github.com/prisma/prisma-client-go"
}

model Post {
  id String @default(cuid()) @id
  createdAt DateTime @default(now())
  updatedAt DateTime @updatedAt
  title String
  published Boolean
  desc String?
  views Int @default(0)
}
```

prismaクライアントの生成
prisma の下に色々生成される
```bash
go run github.com/prisma/prisma-client-go generate
```

DBにスキーマの反映
```bash
go run github.com/prisma/prisma-client-go migrate dev --name init
```

modelに関連付けを行う
```scema.prisma
model Post {


    comments Comment[]
}

model Comment {
  id String @id @default(cuid())
  createdAt DateTime @default(now())
  updatedAt DateTime @updatedAt
  content String

  post Post @relation(fields: [postID], references: [id])
  postID String
}
```

migration を実行
```bash
go run github.com/prisma/prisma-client-go migrate dev --name add_comment_model
```

Prisma Studio を使う
`localhost:5555` で起動する.
```bash
$ npx prisma studio
```

# prisma api syntax
全件取得したいとき
```go
posts, err := client.Post.FindMany().Exec(ctx)
```

クエリで取得するとき
```go
posts, err := client.Post.FindMany(
  db.Post.Title.Equals("hello")
  // <model>.<field>.<method>.(value) 基本的にこの形式で使う
)
```

一意なデータを取得するとき
`schema.prisma` で `@id`,`@unique` でマークされたもののみ使用可能.
```go
post, err := client.Post.FindUnique(
  db.Post.ID.Equals("123")
).Exec(ctx)
if errors.Is(err, db.ErrNotFound) {
  log.Printf("no record with id 123")
} else if err != nil {
  log.Printf("error occurred: %s", err)
}
```

最初に見つかった1件を取得する
```go
post, err := client.Post.FindFirst(
  db.Post.Title.Equals("hi")
).Exec(ctx)
if errors.Is(err, db.ErrNotFound) {
  log.Printf("no record with title 'hi' found")
} else if err != nil {
  log.Printf("error occurred: %s", err)
}
```

Query API
スキーマのデータ型によってよしなにやってくれる
```go
posts, err := client.Post.FindMany(
  db.Post.Title.Contains("What")
).Exec(ctx)
```

string filters
```go
db.Post.Title.Equals("my post")
db.Post.Title.Contains("post")
db.Post.Title.StartsWith("my post")
db.Post.Title.EndsWith("my post")
```

number filters
```go
// views が 50 である post を取得
db.Post.Views.Equals(50)
// views が 50 以下の post を取得
db.Post.Views.Lte(50)
// views が 50 未満の post を取得
db.Post.Views.Lt(50)
// views が 50 以上の post を取得
db.Post.Views.Gte(50)
// views が 50 より大きいの post を取得
db.Post.Views.Gte(50)
```

time filter
```go
// 昨日作成された post を取得する
db.Post.CreatedAt.Equals(yesterday)
// 過去 6 時間で作られた post を取得する (createdAt > 6 hours ago)
db.Post.Gt(time.Now().Add(-6 * time.Hour))
// 過去 6 時間で作られた post を取得する (createdAt >= 6 hours ago)
db.Post.Gte(time.Now().Add(-6 * time.Hour))
// 昨日作成された Post を取得する
db.Post.Lt(time.Now().Truncate(24 * time.Hour))
// 昨日作成された Post を取得する (本日 00:00:00 を含む)
db.Post.Lte(time.Now().Truncate(24 * time.Hour))
```

NULL 関連
```go
db.Post.Content.EqualsOptional(nil)

content := "string"
db.Post.Content.EqualsOptional(&content)
```

他にも
Not
```go
db.Post.Not(
  db.Post.Title.Equals("123")
)
```

Or
```go
db.Post.Or(
  db.Post.Title.Equals("123"),
  db.Post.Content.Equals("456")
)
```


関連付けされたクエリ
```go
posts, err := client.Post.FindMany(
  db.Post.Title.Equals("What up?")
  db.Post.Comments.Some(
    db.Comment.Content.Equals("My Content"),
  ),
).Exec(ctx)
```
