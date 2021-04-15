package data

import (
	"database/sql"
	"fmt"
)

type Account struct {
	Id          int
	FirstName   string
	LastName    string
	PhoneNumber string
	Email       string
}

type Ticket struct {
	Id                  int
	PerformanceName     string
	GenreName           string
	PerformanceDuration string
	DateTime            string
	HallName            string
	HallCapacity        int
	LocationAddress     string
	LocationPhoneNumber string
	SectorName          string
	Place               int
	Price               int
	DateOfIssue         string
	Paid                bool
	Reservation         bool
	Destroyed           bool
}

type TicketInsert struct {
	Id          int
	AccountId   int
	ScheduleId  int
	PlaceId     int
	DateOfIssue string
	Paid        bool
	Reservation bool
	Destroyed   bool
}

type Poster struct {
	Id                  int
	PerformanceName     string
	GenreName           string
	PerformanceDuration string
	DateTime            string
	HallName            string
	HallCapacity        int
	LocationAddress     string
	LocationPhoneNumber string
	Comment             string
}

type User struct {
	Id                  int
	FirstName           string
	LastName            string
	Role                string
	LocationAddress     string
	LocationPhoneNumber string
	PhoneNumber         string
}

type UserInsert struct {
	Id          int
	AccountId   int
	FirstName   string
	LastName    string
	RoleId      int
	LocationId  int
	PhoneNumber string
}

type Hall struct {
	Id         int
	AccountId  int
	Name       string
	Capacity   int
	LocationId int
}

type Location struct {
	Id          int
	AccountId   int
	Address     string
	PhoneNumber string
}

type Performance struct {
	Id        int
	AccountId int
	Name      string
	GenreId   int
	Duration  string
}

type Place struct {
	id       int
	SectorId int
	Name     string
}

type PosterInsert struct {
	Id         int
	AccountId  int
	ScheduleId int
	Comment    string
}

type Price struct {
	Id            int
	AccountId     int
	SectorId      int
	PerformanceId int
	Price         int
}

type Role struct {
	Id   int
	Name string
}

type Schedule struct {
	Id            int
	AccountId     int
	PerformanceId int
	Date          string
	HallId        int
}

type Sector struct {
	Id   int
	Name string
}

type Genre struct {
	Id   int
	Name string
}

type TheaterData struct {
	db *sql.DB
}

func NewTheaterData(db *sql.DB) *TheaterData {
	return &TheaterData{db: db}
}

func (u TheaterData) ReadAllTickets() ([]Ticket, error) {
	var tickets []Ticket
	rows, err := u.db.Query(readAllTicketsQuery)
	if err != nil {
		return nil, fmt.Errorf("can't get tickets from database, error:%w", err)
	}
	for rows.Next() {
		var temp Ticket
		err = rows.Scan(&temp.Id, &temp.PerformanceName, &temp.GenreName, &temp.PerformanceDuration,
			&temp.DateTime, &temp.HallName, &temp.HallCapacity, &temp.LocationAddress, &temp.LocationPhoneNumber,
			&temp.SectorName, &temp.Place, &temp.Price, &temp.DateOfIssue, &temp.Paid, &temp.Reservation, &temp.Destroyed)
		if err != nil {
			return nil, fmt.Errorf("can't scan tickets from database, error:%w", err)
		}
		tickets = append(tickets, temp)
	}
	return tickets, nil
}

func (u TheaterData) ReadAllPosters() ([]Poster, error) {
	var posters []Poster
	rows, err := u.db.Query(readAllPostersQuery)
	if err != nil {
		return nil, fmt.Errorf("can't get posters from database, error:%w", err)
	}
	for rows.Next() {
		var temp Poster
		err = rows.Scan(&temp.Id, &temp.PerformanceName, &temp.GenreName, &temp.PerformanceDuration,
			&temp.DateTime, &temp.HallName, &temp.HallCapacity, &temp.LocationAddress, &temp.LocationPhoneNumber,
			&temp.Comment)
		if err != nil {
			return nil, fmt.Errorf("can't scan posters from database, error:%w", err)
		}
		posters = append(posters, temp)
	}
	return posters, nil
}

