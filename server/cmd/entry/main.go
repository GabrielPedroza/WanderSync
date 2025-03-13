package main

import (
	"GabrielPedroza/WanderSync/ent"
	"GabrielPedroza/WanderSync/ent/migrate"
	"context"
	"log"
	"net/http"

	entry "GabrielPedroza/WanderSync"

	"entgo.io/ent/dialect"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Create ent.Client and run the schema migration.
	client, err := ent.Open(dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatal("opening ent client", err)
	}
	if err := client.Schema.Create(
		context.Background(),
		migrate.WithGlobalUniqueID(true),
	); err != nil {
		log.Fatal("opening ent client", err)
	}

	if err := addUser(client); err != nil {
		log.Fatal("failed adding user:", err)
	}
	// Configure the server and start listening on :8081.
	srv := handler.NewDefaultServer(entry.NewSchema(client))
	http.Handle("/",
		playground.Handler("WanderSync", "/query"),
	)
	http.Handle("/query", srv)
	log.Println("listening on :8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal("http server terminated", err)
	}
}

// addUser adds a new user to the database.
func addUser(client *ent.Client) error {
	// Create a new user
	u, err := client.User.Create().
		SetName("Alice").
		SetAge(30).
		Save(context.Background())
	if err != nil {
		return err
	}
	log.Printf("User created: %v\n", u)
	return nil
}
