package main

import (
	"context"
	"fmt"
	"log"
	"sample-ent-schema-edges/ent"
	"sample-ent-schema-edges/ent/group"
	"sample-ent-schema-edges/ent/user"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	O2OTwoTypes(context.Background(), client)

	O2MTwoTypes(context.Background(), client)

	M2MTwoWypes(context.Background(), client)
}

// O2O Two Types
func O2OTwoTypes(ctx context.Context, client *ent.Client) error {
	u, err := client.User.
		Create().
		SetAge(10).
		SetName("user1").
		Save(ctx)
	if err != nil {
		return fmt.Errorf("creating user: %w", err)
	}
	log.Println("user: ", u)

	ca1, err := client.Card.
		Create().
		SetOwner(u).
		SetNumber("0000").
		SetExpired(time.Now().Add(time.Minute)).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("creating card: %w", err)
	}
	log.Println("card: ", ca1)

	ca2, err := u.QueryCard().Only(ctx)
	if err != nil {
		return fmt.Errorf("querying card: %w", err)
	}
	log.Println("card: ", ca2)

	o, err := ca2.QueryOwner().Only(ctx)
	if err != nil {
		return fmt.Errorf("querying owner: %w", err)
	}
	log.Println("owner: ", o)

	return nil
}

// O2M Two Types​
func O2MTwoTypes(ctx context.Context, client *ent.Client) error {
	p1, err := client.Pet.
		Create().
		SetName("pet1").
		Save(ctx)
	if err != nil {
		return fmt.Errorf("creating pet: %w", err)
	}
	p2, err := client.Pet.
		Create().
		SetName("pet2").
		Save(ctx)
	if err != nil {
		return fmt.Errorf("creating pet: %w", err)
	}

	u, err := client.User.
		Create().
		SetAge(10).
		SetName("user").
		AddPets(p1, p2).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("creating user: %w", err)
	}
	fmt.Println("User created: ", u)

	o := p2.QueryOwner().OnlyX(ctx)
	fmt.Println(o.Name)

	count := p2.
		QueryOwner().
		QueryPets().
		CountX(ctx)
	fmt.Println(count)

	return nil
}

// M2M Two Types​
func M2MTwoWypes(ctx context.Context, client *ent.Client) error {
	g1 := client.Group.
		Create().
		SetName("group1").
		SaveX(ctx)
	g2 := client.Group.
		Create().
		SetName("group2").
		SaveX(ctx)
	u1 := client.User.
		Create().
		SetAge(10).
		SetName("user1").
		AddGroups(g1).
		SaveX(ctx)
	u2 := client.User.
		Create().
		SetAge(20).
		SetName("user2").
		AddGroups(g1, g2).
		SaveX(ctx)

	groups, err := u1.
		QueryGroups().
		All(ctx)
	if err != nil {
		return fmt.Errorf("querying user1 groups: %w", err)
	}
	fmt.Println(groups)

	groups, err = u2.
		QueryGroups().
		All(ctx)
	if err != nil {
		return fmt.Errorf("querying user1 groups: %w", err)
	}
	fmt.Println(groups)

	users, err := u1.
		QueryGroups().
		Where(
			group.Not(group.HasUsersWith(user.Name("user2"))),
		).
		QueryUsers().
		QueryGroups().
		QueryUsers().
		All(ctx)
	if err != nil {
		return fmt.Errorf("traversing the graph %w", err)
	}
	fmt.Println(users)

	return nil
}