func (u TheaterData) ReadAllUsers(account Account) ([]User, error) {
	var users []User
	rows, err := u.db.Query(readAllUsersQuery, account.Id)
	if err != nil {
		return nil, fmt.Errorf("can't get users from database, error:%w", err)
	}
	for rows.Next() {
		var temp User
		err = rows.Scan(&temp.Id, &temp.FirstName, &temp.LastName, &temp.Role,
			&temp.LocationAddress, &temp.LocationPhoneNumber, &temp.PhoneNumber)
		if err != nil {
			return nil, fmt.Errorf("can't scan users from database, error:%w", err)
		}
		users = append(users, temp)
	}
	return users, nil
}

func (u TheaterData) AddAccount(account Account) error {
	_, err := u.db.Exec(insertAccount, account.FirstName, account.LastName, account.PhoneNumber, account.Email)
	if err != nil {
		return fmt.Errorf("can't inser account to database, error: %w", err)
	}
	return nil
}

func (u TheaterData) AddGenre(genre Genre) error {
	_, err := u.db.Exec(insertGenre, genre.Name)
	if err != nil {
		return fmt.Errorf("can't inser genre to database, error: %w", err)
	}
	return nil
}

func (u TheaterData) AddHall(hall Hall) error {
	_, err := u.db.Exec(insertHall, hall.AccountId, hall.Name, hall.Capacity, hall.LocationId)
	if err != nil {
		return fmt.Errorf("can't inser hall to database, error: %w", err)
	}
	return nil
}

func (u TheaterData) AddLocation(location Location) error {
	_, err := u.db.Exec(insertLocation, location.AccountId, location.Address, location.PhoneNumber)
	if err != nil {
		return fmt.Errorf("can't inser location to database, error: %w", err)
	}
	return nil
}

func (u TheaterData) AddPerformance(performance Performance) error {
	_, err := u.db.Exec(insertPerformance, performance.AccountId, performance.Name, performance.GenreId, performance.Duration)
	if err != nil {
		return fmt.Errorf("can't inser Performance to database, error: %w", err)
	}
	return nil
}

func (u TheaterData) AddPlace(place Place) error {
	_, err := u.db.Exec(insertPlace, place.SectorId, place.Name)
	if err != nil {
		return fmt.Errorf("can't inser Place to database, error: %w", err)
	}
	return nil
}

func (u TheaterData) AddPoster(poster PosterInsert) error {
	_, err := u.db.Exec(insertPoster, poster.AccountId, poster.ScheduleId, poster.Comment)
	if err != nil {
		return fmt.Errorf("can't inser Poster to database, error: %w", err)
	}
	return nil
}

func (u TheaterData) AddPrice(price Price) error {
	_, err := u.db.Exec(insertPrice, price.AccountId, price.SectorId, price.PerformanceId, price.Price)
	if err != nil {
		return fmt.Errorf("can't inser Price to database, error: %w", err)
	}
	return nil
}

func (u TheaterData) AddRole(role Role) error {
	_, err := u.db.Exec(insertRole, role.Name)
	if err != nil {
		return fmt.Errorf("can't inser Role to database, error: %w", err)
	}
	return nil
}

func (u TheaterData) AddSchedule(schedule Schedule) error {
	_, err := u.db.Exec(insertSchedule, schedule.AccountId, schedule.PerformanceId, schedule.Date, schedule.HallId)
	if err != nil {
		return fmt.Errorf("can't inser Schedule to database, error: %w", err)
	}
	return nil
}

func (u TheaterData) AddSector(sector Sector) error {
	_, err := u.db.Exec(insertSector, sector.Name)
	if err != nil {
		return fmt.Errorf("can't inser Sector to database, error: %w", err)
	}
	return nil
}

func (u TheaterData) AddTicket(ticket TicketInsert) error {
	_, err := u.db.Exec(insertTicket, ticket.AccountId, ticket.ScheduleId,
		ticket.PlaceId, ticket.DateOfIssue, ticket.Paid, ticket.Reservation, ticket.Destroyed)
	if err != nil {
		return fmt.Errorf("can't inser Ticket to database, error: %w", err)
	}
	return nil
}

func (u TheaterData) AddUser(user UserInsert) error {
	_, err := u.db.Exec(insertUser, user.AccountId, user.FirstName, user.LastName, user.RoleId, user.LocationId, user.PhoneNumber)
	if err != nil {
		return fmt.Errorf("can't inser User to database, error: %w", err)
	}
	return nil
}

