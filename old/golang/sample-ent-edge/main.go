package main

import (
	"context"
	"fmt"
	"log"
	"sample-ent-edge/ent"
	"sample-ent-edge/ent/group"
	"sample-ent-edge/ent/user"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	ctx := context.Background()

	// オートマイグレーションツールを実行する
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	u1, _ := createUser(ctx, client, "user1", 10)
	fmt.Println(u1)

	u11, _ := queryUser(ctx, client, "user1")
	fmt.Println(u11)

	if err := createCars(ctx, client, u1, "tesla"); err != nil {
		log.Fatalf("failed creating car: %v", err)
	}
	if err := createCars(ctx, client, u1, "ford"); err != nil {
		log.Fatalf("failed creating car: %v", err)
	}

	u1HaveCars, _ := u1.QueryCars().All(ctx)
	fmt.Println(u1HaveCars)
	// [Car(id=1, model=tesla, registered_at=Sat Jan  8 17:55:21 2022) Car(id=2, model=ford, registered_at=Sat Jan  8 17:55:21 2022)]

	for _, ca := range u1HaveCars {
		owner, _ := ca.QueryOwner().Only(ctx)
		log.Printf("car %q owner: %q\n", ca.Model, owner.Name)
	}

	createGraph(ctx, client)
	carOfGithub, _ := client.Group.
		Query().
		Where(group.Name("GitHub")).
		QueryUsers().
		QueryCars().
		All(ctx)
	fmt.Println(carOfGithub)
	// [Car(id=3, model=Tesla, registered_at=Sat Jan  8 18:22:44 2022) Car(id=4, model=Mazda, registered_at=Sat Jan  8 18:22:44 2022)]
}

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

func queryUser(ctx context.Context, client *ent.Client, name string) (*ent.User, error) {
	u, err := client.User.
		Query().
		Where(user.Name(name)).
		// ユーザが見つからない場合、 Only は error を返す
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}
	log.Println("user returned: ", u)

	return u, nil
}

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
