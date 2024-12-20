package dbrepo

import (
	"errors"
	"time"

	"github.com/jordanryanoFA/bookings/internal/models"
)

func (m *tesDBRepo) AllUsers() bool {
	return true
}

// InsertReservation inserts a reservation into the database
func (m *tesDBRepo) InsertReservation(res models.Reservation) (int, error) {
	// if the room id is 2, then fail; otherwise, pass
	if res.RoomID == 2 {
		return 0, errors.New("some error)")
	}
	return 1, nil
}

// InsertRoomRestriction insert a room restriction into the database
func (m *tesDBRepo) InsertRoomRestriction(r models.RoomRestriction) error {
	if r.RoomID == 1000 {
		return errors.New("some error")
	}
	return nil
}

// SearchAvailabilityByDatesByRoomID returns true if availability exists for roomID, returns false if availibility not available
func (m *tesDBRepo) SearchAvailabilityByDatesByRoomID(start, end time.Time, roomID int) (bool, error) {
	return false, nil
}

// SearchAvailabilityForAllRooms returns a slice of available rooms, if any, for given date range
func (m *tesDBRepo) SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error) {
	var rooms []models.Room
	return rooms, nil
}

// GetRoomByID gets a room by id
func (m *tesDBRepo) GetRoomByID(id int) (models.Room, error) {
	var room models.Room
	if id > 2 {
		return room, errors.New("some error")
	}
	return room, nil
}

func (m *tesDBRepo) GetUserByID(id int) (models.User, error) {
	var u models.User

	return u, nil
}

func (m *tesDBRepo) UpdateUser(u models.User) error {
	return nil
}

func (m *tesDBRepo) Autenticate(email, testPassword string) (int, string, error) {
	return 1, "", nil
}

func (m *tesDBRepo) AllReservation() ([]models.Reservation, error) {
	var reservations []models.Reservation
	return reservations, nil
}

// returns a slice of all reservations
func (m *tesDBRepo) AllNewReservation() ([]models.Reservation, error) {
	var reservations []models.Reservation

	return reservations, nil
}

// returns one reservation by ID
func (m *tesDBRepo) GetReservationByID(id int) (models.Reservation, error) {
	var res models.Reservation
	return res, nil
}

// UpdateReservation updates a reservation in the database
func (m *tesDBRepo) UpdateReservation(u models.Reservation) error {
	return nil
}

// deletes one reservation by id
func (m *tesDBRepo) DeleteReservation(id int) error {
	return nil
}

// updates a processed for a reservation by id
func (m *tesDBRepo) UpdateProcessedForReservation(id, processed int) error {
	return nil
}

func (m *tesDBRepo) AllRooms() ([]models.Room, error) {
	var rooms []models.Room

	return rooms, nil
}

// returns restrictions for a room by date range
func (m *tesDBRepo) GetRestrictionForRoomByDate(roomID int, start, end time.Time) ([]models.RoomRestriction, error) {
	var restrictions []models.RoomRestriction

	return restrictions, nil
}

// insert a room restrictions
func (m *tesDBRepo) InsertBlockForRoom(id int, startDate time.Time) error {
	return nil
}

// deletes a room restrictions
func (m *tesDBRepo) DeleteBlockByID(id int) error {
	return nil
}
