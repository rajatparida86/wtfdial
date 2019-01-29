package bolt

import (
	"time"

	"github.com/boltdb/bolt"
	"github.com/rajatparida86/wtfdial"
)

type Client struct {
	Authenticator wtf.Authenticator
	Path          string
	db            *bolt.DB
	Now           func() time.Time
}

func (c *Client) Open() error {
	// Open connectin to BoltDB
	db, err := bolt.Open(c.Path, 0666, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}
	c.db = db
	// Create the bucket if not exists within a transaction
	tx, err := c.db.Begin(true)
	if err != nil {
		return err
	}
	// Rollback in case of premature errors
	defer tx.Rollback()

	if _, err = tx.CreateBucketIfNotExists([]byte("Dials")); err != nil {
		return err
	}
	return tx.Commit()
}

// Create a session
func (c *Client) Connect() *Session {
	session := newSession(c.db)
	session.authenticator = c.Authenticator
	session.now = c.Now()
	return session
}
