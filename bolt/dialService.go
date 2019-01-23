package bolt

import (
	"encoding/binary"

	wtf "github.com/rajatparida86/wtfdial"
	"github.com/rajatparida86/wtfdial/bolt/internal"
)

type DialService struct {
	session *Session
}

func (d *DialService) CreateDial(dial *wtf.Dial) error {
	// Get the user to associate the Dial to
	user, err := d.session.Authenticate()
	if err != nil {
		return err
	}

	// Start a writable transaction in boltdb
	tx, err := d.session.db.Begin(true)
	if err != nil {
		return nil
	}
	defer tx.Rollback()

	// Get the dial Bucket
	dials := tx.Bucket([]byte("Dials"))
	// Grab next dialid (to be created) from bolt bucket
	dialID, err := dials.NextSequence()
	if err != nil {
		return err
	}
	dial.ID = int(dialID)
	dial.UserID = user.ID
	dial.ModifiedTime = d.session.now

	// Marshal the Dial object and put into BoltDB bucket
	if dialproto, err := internal.MarshalDial(dial); err != nil {
		return err
	} else if err := dials.Put(itob(dial.ID), dialproto); err != nil {
		return err
	}

	return tx.Commit()
}

func (d *DialService) GetDial(dialID int) (*wtf.Dial, error) {

	return nil, nil
}

func (d *DialService) SetStatus(status int, dialId int) error {

	return nil
}

//Converts an integer to byte slice
func itob(value int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(value))
	return b
}
