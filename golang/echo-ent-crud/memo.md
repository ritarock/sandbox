---
title: ent について調べた
date: 20220108155408
tags: ['golang']
---

### インストール
```bash
$ go get -d entgo.io/ent/cmd/ent
```

### スキーマの作成
`User` を作成する。
```bash
$ go run entgo.io/ent/cmd/ent init User
```

### コード生成
スキーマの内容からコードを生成する。
```bash
$ go generate ./ent
```

### フィールドの作成
`ent/schema/user.go` を編集する。編集後は `go generate ./ent` の実行が必要。
```go
// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("age").
			Positive(),
		field.String("name").
			Default("unknown"),
	}
}
```

### エンティティの作成
```go
func createUser(ctx context.Context, client *ent.Client, name string, age int) (*ent.User, error) {
	u, err := client.User.
		Create().
		SetName(name).
		SetAge(age).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	log.Println("user was created: ", u)

	return u, nil
}
```

### エンティティの問い合わせ
`Only` を使った場合、エンティティがなかったらエラーを返す。
```go
func queryUser(ctx context.Context, client *ent.Client, name string) (*ent.User, error) {
	u, err := client.User.
		Query().
		Where(user.Name(name)).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}
	log.Println("user returned: ", u)

	return u, nil
}
```

### エッジを追加する
`Car / Group` スキーマを作成する。
```bash
$ go run entgo.io/ent/cmd/ent init Car Group
```

User は複数の Car を持ち、 Car は一人の User しか持てない事 ( 1 対多 ) を定義する。

`ent/schema/user.go` を編集する。編集後は `go generate ./ent` の実行が必要。

`cars` エッジを User スキーマに追加する。
```go
// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("cars", Car.Type),
	}
}
```

1 台の car を作成して user に関連付ける。
```go
// 呼び出し側 ~~
	if err := createCars(ctx, client, u1, "tesla"); err != nil {
		log.Fatalf("failed creating car: %v", err)
	}
	if err := createCars(ctx, client, u1, "ford"); err != nil {
		log.Fatalf("failed creating car: %v", err)
	}

  // u1 に関連付けられた cars の一覧を取得する
	u1HaveCars, _ := u1.QueryCars().All(ctx)
	fmt.Println(u1HaveCars) // [Car(id=1, model=tesla, registered_at=Sat Jan  8 17:55:21 2022) Car(id=2, model=ford, registered_at=Sat Jan  8 17:55:21 2022)]
// ~~

func createCars(ctx context.Context, client *ent.Client, user *ent.User, carModelName string) error {
	c, err := client.Car.
		Create().
		SetModel(carModelName).
		SetRegisteredAt(time.Now()).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("failed creating car: %w", err)
	}
	log.Println("car was created: ", c.Model)

	if err := user.Update().AddCars(c).Exec(ctx); err != nil {
		return fmt.Errorf("failed add cars: %w", err)
	}
	return nil
}
```

### 逆方向のエッジ ( BackRef ) を追加する
`ent/schema/user.go` を編集する。編集後は `go generate ./ent` の実行が必要。

`edge.From` を使って逆方向エッジを定義する。
```go
// Edges of the Car.
func (Car) Edges() []ent.Edge {
	return []ent.Edge{
		// User 型の owner という逆エッジを作成
		edge.From("owner", User.Type).
			// Ref を使って明示的に User スキーマの cars エッジを参照する
			Ref("cars").
			// cars は 一人のオーナーのみが所有することを保証する
			Unique(),
	}
}
```

`cars` から Owner を紹介できる。
```go
	u1HaveCars, _ := u1.QueryCars().All(ctx)
	for _, ca := range u1HaveCars {
		owner, _ := ca.QueryOwner().Only(ctx)
		log.Printf("car %q owner: %q\n", ca.Model, owner.Name)
	}
```

### M2M ( 多対多 ) のリレーションを追加する
User と Group に M2M のリレーションを追加する。

Group は users を持つ。
```go
// Edges of the Group.
func (Group) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("users", User.Type),
	}
}
```
User は groups の逆エッジを持ち、それは users を参照する。
```go
// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("cars", Car.Type),
		edge.From("groups", Group.Type).
			Ref("users"),
	}
}
```

### グラフ探索
グラフを作成する。
```go
func createGraph(ctx context.Context, client *ent.Client) error {
	// user を複数作成する
	u1, err := client.User.
		Create().
		SetAge(10).
		SetName("user1").
		Save(ctx)
	if err != nil {
		return err
	}
	u2, err := client.User.
		Create().
		SetAge(10).
		SetName("user2").
		Save(ctx)
	if err != nil {
		return err
	}

	// car を複数作成し、ユーザと関連付ける
	err = client.Car.
		Create().
		SetModel("Tesla").
		SetRegisteredAt(time.Now()).
		SetOwner(u1).
		Exec(ctx)
	if err != nil {
		return err
	}
	err = client.Car.
		Create().
		SetModel("Mazda").
		SetRegisteredAt(time.Now()).
		SetOwner(u1).
		Exec(ctx)
	if err != nil {
		return err
	}
	err = client.Car.
		Create().
		SetModel("Mazda").
		SetRegisteredAt(time.Now()).
		SetOwner(u2).
		Exec(ctx)
	if err != nil {
		return err
	}

	// Group を作成し、ユーザーと関連付ける
	err = client.Group.
		Create().
		SetName("GitLab").
		AddUsers(u1, u2).
		Exec(ctx)
	if err != nil {
		return err
	}
	err = client.Group.
		Create().
		SetName("GitHub").
		AddUsers(u1).
		Exec(ctx)
	if err != nil {
		return err
	}
	fmt.Println("Success")
	return nil
}


// 呼び出し側 ~~
	createGraph(ctx, client)
	carOfGithub, _ := client.Group.
		Query().
		Where(group.Name("GitHub")).
		QueryUsers().
		QueryCars().
		All(ctx)
	fmt.Println(carOfGithub)
	// [Car(id=3, model=Tesla, registered_at=Sat Jan  8 18:22:44 2022) Car(id=4, model=Mazda, registered_at=Sat Jan  8 18:22:44 2022)]
// ~~
```

