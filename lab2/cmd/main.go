package main

import (
	"context"
	"fmt"
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
		printMenu()
		fmt.Print("\nSelect option: ")
		_, err := fmt.Scan(&option)
		if err != nil {
			return fmt.Errorf("invalid option was provided")
		}

		switch option {
		case 1:
			customers, err := handler.GetCustomers()
			if err != nil {
				return err
			}

			for _, mov := range customers {
				fmt.Println("Відвідувач id =", mov.Id)
				fmt.Println("\tІм'я:", mov.FirstName)
				fmt.Println("\tПрізвище:", mov.LastName)
			}

		case 2:
			movies, err := handler.GetMovies()
			if err != nil {
				return err
			}

			for _, mov := range movies {
				fmt.Println("Фільм id =", mov.Id)
				fmt.Println("\tНазва:", mov.Title)
				fmt.Println("\tОпис:", mov.Description)
				fmt.Println("\tТривалість:", mov.Duration)
			}

			fmt.Println("Кількість фільмів =", len(movies))

		case 3:
			halls, err := handler.GetHalls()
			if err != nil {
				return err
			}

			for _, h := range halls {
				fmt.Println("Кінозал id =", h.Id)
				fmt.Println("\tНазва:", h.Title)
				fmt.Println("\tОпис:", h.Description)
				fmt.Println("\tВмістимість:", h.Capacity)
			}

			fmt.Println("Кількість кінозалів =", len(halls))

		case 4:
			sessions, err := handler.GetSessions()
			if err != nil {
				return err
			}

			for _, h := range sessions {
				fmt.Println("Сеанс id =", h.Id)
				fmt.Println("\tФільм:", h.MovieId)
				fmt.Println("\tКінозал:", h.HallId)
				fmt.Println("\tПочинається о:", h.StartAt)
			}

			fmt.Println("Кількість сеансів =", len(sessions))
		default:
			break
		}
	}

	return nil
}

func printMenu() {
	fmt.Println("--------------------------------------")
	fmt.Println("* Menu:")
	fmt.Println("\t~ 1. Get all customers from database")
	fmt.Println("\t~ 2. Get all movies from database")
	fmt.Println("\t~ 3. Get all halls from database")
	fmt.Println("\t~ 4. Get all sessions from database")
}