func (u TheaterData) DeleteEntry(table Table, id int) error {
	_, err := u.db.Exec(deleteBegin+table.String()+deleteEnd, id)
	if err != nil {
		return fmt.Errorf("can't Delete to database, error: %w", err)
	}
	return nil
}

func (u TheaterData) UpdateAccount(account Account) error {
	_, err := u.db.Exec(updateAccount, account.FirstName, account.LastName, account.PhoneNumber, account.Email, account.Id)
	if err != nil {
		return fmt.Errorf("can't update account to database, error: %w", err)
	}
	return nil
}

func (u TheaterData) UpdateGenre(genre Genre) error {
	_, err := u.db.Exec(updateGenre, genre.Name, genre.Id)
	if err != nil {
		return fmt.Errorf("can't update genre to database, error: %w", err)
	}
	return nil
}

func (u TheaterData) UpdateHall(hall Hall) error {
	_, err := u.db.Exec(updateHall, hall.AccountId, hall.Name, hall.Capacity, hall.LocationId, hall.Id)
	if err != nil {
		return fmt.Errorf("can't update hall to database, error: %w", err)
	}
	return nil
}

func (u TheaterData) UpdateLocation(location Location) error {
	_, err := u.db.Exec(updateLocation, location.AccountId, location.Address, location.PhoneNumber, location.Id)
	if err != nil {
		return fmt.Errorf("can't update location to database, error: %w", err)
	}
	return nil
}

func (u TheaterData) UpdatePerformance(performance Performance) error {
	_, err := u.db.Exec(updatePerformance, performance.AccountId, performance.Name, performance.GenreId, performance.Duration, performance.Id)
	if err != nil {
		return fmt.Errorf("can't update Performance to database, error: %w", err)
	}
	return nil
}

func (u TheaterData) UpdatePlace(place Place) error {
	_, err := u.db.Exec(updatePlace, place.SectorId, place.Name, place.id)
	if err != nil {
		return fmt.Errorf("can't update Place to database, error: %w", err)
	}
	return nil
}

func (u TheaterData) UpdatePoster(poster PosterInsert) error {
	_, err := u.db.Exec(updatePoster, poster.AccountId, poster.ScheduleId, poster.Comment, poster.Id)
	if err != nil {
		return fmt.Errorf("can't update Poster to database, error: %w", err)
	}
	return nil
}

func (u TheaterData) UpdatePrice(price Price) error {
	_, err := u.db.Exec(updatePrice, price.AccountId, price.SectorId, price.PerformanceId, price.Price, price.Id)
	if err != nil {
		return fmt.Errorf("can't update Price to database, error: %w", err)
	}
	return nil
}

func (u TheaterData) UpdateRole(role Role) error {
	_, err := u.db.Exec(updateRole, role.Name, role.Id)
	if err != nil {
		return fmt.Errorf("can't update Role to database, error: %w", err)
	}
	return nil
}

func (u TheaterData) UpdateSchedule(schedule Schedule) error {
	_, err := u.db.Exec(updateSchedule, schedule.AccountId, schedule.PerformanceId, schedule.Date, schedule.HallId, schedule.Id)
	if err != nil {
		return fmt.Errorf("can't update Schedule to database, error: %w", err)
	}
	return nil
}

func (u TheaterData) UpdateSector(sector Sector) error {
	_, err := u.db.Exec(updateSector, sector.Name, sector.Id)
	if err != nil {
		return fmt.Errorf("can't update Sector to database, error: %w", err)
	}
	return nil
}

func (u TheaterData) UpdateTicket(ticket TicketInsert) error {
	_, err := u.db.Exec(updateTicket, ticket.AccountId, ticket.ScheduleId,
		ticket.PlaceId, ticket.DateOfIssue, ticket.Paid, ticket.Reservation, ticket.Destroyed, ticket.Id)
	if err != nil {
		return fmt.Errorf("can't update Ticket to database, error: %w", err)
	}
	return nil
}

func (u TheaterData) UpdateUser(user UserInsert) error {
	_, err := u.db.Exec(updateUser, user.AccountId, user.FirstName, user.LastName, user.RoleId,
		user.LocationId, user.PhoneNumber, user.Id)
	if err != nil {
		return fmt.Errorf("can't update User to database, error: %w", err)
	}
	return nil
}
