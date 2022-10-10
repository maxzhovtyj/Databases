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
			
		case 10:
			err = handler.SearchSessions()
			if err != nil {
				return err
			}

		case 11:
			err = handler.SearchTickets()
			if err != nil {
				return err
			}

		case 12:
			err = handler.SearchHalls()
			if err != nil {
				return err
			}

		case 13:
			err = handler.NewRandomMovies()
			if err != nil {
				return err
			}

		case 14:
			err = handler.NewRandomSessions()
			if err != nil {
				return err
			}

		case 15:
			err = handler.DeleteMovie()
			if err != nil {
				return err
			}

		case 16:
			err = handler.DeleteCustomer()
			if err != nil {
				return err
			}

		case 17:
			err = handler.DeleteSession()
			if err != nil {
				return err
			}

		case 18:
			err = handler.DeleteTicket()
			if err != nil {
				return err
			}

		case 19:
			err = handler.UpdateMovie()
			if err != nil {
				return err
			}

		case 20:
			err = handler.UpdateCustomer()
			if err != nil {
				return err
			}

		case 21:
			err = handler.UpdateSession()
			if err != nil {
				return err
			}

		case 22:
			err = handler.UpdateTicket()
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

	fmt.Println("")

	fmt.Println("\t~ 6. Create new movie")
	fmt.Println("\t~ 7. Create new customer")
	fmt.Println("\t~ 8. Create new session")
	fmt.Println("\t~ 9. Create new ticket")

	fmt.Println("")

	fmt.Println("\t~ 10. Search sessions")
	fmt.Println("\t~ 11. Search tickets")
	fmt.Println("\t~ 12. Search halls")

	fmt.Println("")

	fmt.Println("\t~ 13. Insert random movies")
	fmt.Println("\t~ 14. Insert random sessions")

	fmt.Println("")

	fmt.Println("\t~ 16. Delete movie")
	fmt.Println("\t~ 15. Delete customer")
	fmt.Println("\t~ 17. Delete session")
	fmt.Println("\t~ 18. Delete ticket")

	fmt.Println("")

	fmt.Println("\t~ 19. Edit movie")
	fmt.Println("\t~ 20. Edit customer")
	fmt.Println("\t~ 21. Edit session")
	fmt.Println("\t~ 22. Edit ticket")
}
