package dto

type RatingResponse struct {
	ID       uint64  `json:"id,omitempty"`
	Username *string `json:"username,omitempty"`
	Stars    int     `json:"stars"`
}

type UpdateRatingRequest struct {
	StarsDiff int `uri:"stars_diff" binding:"required,min=1,max=100,ne=0""`
}
