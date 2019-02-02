package bolt

import (
	"encoding/binary"
	"fmt"

	wtf "github.com/rajatparida86/wtfdial"
	"github.com/rajatparida86/wtfdial/bolt/internal"
)

// DialService ... Provides a service to get/post/update Dial to Boltdb in a session
type DialService struct {
	client *Client
}

const (
	dials string = "Dials"
)

// CreateDial ... Creates a new dial and associates the user with it
func (d *DialService) CreateDial(dial *wtf.Dial) error {
	// Start a writable transaction in boltdb
	tx, err := d.client.db.Begin(true)
	if err != nil {
		return nil
	}
	defer tx.Rollback()

	// Chek if Dial already exists
	_, err = d.GetDial(dial.ID)
	if _, ok := err.(*wtf.ErrDialIdNotFound); !ok {
		return wtf.NewErrDialAlreadyExists("Can not create an existing Dial")
	}

	// Get the dial Bucket and the next dialID
	dials := tx.Bucket([]byte(dials))
	dialID, err := dials.NextSequence()
	if err != nil {
		return err
	}
	//Prepare the new Dial
	dial.ID = wtf.DialID(dialID)
	dial.ModifiedTime = d.client.Now()

	// Marshal the Dial object and put into BoltDB bucket
	if dialproto, err := internal.MarshalDial(dial); err != nil {
		return err
	} else if err := dials.Put(itob(int(dial.ID)), dialproto); err != nil {
		return err
	}

	return tx.Commit()
}

// GetDial ... Gets a Dial from "Dials" bucket for a given dialId
func (d *DialService) GetDial(dialID wtf.DialID) (*wtf.Dial, error) {
	// Create non writable transaction
	tx, err := d.client.db.Begin(false)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	// Read the dial from bucket and unmarshal
	dial := &wtf.Dial{}
	if dialProto := tx.Bucket([]byte(dials)).Get(itob(int(dialID))); dialProto == nil {
		return nil, wtf.NewErrDialIdNotFound(fmt.Sprintf("Dial does not exists for given DialID"))
	} else if dial, err = internal.UnMarshalDial(dialProto); err != nil {
		return nil, err
	}
	return dial, nil
}

// SetStatus ... Sets the status for a dial
func (d *DialService) SetStatus(status float64, token string, dialID wtf.DialID) error {
	// Check if requested user owns the dial
	dial, err := d.GetDial(dialID)
	if err != nil {
		return err
	} else if dial.Token != token {
		return wtf.NewAuthenticationError("Dial doesn't belong to the user.")
	}
	// Update Dial status
	dial.Status = status
	dial.ModifiedTime = d.client.Now()
	dialProto, err := internal.MarshalDial(dial)
	if err != nil {
		return err
	}
	//Get writable transaction
	tx, err := d.client.db.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	bucket := tx.Bucket([]byte(dials))
	if err := bucket.Put(itob(int(dialID)), dialProto); err != nil {
		return err
	}

	return tx.Commit()
}

//Converts an integer to byte slice
func itob(value int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(value))
	return b
}
