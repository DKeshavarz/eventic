package entity

type Organization struct {
	ID          int     `json:"organizer_id"`
	OwnerID     int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	LogoPic     *string `json:"logo_pic"`
	Email       *string `json:"email"`
	Phone       *string `json:"phone"`
}
