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
}
```

prismaクライアントの生成
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
