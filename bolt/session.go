package bolt

import (
	"time"

	"github.com/boltdb/bolt"
	wtf "github.com/rajatparida86/wtfdial"
)

type Session struct {
	db            *bolt.DB
	now           time.Time
	user          *wtf.User
	authenticator wtf.Authenticator
	token         string
	dialService   DialService
}

func newSession(db *bolt.DB) *Session {
	s := &Session{db: db}
	s.dialService.session = s
	return s
}

func (s *Session) SetAuthToken(token string) {
	s.token = token
}

func (s *Session) Authenticate() (*wtf.User, error) {
	if s.user != nil {
		return s.user, nil
	}
	u, err := s.authenticator.Authenticate(s.token)
	if err != nil {
		return nil, err
	}
	s.user = u
	return u, nil
}

func (s *Session) DialService() *DialService { return &s.dialService }
