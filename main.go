package main

import (
	"context"
	"entgo/db"
	"entgo/ent"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	id, passwd := "root", "root"
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/service?parseTime=True", id, passwd)
	client, err := ent.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resource: %v", err)
	}

	db.CreateUser(context.TODO(), client)
}
