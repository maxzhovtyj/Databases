package handler

import (
	"bufio"
	"fmt"
	"lab3/internal/domain"
	"lab3/internal/service"
	"os"
	"strings"
	"time"
)

type handler struct {
	service service.Service
	logger  *bufio.Writer
}

type Handler interface {
	GetCustomers() error
	GetMovies() error
	GetHalls() error
	GetSessions() error
	GetTickets() error
	NewMovie() error
	NewCustomer() error
	NewSession() error
	NewTicket() error
	SearchSessions() error
	SearchTickets() error
	SearchHalls() error
	NewRandomMovies() error
	NewRandomSessions() error
	DeleteCustomer() error
	DeleteMovie() error
	DeleteSession() error
	DeleteTicket() error
	UpdateCustomer() error
	UpdateMovie() error
	UpdateSession() error
	UpdateTicket() error
}

func NewHandler(service service.Service, logger *bufio.Writer) Handler {
	return &handler{
		service: service,
		logger:  logger,
	}
}

func (h *handler) GetCustomers() error {
	customers, err := h.service.SelectCustomers()
	if err != nil {
		return err
	}

	for _, mov := range customers {
		_, _ = fmt.Fprintln(h.logger, "Відвідувач id =", mov.Model.ID)
		_, _ = fmt.Fprintln(h.logger, "\tІм'я:", mov.FirstName)
		_, _ = fmt.Fprintln(h.logger, "\tПрізвище:", mov.LastName)
	}

	_, err = fmt.Fprintln(h.logger, "Кількість відвідувачів =", len(customers))
	if err != nil {
		return err
	}

	err = h.logger.Flush()
	if err != nil {
		return err
	}

	return nil
}

func (h *handler) GetMovies() error {
	movies, err := h.service.SelectMovies()
	if err != nil {
		return err
	}

	for _, mov := range movies {
		_, _ = fmt.Fprintln(h.logger, "Фільм id =", mov.Model.ID)
		_, _ = fmt.Fprintln(h.logger, "\tНазва:", mov.Title)
		_, _ = fmt.Fprintln(h.logger, "\tОпис:", mov.Description)
		_, _ = fmt.Fprintln(h.logger, "\tТривалість:", mov.Duration)
	}

	_, _ = fmt.Fprintln(h.logger, "Кількість фільмів =", len(movies))

	err = h.logger.Flush()
	if err != nil {
		return err
	}

	return err
}

func (h *handler) GetHalls() error {
	halls, err := h.service.SelectHalls()
	if err != nil {
		return err
	}

	for _, hls := range halls {
		_, _ = fmt.Fprintln(h.logger, "Кінозал id =", hls.Model.ID)
		_, _ = fmt.Fprintln(h.logger, "\tНазва:", hls.Title)
		_, _ = fmt.Fprintln(h.logger, "\tОпис:", hls.Description)
		_, _ = fmt.Fprintln(h.logger, "\tВмістимість:", hls.Capacity)
	}

	_, _ = fmt.Fprintln(h.logger, "Кількість кінозалів =", len(halls))

	err = h.logger.Flush()
	if err != nil {
		return err
	}

	return err
}

func (h *handler) GetSessions() error {
	sessions, err := h.service.SelectSessions()
	if err != nil {
		return err
	}

	for _, s := range sessions {
		_, _ = fmt.Fprintln(h.logger, "Сеанс id =", s.Model.ID)
		_, _ = fmt.Fprintln(h.logger, "\tФільм:", s.MovieID)
		_, _ = fmt.Fprintln(h.logger, "\tКінозал:", s.HallID)
		_, _ = fmt.Fprintln(h.logger, "\tПочинається о:", s.StartAt)
	}

	_, _ = fmt.Fprintln(h.logger, "Кількість сеансів =", len(sessions))

	err = h.logger.Flush()
	if err != nil {
		return err
	}
	return err
}

