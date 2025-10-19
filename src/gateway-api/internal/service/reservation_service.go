package service

import (
	"fmt"
	"gateway-api/internal/client"
	"gateway-api/internal/dto"
)

type ReservationService struct {
	ClientRes  *client.Reservation
	ClientLib  *client.Library
	ClientRate *client.Rating
}

func NewReservationService(clRes *client.Reservation, clLib *client.Library, clRate *client.Rating) *ReservationService {
	return &ReservationService{ClientRes: clRes, ClientLib: clLib, ClientRate: clRate}
}

func (s *ReservationService) Get(username string) ([]dto.ReservationFullResponse, error) {
	raw, err := s.ClientRes.Get(username)
	if err != nil {
		return nil, err
	}

	result := make([]dto.ReservationFullResponse, 0, len(raw))
	for _, r := range raw {
		book, err := s.ClientLib.GetBookByUID(r.BookUID)
		if err != nil {
			return nil, err
		}

		lib, err := s.ClientLib.GetLibraryByUID(r.LibraryUID)
		if err != nil {
			return nil, err
		}

		fullRes := dto.ReservationToFull(r, dto.BookToRaw(*book), *lib)
		result = append(result, fullRes)
	}
	return result, nil
}

func (s *ReservationService) CreateReservation(username string, req dto.CreateReservationRequest) (*dto.ReservationFullResponse, error) {

	resCount, err := s.ClientRes.GetCurrentAmount(username)
	if err != nil {
		return nil, fmt.Errorf("failed to get current amount: %s", err)
	}
	starsCount, err := s.ClientRate.Get(username)
	if err != nil {
		return nil, fmt.Errorf("failed to get rating: %s", err)
	}
	if resCount >= starsCount.Stars {
		return nil, fmt.Errorf("You rented maximum amount of books", resCount)
	}
	result, err := s.ClientRes.Create(username, req)
	if err != nil {
		return nil, fmt.Errorf("failed to create reservation: %s", err)
	}
	err = s.ClientLib.UpdateBookCount(result.LibraryUID, result.BookUID, -1)
	if err != nil {
		return nil, fmt.Errorf("failed to update book count: %s", err)
	}

	book, err := s.ClientLib.GetBookByUID(result.BookUID)
	if err != nil {
		return nil, err
	}
	lib, err := s.ClientLib.GetLibraryByUID(result.LibraryUID)
	if err != nil {
		return nil, err
	}
	fullRes := dto.ReservationToFull(*result, dto.BookToRaw(*book), *lib)

	return &fullRes, nil
}
