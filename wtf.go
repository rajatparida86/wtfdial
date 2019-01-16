package wtf

import "time"

type User struct {
	ID       int `json:"id"`
	UserName int `json:"username"`
}

type Dial struct {
	ID           int       `json:"Id"`
	UserID       int       `json:"userId"`
	status       float64   `json:"status"`
	ModifiedTime time.Time `json:"lastModified"`
}

type UserService interface {
	Authenticate(token string) *User
}

type DialService interface {
	CreateDial(dial *Dial) error
	GetDial(dialID int) (*Dial, error)
	SetStatus(status int, dialID int) error
}