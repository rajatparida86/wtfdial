package mock

import (
	"github.com/rajatparida86/wtfdial"
)

// Authenticator represents a service for authenticating users.
type Authenticator struct {
	AuthenticateFn      func(token string) (*wtf.User, error)
	AuthenticateInvoked bool
}

func (a *Authenticator) Authenticate(token string) (*wtf.User, error) {
	a.AuthenticateInvoked = true
	return a.AuthenticateFn(token)
}

// DefaultUser is the user authenticated by DefaultAuthenticator.
var DefaultUser = &wtf.User{ID: 100}

// DefaultAuthenticator returns an authenticator that returns the default user.
func DefaultAuthenticator() Authenticator {
	return Authenticator{
		AuthenticateFn: func(token string) (*wtf.User, error) { return DefaultUser, nil },
	}
}

type DialService struct {
	DialFn      func(id wtf.DialID) (*wtf.Dial, error)
	DialInvoked bool

	CreateDialFn      func(dial *wtf.Dial) error
	CreateDialInvoked bool

	SetLevelFn      func(id wtf.DialID, level float64) error
	SetLevelInvoked bool
}

func (s *DialService) Dial(id wtf.DialID) (*wtf.Dial, error) {
	s.DialInvoked = true
	return s.DialFn(id)
}

func (s *DialService) CreateDial(dial *wtf.Dial) error {
	s.CreateDialInvoked = true
	return s.CreateDialFn(dial)
}

func (s *DialService) SetLevel(id wtf.DialID, level float64) error {
	s.SetLevelInvoked = true
	return s.SetLevelFn(id, level)
}
