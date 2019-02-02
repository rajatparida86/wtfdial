package wtf

import "time"

type DialID int
type Dial struct {
	ID           DialID    `json:"Id"`
	Token        string    `json:"-"`
	Status       float64   `json:"status"`
	ModifiedTime time.Time `json:"lastModified"`
	Name         string
}

type DialService interface {
	CreateDial(dial *Dial) error
	GetDial(dialID DialID) (*Dial, error)
	SetStatus(status float64, token string, dialID DialID) error
}

type Client interface {
	DialService() DialService
}
