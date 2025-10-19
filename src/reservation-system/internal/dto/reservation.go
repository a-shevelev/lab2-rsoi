package dto

import (
	"reservation-system/internal/models"
)

type CreateReservationRequest struct {
	BookUID    string `json:"book_uid" binding:"required"`
	LibraryUID string `json:"library_uid" binding:"required"`
	TillDate   string `json:"till_date" binding:"required,datetime=2006-01-02"`
}

type ReservationResponse struct {
	ReservationUID string `json:"reservation_uid"`
	Username       string `json:"username"`
	BookUID        string `json:"book_uid"`
	LibraryUID     string `json:"library_uid"`
	Status         string `json:"status"`
	StartDate      string `json:"start_date"`
	TillDate       string `json:"till_date"`
}

type ReservationsListResponse struct {
	Items []ReservationResponse
}

func ToReservationDTO(m *models.Reservation) ReservationResponse {
	return ReservationResponse{
		ReservationUID: m.ReservationUID.String(),
		Username:       m.Username,
		BookUID:        m.BookUID.String(),
		LibraryUID:     m.LibraryUID.String(),
		Status:         m.Status,
		StartDate:      m.StartDate.Format("2006-01-02"),
		TillDate:       m.TillDate.Format("2006-01-02"),
	}
}

func ToReservationsDTO(list []models.Reservation) []ReservationResponse {
	out := make([]ReservationResponse, 0, len(list))
	for _, r := range list {
		out = append(out, ToReservationDTO(&r))
	}
	return out
}
