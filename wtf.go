package wtf

import "time"

type UserID int
type User struct {
	ID       UserID `json:"id"`
	UserName string `json:"username"`
}

type DialID int
type Dial struct {
	ID           DialID    `json:"Id"`
	UserID       UserID    `json:"userId"`
	Status       float64   `json:"status"`
	ModifiedTime time.Time `json:"lastModified"`
	Name         string
}

type Authenticator interface {
	Authenticate(token string) (*User, error)
}

type DialService interface {
	CreateDial(dial *Dial) error
	GetDial(dialID int) (*Dial, error)
	SetStatus(status int, dialID int) error
}

type Client interface {
	Connect() Session
}

type Session interface {
	DialService() DialService
	SetAuthToken(token string)
}
