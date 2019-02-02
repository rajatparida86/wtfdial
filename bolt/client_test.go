package bolt_test

import (
	"io/ioutil"
	"os"
	"time"

	"github.com/rajatparida86/wtfdial/bolt"
)

// Client ... Wrapper for bolt.Client
type Client struct {
	*bolt.Client
}

// OpenClient ... Creates a new Client wrapper and opens the connection before each test run
func OpenClient() *Client {
	client := NewClient()
	if err := client.Client.Open(); err != nil {
		panic(err)
	}
	return client
}

// NewClient ... Wires all the necessary fields for the Client
func NewClient() *Client {
	file, err := ioutil.TempFile("", "wtf-bolt-client-")
	if err != nil {
		panic(err)
	}
	file.Close()

	client := &Client{
		Client: bolt.NewClient(),
	}
	// Mocking
	client.Path = file.Name()
	client.Now = func() time.Time { return Now }

	return client
}

// Mocked time of testing
var Now = time.Date(2019, time.January, 29, 11, 47, 25, 23, time.UTC)

// Close ... Closes a connection to DB after test run
func (c *Client) Close() error {
	defer os.Remove(c.Path)
	return c.Client.Close()
}
