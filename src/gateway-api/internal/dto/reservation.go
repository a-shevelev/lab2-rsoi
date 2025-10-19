package dto

type CreateReservationRequest struct {
	BookUID    string `json:"book_uid" binding:"required"`
	LibraryUID string `json:"library_uid" binding:"required"`
	TillDate   string `json:"till_date" binding:"required,datetime=2006-01-02"`
}

type ReservationResponse struct {
	ReservationUID string `json:"reservationUid"`
	Username       string `json:"username"`
	BookUID        string `json:"bookUid"`
	LibraryUID     string `json:"libraryUid"`
	Status         string `json:"status"`
	StartDate      string `json:"startDate"`
	TillDate       string `json:"tillDate"`
}

type ReservationFullResponse struct {
	ReservationUID string          `json:"reservationUid"`
	Username       string          `json:"username"`
	Book           BookResponseRaw `json:"book"`
	Library        LibraryResponse `json:"library"`
	Status         string          `json:"status"`
	StartDate      string          `json:"startDate"`
	TillDate       string          `json:"tillDate"`
}

func ReservationToFull(
	r ReservationResponse,
	book BookResponseRaw,
	library LibraryResponse,
) ReservationFullResponse {
	return ReservationFullResponse{
		ReservationUID: r.ReservationUID,
		Username:       r.Username,
		Status:         r.Status,
		StartDate:      r.StartDate,
		TillDate:       r.TillDate,
		Book:           book,
		Library:        library,
	}
}
