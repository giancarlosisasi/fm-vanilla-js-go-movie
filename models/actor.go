package models

type Actor struct {
	ID        int     `json:"id"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	ImageURL  *string `json:"image_url"` // we are making it a pointer so we can know when its nil (nulleable)
}
