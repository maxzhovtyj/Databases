package main

import (
	"context"
	"fmt"
	"lab2/pkg/client/postgresql"
	"log"
)

func main() {
	_, err := postgresql.NewPostgresClient(context.Background(), postgresql.StorageConfig{
		Username: "postgres",
		Password: "30042003",
		Host:     "localhost",
		Port:     5432,
		Database: "Lab1",
	})
	if err != nil {
		log.Fatalf("failed to connect to postgres database, %v", err)
	}

	fmt.Println("successfully connected to db")
}
