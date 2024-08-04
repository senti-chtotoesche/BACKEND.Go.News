package News

import "time"

type News struct {
	ID              int       `json:"id"`
	Title           string    `json:"title"`
	Description     string    `json:"description"`
	CategoryID      int       `json:"category_id"`
	NDate           time.Time `json:"ndate"`
	FullDescription string    `json:"full_description"`
}
