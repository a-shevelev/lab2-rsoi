package service

import (
	"gateway-api/internal/client"
	"gateway-api/internal/dto"
)

type ReservationService struct {
	ClientRes *client.Reservation
	ClientLib *client.Library
}

func NewReservationService(cr *client.Reservation, cl *client.Library) *ReservationService {
	return &ReservationService{ClientRes: cr, ClientLib: cl}
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
