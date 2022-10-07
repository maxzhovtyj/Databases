package main

import (
	"context"
	"lab2/internal/handler"
	"lab2/internal/repository"
	"lab2/internal/service"
	"lab2/pkg/client/postgresql"
	"log"
)

func main() {
	db, err := postgresql.NewPostgresClient(context.Background(), postgresql.StorageConfig{
		Username: "postgres",
		Password: "30042003",
		Host:     "localhost",
		Port:     5432,
		Database: "Lab1",
	})
	if err != nil {
		log.Fatalf("failed to connect to postgres database, %v", err)
	}

	repo := repository.NewRepository(db)
	serviceInstance := service.NewService(repo)
	handlerInstance := handler.NewHandler(serviceInstance)

	err = runServer(handlerInstance)
	if err != nil {
		log.Fatalln(err)
	}
}

func runServer(handler handler.Handler) error {
	var option int

	for true {
		switch option {
		case 1:
			err := handler.Search()
			if err != nil {
				return err
			}
		}
	}

	return nil
}
