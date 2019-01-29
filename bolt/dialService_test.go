package bolt_test

import (
	"github.com/rajatparida86/wtfdial"
	"reflect"
	"testing"
)

// TestDialService_SetStatus ... Test Dial creation
func TestDialService_CreateDial(t *testing.T) {
	client := OpenClient()
	defer client.Close()
	// Create a session and get the dialService
	ds := client.Connect().DialService()

	dial := wtf.Dial{
		Name:   "Test Dial",
		Status: 50,
	}

	if err := ds.CreateDial(&dial); err != nil {
		t.Fatal(err)
	} else if dial.ID != 1 {
		t.Fatalf("Expected Id to be: %d, Got: %d", 1, dial.ID)
	} else if dial.UserID != 100 {
		t.Fatalf("Expected Dials owner's ID to be: %d, Got: %d", 100, dial.UserID)
	}
	retrievedDial, err := ds.GetDial(1)
	if err != nil {
		t.Fatal(err)
	} else if !reflect.DeepEqual(retrievedDial, &dial) {
		t.Fatalf("Expected dial:%#v , Got dial: %#v", &dial, retrievedDial)
	}
}

// TestDialService_SetStatus ... Test Dial status update
func TestDialService_SetStatus(t *testing.T) {
	client := OpenClient()
	defer client.Close()
	ds := client.Connect().DialService()

	// Define test cases
	testcases := []struct {
		dial     wtf.Dial
		expected float64
	}{
		{wtf.Dial{Name: "Dial1", Status: 50}, 20},
		{wtf.Dial{Name: "Dial2", Status: 45}, 10},
	}
	// Test all test cases
	for _, test := range testcases {
		// Create test dial
		if err := ds.CreateDial(&test.dial); err != nil {
			t.Fatal(err)
		} else //Update Dial status
		if err := ds.SetStatus(test.expected, test.dial.ID); err != nil {
			t.Fatal(err)
		} else // Retrieve updated dial
		if updated, err := ds.GetDial(test.dial.ID); err != nil {
			t.Fatal(err)
		} else // Match retreived dial status with expected status
		if updated.Status != test.expected {
			t.Fatalf("Expected status: %v, Got: %v", test.expected, updated.Status)
		}
	}
}

// TestDialService_SetStatus_ErrUnauthorized ...  Update Dial status with unauthorised user
func TestDialService_SetStatus_ErrUnauthorized(t *testing.T) {
	client := OpenClient()
	defer client.Close()
	ds := client.Connect().DialService()

	// Define test cases
	testcases := []struct {
		dial   wtf.Dial
		userID int
		err    error
	}{
		{wtf.Dial{Name: "Dial1", Status: 50}, 23, wtf.NewAuthenticationError("Dial doesn't belong to the user.")},
	}

	for _, test := range testcases {
		if err := ds.CreateDial(&test.dial); err != nil {
			t.Fatal(err)
		}
		client.Authenticator.AuthenticateFn = func(token string) (*wtf.User, error) {
			return &wtf.User{wtf.UserID(test.userID), "test"}, nil
		}
		unAuthorizedSession := client.Connect()
		err := unAuthorizedSession.DialService().SetStatus(1, test.dial.ID)

		if _, ok := err.(*wtf.AuthenticationError); !ok {
			t.Fatalf("Expected: %t, Got: %t", test.err, err)
		}
	}
}