func (h *handler) GetTickets() error {
	tickets, err := h.service.SelectTickets()
	if err != nil {
		return err
	}

	for _, t := range tickets {
		_, _ = fmt.Fprintln(h.logger, "Квиток id =", t.Model.ID)
		_, _ = fmt.Fprintln(h.logger, "\tКористувач:", t.CustomerID)
		_, _ = fmt.Fprintln(h.logger, "\tСеанс:", t.SessionID)
		_, _ = fmt.Fprintln(h.logger, "\tЦіна:", t.Price)
		_, _ = fmt.Fprintln(h.logger, "\tРяд (id):", t.RowID)
		_, _ = fmt.Fprintln(h.logger, "\tМісце (id):", t.PositionID)
	}

	_, _ = fmt.Fprintln(h.logger, "Кількість квитків =", len(tickets))

	err = h.logger.Flush()
	if err != nil {
		return err
	}

	return err
}

//

func (h *handler) NewMovie() error {
	fmt.Print("Enter movie name: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	movieName := scanner.Text()

	fmt.Print("Enter movie description: ")
	scanner.Scan()
	movieDescription := scanner.Text()

	var movieDuration string
	fmt.Print("Enter movie duration (ex: 1:20:30): ")
	_, err := fmt.Scan(&movieDuration)
	if err != nil {
		return err
	}

	movieId, err := h.service.CreateMovie(domain.Movie{
		Title:       movieName,
		Description: movieDescription,
		Duration:    movieDuration,
	})
	if err != nil {
		return err
	}

	_, _ = fmt.Fprintln(h.logger, "Inserted movie id", movieId)

	err = h.logger.Flush()
	if err != nil {
		return err
	}

	return nil
}

func (h *handler) NewCustomer() error {
	fmt.Print("Enter customer firstname: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	customerFirstName := scanner.Text()

	fmt.Print("Enter customer lastname: ")
	scanner.Scan()
	customerLastname := scanner.Text()

	movieId, err := h.service.CreateCustomer(domain.Customer{
		FirstName: customerFirstName,
		LastName:  customerLastname,
	})
	if err != nil {
		return err
	}

	_, _ = fmt.Fprintln(h.logger, "Inserted customer id", movieId)

	err = h.logger.Flush()
	if err != nil {
		return err
	}

	return nil
}

func (h *handler) NewSession() error {
	var session domain.Session

	fmt.Print("Enter movie id: ")
	_, err := fmt.Scan(&session.MovieID)
	if err != nil {
		return fmt.Errorf("invalid movie id input")
	}

	fmt.Print("Enter hall id: ")
	_, err = fmt.Scan(&session.HallID)
	if err != nil {
		return fmt.Errorf("invalid hall id input")
	}

	var startAt string
	fmt.Print("Enter start time (ex.: 02/01/06,15:04): ")
	_, err = fmt.Scan(&startAt)
	if err != nil {
		return fmt.Errorf("invalid start at time input")
	}

	session.StartAt, err = time.Parse("02/01/06,15:04", startAt)
	if err != nil {
		return fmt.Errorf("invalid start at time input")
	}

	sessionId, err := h.service.CreateSession(session)
	if err != nil {
		return err
	}

	_, _ = fmt.Fprintln(h.logger, "Inserted session id", sessionId)

	err = h.logger.Flush()
	if err != nil {
		return err
	}
	return nil
}

func (h *handler) NewTicket() error {
	var ticket domain.Ticket

	fmt.Print("Enter session id: ")
	_, err := fmt.Scan(&ticket.SessionID)
	if err != nil {
		return err
	}

	fmt.Print("Enter customer id: ")
	_, err = fmt.Scan(&ticket.CustomerID)
	if err != nil {
		return err
	}

	fmt.Print("Enter price: ")
	_, err = fmt.Scan(&ticket.Price)
	if err != nil {
		return err
	}

	fmt.Print("Enter row id: ")
	_, err = fmt.Scan(&ticket.RowID)
	if err != nil {
		return err
	}

	fmt.Print("Enter position id: ")
	_, err = fmt.Scan(&ticket.PositionID)
	if err != nil {
		return err
	}

	ticketId, err := h.service.CreateTicket(ticket)
	if err != nil {
		return err
	}

	_, _ = fmt.Fprintln(h.logger, "Inserted ticket id", ticketId)

	err = h.logger.Flush()
	if err != nil {
		return err
	}
	return nil
}

//

func (h *handler) SearchSessions() (err error) {
	var searchParams domain.SessionsSearchParams

	fmt.Print("Enter movie name: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	searchParams.MovieName = scanner.Text()

	var startAtGt string
	fmt.Print("Enter greater than (ex.: 02/01/06,15:04): ")
	_, err = fmt.Scan(&startAtGt)
	if err != nil {
		return fmt.Errorf("invalid start at time input")
	}

	searchParams.StartAtGt, err = time.Parse("02/01/06,15:04", startAtGt)
	if err != nil {
		return fmt.Errorf("invalid start at time input")
	}

	var startAtLt string
	fmt.Print("Enter lower than time (ex.: 02/01/06,15:04): ")
	_, err = fmt.Scan(&startAtLt)
	if err != nil {
		return fmt.Errorf("invalid start at time input")
	}

	searchParams.StartAtLt, err = time.Parse("02/01/06,15:04", startAtLt)
	if err != nil {
		return fmt.Errorf("invalid start at time input")
	}

	sessions, queryTime, err := h.service.SearchSessions(searchParams)
	if err != nil {
		return err
	}

	for _, s := range sessions {
		_, _ = fmt.Fprintln(h.logger, "Сеанс id =", s.Model.ID)
		_, _ = fmt.Fprintln(h.logger, "\tНазва фільму:", s.MovieID)
		_, _ = fmt.Fprintln(h.logger, "\tПочаток о:", s.StartAt)
		_, _ = fmt.Fprintln(h.logger, "\tКінозал:", s.HallID)
	}

	fmt.Println("Кількість сеансів =", len(sessions))

	_, _ = fmt.Fprintln(h.logger, "====================================")
	_, _ = fmt.Fprintln(h.logger, "Query time:", queryTime)
	_, _ = fmt.Fprintln(h.logger, "====================================")

	err = h.logger.Flush()
	if err != nil {
		return err
	}
	return err
}

func (h *handler) SearchTickets() (err error) {
	var searchParams domain.TicketsSearchParams

	fmt.Println("Enter ticket price greater than: ")
	_, err = fmt.Scan(&searchParams.PriceGt)
	if err != nil {
		return err
	}

	fmt.Println("Enter ticket price lower than: ")
	_, err = fmt.Scan(&searchParams.PriceLt)
	if err != nil {
		return err
	}

	fmt.Println("Enter duration greater than (ex: 1:20:30): ")
	_, err = fmt.Scan(&searchParams.MovieDurationGt)
	if err != nil {
		return fmt.Errorf("invalid duration input")
	}

	fmt.Println("Enter duration lower than (ex: 1:20:30): ")
	_, err = fmt.Scan(&searchParams.MovieDurationLt)
	if err != nil {
		return fmt.Errorf("invalid duration input")
	}

	tickets, queryTime, err := h.service.SearchTickets(searchParams)
	if err != nil {
		return err
	}

	_, _ = fmt.Fprintln(h.logger, "Квитки: \n", tickets)

	_, _ = fmt.Fprintln(h.logger, "Кількість квитків =", len(tickets))

	_, _ = fmt.Fprintln(h.logger, "====================================")
	_, _ = fmt.Fprintln(h.logger, "Query time:", queryTime)
	_, _ = fmt.Fprintln(h.logger, "====================================")

	err = h.logger.Flush()
	if err != nil {
		return err
	}
	return err
}

func (h *handler) SearchHalls() (err error) {
	var searchParams domain.HallsSearchParams

	fmt.Print("Enter hall title: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	searchParams.HallTitle = scanner.Text()

	fmt.Print("Enter capacity greater than: ")
	_, err = fmt.Scan(&searchParams.CapacityGt)
	if err != nil {
		return err
	}

	fmt.Print("Enter capacity lower than: ")
	_, err = fmt.Scan(&searchParams.CapacityLt)
	if err != nil {
		return err
	}

	halls, queryTime, err := h.service.SearchHalls(domain.HallsSearchParams{
		HallTitle:  "par",
		CapacityGt: 0,
		CapacityLt: 1000,
	})
	if err != nil {
		return err
	}

	for _, hl := range halls {
		_, _ = fmt.Fprintln(h.logger, "Кінозал id =", hl.Model.ID)
		_, _ = fmt.Fprintln(h.logger, "\tНазва:", hl.Title)
		_, _ = fmt.Fprintln(h.logger, "\tОпис:", hl.Description)
		_, _ = fmt.Fprintln(h.logger, "\tРяди:", hl.Capacity)
	}

	_, _ = fmt.Fprintln(h.logger, "Кількість кінозалів =", len(halls))

	_, _ = fmt.Fprintln(h.logger, "====================================")
	_, _ = fmt.Fprintln(h.logger, "Query time:", queryTime)
	_, _ = fmt.Fprintln(h.logger, "====================================")

	err = h.logger.Flush()
	if err != nil {
		return err
	}
	return err
}

//

func (h *handler) NewRandomMovies() error {
	var movieAmount int
	fmt.Print("Enter inserted movie amount: ")
	_, err := fmt.Scan(&movieAmount)
	if err != nil {
		return fmt.Errorf("invalid amount")
	}

	err = h.service.CreateRandomMovies(movieAmount)
	if err != nil {
		return err
	}

	fmt.Println("Successfully inserted:", movieAmount)

	return err
}

func (h *handler) NewRandomSessions() error {
	var sessionsAmount int
	fmt.Print("Enter inserted sessions amount: ")
	_, err := fmt.Scan(&sessionsAmount)
	if err != nil {
		return fmt.Errorf("invalid amount")
	}

	err = h.service.CreateRandomSessions(sessionsAmount)
	if err != nil {
		return err
	}

	fmt.Println("Successfully inserted:", sessionsAmount)

	return err
}

//

func (h *handler) DeleteCustomer() error {
	var id int
	fmt.Print("Enter customer id you want to delete: ")
	_, err := fmt.Scan(&id)
	if err != nil {
		return err
	}

	var option string
	fmt.Print("Deleting customer will lead to deleting all customer tickets. Are you sure (y/n): ")
	_, err = fmt.Scan(&option)
	if err != nil {
		return err
	}

	if strings.ToLower(option) == "y" {
		err = h.service.DeleteCustomer(id)
		if err != nil {
			return err
		}

		fmt.Printf("Customer id = %d successfully deleted\n", id)
	} else {
		return nil
	}

	return err
}

func (h *handler) DeleteMovie() error {
	var id int
	fmt.Print("Enter movie id you want to delete: ")
	_, err := fmt.Scan(&id)
	if err != nil {
		return err
	}

	var option string
	fmt.Print("Deleting movie will lead to deleting all sessions and tickets. Are you sure (y/n): ")
	_, err = fmt.Scan(&option)
	if err != nil {
		return err
	}

	if strings.ToLower(option) == "y" {
		err = h.service.DeleteMovie(id)
		if err != nil {
			return err
		}
	} else {
		return nil
	}

	fmt.Printf("Movie id = %d successfully deleted\n", id)
	return err
}

func (h *handler) DeleteSession() error {
	var id int
	fmt.Print("Enter session id you want to delete: ")
	_, err := fmt.Scan(&id)
	if err != nil {
		return err
	}

	var option string
	fmt.Print("Deleting session will lead to deleting all tickets. Are you sure (y/n): ")
	_, err = fmt.Scan(&option)
	if err != nil {
		return err
	}

	if strings.ToLower(option) == "y" {
		err = h.service.DeleteSession(id)
		if err != nil {
			return err
		}
	} else {
		return nil
	}

	fmt.Printf("Session %d successfully deleted\n", id)
	return err
}

func (h *handler) DeleteTicket() error {
	var id int
	fmt.Print("Enter session id you want to delete: ")
	_, err := fmt.Scan(&id)
	if err != nil {
		return err
	}

	var option string
	fmt.Print("Are you sure you want to delete ticket (y/n): ")
	_, err = fmt.Scan(&option)
	if err != nil {
		return err
	}

	if strings.ToLower(option) == "y" {
		err = h.service.DeleteTicket(id)
		if err != nil {
			return err
		}
	} else {
		return nil
	}

	fmt.Printf("Ticket %d successfully deleted\n", id)
	return err
}

//

func (h *handler) UpdateCustomer() error {
	var customer domain.Customer
	fmt.Print("Enter customer id you want to update: ")
	_, err := fmt.Scan(&customer.Model.ID)
	if err != nil {
		return err
	}

	fmt.Print("Enter customer new firstname: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	customer.FirstName = scanner.Text()

	fmt.Print("Enter customer new lastname: ")
	scanner.Scan()
	customer.LastName = scanner.Text()

	err = h.service.UpdateCustomer(customer)
	if err != nil {
		return err
	}

	fmt.Println("Customer successfully updated")
	return err
}

func (h *handler) UpdateMovie() error {
	var movie domain.Movie
	fmt.Print("Enter movie id you want to update: ")
	_, err := fmt.Scan(&movie.Model.ID)
	if err != nil {
		return err
	}

	fmt.Print("Enter movie new title: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	movie.Title = scanner.Text()

	fmt.Print("Enter movie new description: ")
	scanner.Scan()
	movie.Description = scanner.Text()

	fmt.Print("Enter movie new duration (ex: 1:20:30): ")
	_, err = fmt.Scan(&movie.Duration)
	if err != nil {
		return err
	}

	err = h.service.UpdateMovie(movie)
	if err != nil {
		return err
	}

	fmt.Printf("Movie id = %d successfully updated\n", movie.Model.ID)
	return err
}

func (h *handler) UpdateSession() error {
	var session domain.Session
	fmt.Print("Enter session id you want to update: ")
	_, err := fmt.Scan(&session.Model.ID)
	if err != nil {
		return err
	}

	fmt.Print("Enter session new movie id: ")
	_, err = fmt.Scan(&session.MovieID)
	if err != nil {
		return err
	}

	fmt.Print("Enter session new description: ")
	_, err = fmt.Scan(&session.HallID)
	if err != nil {
		return err
	}

	var startAt string
	fmt.Print("Enter start time (ex.: 02/01/06,15:04): ")
	_, err = fmt.Scan(&startAt)
	if err != nil {
		return fmt.Errorf("invalid start at time input")
	}

	session.StartAt, err = time.Parse("02/01/06,15:04", startAt)
	if err != nil {
		return fmt.Errorf("invalid start at time input")
	}

	err = h.service.UpdateSession(session)
	if err != nil {
		return err
	}

	fmt.Printf("Session id = %d successfully updated\n", session.Model.ID)
	return err
}

func (h *handler) UpdateTicket() error {
	var ticket domain.Ticket
	fmt.Print("Enter ticket id you want to update: ")
	_, err := fmt.Scan(&ticket.Model.ID)
	if err != nil {
		return err
	}

	fmt.Print("Enter ticket new customer id: ")
	_, err = fmt.Scan(&ticket.CustomerID)
	if err != nil {
		return err
	}

	fmt.Print("Enter ticket new session id: ")
	_, err = fmt.Scan(&ticket.SessionID)
	if err != nil {
		return err
	}

	fmt.Print("Enter ticket new price: ")
	_, err = fmt.Scan(&ticket.Price)
	if err != nil {
		return err
	}

	fmt.Print("Enter ticket new row id: ")
	_, err = fmt.Scan(&ticket.RowID)
	if err != nil {
		return err
	}

	fmt.Print("Enter ticket new position id: ")
	_, err = fmt.Scan(&ticket.PositionID)
	if err != nil {
		return err
	}

	err = h.service.UpdateTicket(ticket)
	if err != nil {
		return err
	}

	fmt.Printf("Ticket id = %d successfully updated\n", ticket.Model.ID)
	return err
}
