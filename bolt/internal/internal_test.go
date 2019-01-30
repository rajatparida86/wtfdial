package internal_test

import (
	"reflect"
	"testing"
	"time"

	"github.com/rajatparida86/wtfdial"
	"github.com/rajatparida86/wtfdial/bolt/internal"
)

func TestMarshalDial(t *testing.T) {
	dial := &wtf.Dial{
		ID:           12,
		Status:       45,
		Name:         "test dial",
		ModifiedTime: time.Now().UTC(),
	}
	if dialBytes, err := internal.MarshalDial(dial); err != nil {
		t.Fatal(err)
	} else if unmarshalledDial, err := internal.UnMarshalDial(dialBytes); err != nil {
		t.Fatal(err)
	} else if !reflect.DeepEqual(dial, unmarshalledDial) {
		t.Fatalf("Expected: %#v, Got: %#v", dial, unmarshalledDial)
	}
}
