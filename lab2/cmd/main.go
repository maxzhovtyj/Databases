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
			err = handler.GetCustomers()
			if err != nil {
				return err
			}

		case 2:
			err = handler.GetMovies()
			if err != nil {
				return err
			}

		case 3:
			err = handler.GetHalls()
			if err != nil {
				return err
			}

		case 4:
			err = handler.GetSessions()
			if err != nil {
				return err
			}

		case 5:
			err = handler.GetTickets()
			if err != nil {
				return err
			}

		case 6:
			err = handler.NewMovie()
			if err != nil {
				return err
			}

		case 7:
			err = handler.NewCustomer()
			if err != nil {
				return err
			}

		case 8:
			err = handler.NewSession()
			if err != nil {
				return err
			}

		case 9:
			err = handler.NewTicket()
			if err != nil {
				return err
			}
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
	fmt.Println("\t~ 5. Get all tickets from database")

	fmt.Println("\t~ 6. Create new movie")
	fmt.Println("\t~ 7. Create new customer")
	fmt.Println("\t~ 8. Create new session")
	fmt.Println("\t~ 9. Create new ticket")
}
