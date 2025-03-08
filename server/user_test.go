package user

import (
	"GabrielPedroza/WanderSync/ent"
	"context"
	"fmt"
	"log"

	"entgo.io/ent/dialect"
	_ "github.com/mattn/go-sqlite3"
)

func User_Test() {
	// Create an ent.Client with in-memory SQLite database.
	client, err := ent.Open(dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	ctx := context.Background()
	// Run the automatic migration tool to create all schema resources.
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	task1, err := client.User.Create().SetName("Gabriel").SetAge(23).Save(ctx)
	if err != nil {
		log.Fatalf("failed creating a user: %v", err)
	}

	fmt.Println(task1)
	// Output:
	// User(id=1, name=Gabriel, age=23)
}
