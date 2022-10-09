package handler

import (
	"bufio"
	"fmt"
	"lab2/internal/domain"
	"lab2/internal/service"
	"os"
	"time"
)

type handler struct {
	service service.Service
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
	SearchHalls() (err error)
}

func NewHandler(service service.Service) Handler {
	return &handler{service: service}
}

func (h *handler) GetCustomers() error {
	customers, err := h.service.SelectCustomers()
	if err != nil {
		return err
	}

	for _, mov := range customers {
		fmt.Println("Відвідувач id =", mov.Id)
		fmt.Println("\tІм'я:", mov.FirstName)
		fmt.Println("\tПрізвище:", mov.LastName)
	}

	fmt.Println("Кількість відвідувачів =", len(customers))

	return nil
}

func (h *handler) GetMovies() error {
	movies, err := h.service.SelectMovies()
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

	return err
}

func (h *handler) GetHalls() error {
	halls, err := h.service.SelectHalls()
	if err != nil {
		return err
	}

	for _, hls := range halls {
		fmt.Println("Кінозал id =", hls.Id)
		fmt.Println("\tНазва:", hls.Title)
		fmt.Println("\tОпис:", hls.Description)
		fmt.Println("\tВмістимість:", hls.Capacity)
	}

	fmt.Println("Кількість кінозалів =", len(halls))

	return err
}

func (h *handler) GetSessions() error {
	sessions, err := h.service.SelectSessions()
	if err != nil {
		return err
	}

	for _, s := range sessions {
		fmt.Println("Сеанс id =", s.Id)
		fmt.Println("\tФільм:", s.MovieId)
		fmt.Println("\tКінозал:", s.HallId)
		fmt.Println("\tПочинається о:", s.StartAt)
	}

	fmt.Println("Кількість сеансів =", len(sessions))

	return err
}

func (h *handler) GetTickets() error {
	tickets, err := h.service.SelectTickets()
	if err != nil {
		return err
	}

	for _, t := range tickets {
		fmt.Println("Квиток id =", t.Id)
		fmt.Println("\tКористувач:", t.CustomerId)
		fmt.Println("\tСеанс:", t.SessionId)
		fmt.Println("\tЦіна:", t.Price)
		fmt.Println("\tРяд (id):", t.RowId)
		fmt.Println("\tМісце (id):", t.PositionId)
	}

	fmt.Println("Кількість квитків =", len(tickets))

	return err
}

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

	fmt.Println("Inserted movie id", movieId)

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

	fmt.Println("Inserted customer id", movieId)

	return nil
}

func (h *handler) NewSession() error {
	var session domain.Session

	fmt.Print("Enter movie id: ")
	_, err := fmt.Scan(&session.MovieId)
	if err != nil {
		return fmt.Errorf("invalid movie id input")
	}

	fmt.Print("Enter hall id: ")
	_, err = fmt.Scan(&session.HallId)
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

	fmt.Println("Inserted session id", sessionId)

	return nil
}

func (h *handler) NewTicket() error {
	var ticket domain.Ticket

	fmt.Print("Enter session id: ")
	_, err := fmt.Scan(&ticket.SessionId)
	if err != nil {
		return err
	}

	fmt.Print("Enter customer id: ")
	_, err = fmt.Scan(&ticket.CustomerId)
	if err != nil {
		return err
	}

	fmt.Print("Enter price: ")
	_, err = fmt.Scan(&ticket.Price)
	if err != nil {
		return err
	}

	fmt.Print("Enter row id: ")
	_, err = fmt.Scan(&ticket.RowId)
	if err != nil {
		return err
	}

	fmt.Print("Enter position id: ")
	_, err = fmt.Scan(&ticket.PositionId)
	if err != nil {
		return err
	}

	ticketId, err := h.service.CreateTicket(ticket)
	if err != nil {
		return err
	}

	fmt.Println("Inserted ticket id", ticketId)

	return nil
}

func (h *handler) SearchSessions() (err error) {
	var searchParams domain.SessionsSearchParams

	fmt.Print("Enter movie name: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	searchParams.MovieName = scanner.Text()

	var startAtGt string
	fmt.Print("Enter start time (ex.: 02/01/06,15:04): ")
	_, err = fmt.Scan(&startAtGt)
	if err != nil {
		return fmt.Errorf("invalid start at time input")
	}

	searchParams.StartAtGt, err = time.Parse("02/01/06,15:04", startAtGt)
	if err != nil {
		return fmt.Errorf("invalid start at time input")
	}

	var startAtLt string
	fmt.Print("Enter start time (ex.: 02/01/06,15:04): ")
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
		fmt.Println("Сеанс id =", s.Id)
		fmt.Println("\tНазва фільму:", s.Movie)
		fmt.Println("\tПочаток о:", s.StartAt)
		fmt.Println("\tКінозал:", s.Hall)
	}

	fmt.Println("Кількість сеансів =", len(sessions))

	fmt.Println("====================================")
	fmt.Println("Query time:", queryTime)
	fmt.Println("====================================")

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

	for _, t := range tickets {
		fmt.Println("Квиток id =", t.Id)
		fmt.Println("\tВідвідувач:", t.CustomerLastname, t.CustomerFirstname)
		fmt.Println("\tНазва фільму:", t.MovieTitle)
		fmt.Println("\tТривалість фільму:", t.MovieDuration)
		fmt.Println("\tПочинається о:", t.SessionStartAt)
		fmt.Println("\tКінозал:", t.HallTitle)
		fmt.Println("\tЦіна:", t.Price)
		fmt.Println("\tРяд:", t.Row)
		fmt.Println("\tМісце:", t.Position)
	}

	fmt.Println("Кількість квитків =", len(tickets))

	fmt.Println("====================================")
	fmt.Println("Query time:", queryTime)
	fmt.Println("====================================")

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
		fmt.Println("Кінозал id =", hl.Id)
		fmt.Println("\tНазва:", hl.Title)
		fmt.Println("\tОпис:", hl.Description)
		fmt.Println("\tРяди:", hl.Rows)
	}

	fmt.Println("Кількість кінозалів =", len(halls))

	fmt.Println("====================================")
	fmt.Println("Query time:", queryTime)
	fmt.Println("====================================")

	return err
}
