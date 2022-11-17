package main

import (
	"bufio"
	"fmt"
	"lab3/internal/domain"
	"lab3/internal/handler"
	"lab3/internal/repository"
	"lab3/internal/service"
	"lab3/pkg/client/orm"
	"log"
	"os"
)

func main() {
	dbClient, err := orm.NewORMClient(&orm.StorageConfig{
		Username: "postgres",
		Password: "30042003",
		Host:     "localhost",
		Port:     "5432",
		Database: "Lab1",
	})
	if err != nil {
		log.Fatalf("failed to connect to postgres database, %v", err)
	}

	f, err := os.Create("./log/out.log")
	if err != nil {
		log.Fatalf("failed to open log file, error: %v", err)
		return
	}

	writer := bufio.NewWriter(f)

	defer func(f *os.File) {
		err = f.Close()
		if err != nil {
			log.Fatalf("failed to close log file, error: %v", err)
			return
		}
	}(f)

	err = dbClient.AutoMigrate(
		&domain.Customer{},
		&domain.Hall{},
		&domain.Movie{},
		&domain.Position{},
		&domain.Row{},
		&domain.Session{},
		&domain.Ticket{},
	)
	if err != nil {
		log.Fatalf("failed to apply migrations, %v", err)
		return
	}

	repo := repository.NewRepository(dbClient)
	serviceInstance := service.NewService(repo)
	handlerInstance := handler.NewHandler(serviceInstance, writer)

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
				fmt.Println(err.Error())
			}

		case 2:
			err = handler.GetMovies()
			if err != nil {
				fmt.Println(err.Error())
			}

		case 3:
			err = handler.GetHalls()
			if err != nil {
				fmt.Println(err.Error())
			}

		case 4:
			err = handler.GetSessions()
			if err != nil {
				fmt.Println(err.Error())
			}

		case 5:
			err = handler.GetTickets()
			if err != nil {
				fmt.Println(err.Error())
			}

		case 6:
			err = handler.NewMovie()
			if err != nil {
				fmt.Println(err.Error())
			}

		case 7:
			err = handler.NewCustomer()
			if err != nil {
				fmt.Println(err.Error())
			}

		case 8:
			err = handler.NewSession()
			if err != nil {
				fmt.Println(err.Error())
			}

		case 9:
			err = handler.NewTicket()
			if err != nil {
				fmt.Println(err.Error())
			}

		case 10:
			err = handler.SearchSessions()
			if err != nil {
				fmt.Println(err.Error())
			}

		case 11:
			err = handler.SearchTickets()
			if err != nil {
				fmt.Println(err.Error())
			}

		case 12:
			err = handler.SearchHalls()
			if err != nil {
				fmt.Println(err.Error())
			}

		case 13:
			err = handler.NewRandomMovies()
			if err != nil {
				fmt.Println(err.Error())
			}

		case 14:
			err = handler.NewRandomSessions()
			if err != nil {
				fmt.Println(err.Error())
			}

		case 15:
			err = handler.DeleteMovie()
			if err != nil {
				fmt.Println(err.Error())
			}

		case 16:
			err = handler.DeleteCustomer()
			if err != nil {
				fmt.Println(err.Error())
			}

		case 17:
			err = handler.DeleteSession()
			if err != nil {
				fmt.Println(err.Error())
			}

		case 18:
			err = handler.DeleteTicket()
			if err != nil {
				fmt.Println(err.Error())
			}

		case 19:
			err = handler.UpdateMovie()
			if err != nil {
				fmt.Println(err.Error())
			}

		case 20:
			err = handler.UpdateCustomer()
			if err != nil {
				fmt.Println(err.Error())
			}

		case 21:
			err = handler.UpdateSession()
			if err != nil {
				fmt.Println(err.Error())
			}

		case 22:
			err = handler.UpdateTicket()
			if err != nil {
				fmt.Println(err.Error())
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

	fmt.Println("\t~ 15. Delete movie")
	fmt.Println("\t~ 16. Delete customer")
	fmt.Println("\t~ 17. Delete session")
	fmt.Println("\t~ 18. Delete ticket")

	fmt.Println("")

	fmt.Println("\t~ 19. Edit movie")
	fmt.Println("\t~ 20. Edit customer")
	fmt.Println("\t~ 21. Edit session")
	fmt.Println("\t~ 22. Edit ticket")
}
