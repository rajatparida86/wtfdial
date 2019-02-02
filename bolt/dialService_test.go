package bolt_test

import (
	"testing"

	"github.com/rajatparida86/wtfdial"
	"github.com/stretchr/testify/assert"
)

// TestDialService_SetStatus ... Test Dial creation
func TestDialService_CreateDial(t *testing.T) {
	testCases := []struct {
		dial          wtf.Dial
		expectedError error
	}{
		{wtf.Dial{Name: "Dial1", Status: 50, Token: "xxx"}, nil},
		{wtf.Dial{ID: 1, Name: "Dial2", Status: 45, Token: "xxx"}, &wtf.ErrDialAlreadyExists{}},
	}

	client := OpenClient()
	defer client.Close()
	// Create a session and get the dialService
	ds := client.DialService()

	for _, test := range testCases {
		if err := ds.CreateDial(&test.dial); err != nil {
			assert.IsType(t, test.expectedError, err)
		}
	}
}

// TestDialService_CreateDial_Dial_Exists -- Test Existing dial cannot be created
func TestDialService_CreateDial_Dial_Exists(t *testing.T) {
	client := OpenClient()
	defer client.Close()
	ds := client.DialService()
	dial := wtf.Dial{
		Token:  "xxx",
		Name:   "Test Dial",
		Status: 50,
	}
	if err := ds.CreateDial(&dial); err != nil {
		t.Fatal(err)
	}
	if err := ds.CreateDial(&dial); err == nil {
		t.Fatal(err)
	} else if _, ok := err.(*wtf.ErrDialAlreadyExists); !ok {
		t.Fatalf("Expected ErrDialAlreadyExists. Got error:%t", err)
	}
}

// TestDialService_GetDial-- Test get dial
func TestDialService_GetDial(t *testing.T) {
	// Define test cases
	testcases := []struct {
		dial             wtf.Dial
		dialID           wtf.DialID
		expectedError    error
		expectedDialName string
	}{
		{wtf.Dial{Name: "Dial1", Status: 50, Token: "xxx"}, 1, nil, "Dial1"},
		{wtf.Dial{Name: "Dial2", Status: 45, Token: "xxx"}, 10, &wtf.ErrDialIdNotFound{}, ""},
	}

	client := OpenClient()
	defer client.Close()
	ds := client.DialService()

	for _, test := range testcases {
		if err := ds.CreateDial(&test.dial); err != nil {
			t.Fatal(err)
		}
		result, err := ds.GetDial(test.dialID)
		assert.IsType(t, test.expectedError, err)
		if err == nil {
			assert.Equal(t, test.expectedDialName, result.Name)
		}
	}
}

// TestDialService_SetStatus ... Test Dial status update
func TestDialService_SetStatus(t *testing.T) {
	// Define test cases
	testcases := []struct {
		dial           wtf.Dial
		token          string
		expectedStatus float64
		expectedError  error
	}{
		{wtf.Dial{Name: "Dial1", Status: 50, Token: "xxx"}, "xxx", 20, nil},
		{wtf.Dial{Name: "Dial2", Status: 45, Token: "yyy"}, "yyy", 10, nil},
		{wtf.Dial{Name: "Dial3", Status: 18, Token: "zzz"}, "xxx", 10, wtf.NewAuthenticationError("")},
	}

	client := OpenClient()
	defer client.Close()
	ds := client.DialService()

	// Test all test cases
	for _, test := range testcases {
		// Create test dial
		if err := ds.CreateDial(&test.dial); err != nil {
			t.Fatal(err)
		} else //Update Dial status
		if err := ds.SetStatus(test.expectedStatus, test.token, test.dial.ID); err != nil {
			assert.IsType(t, test.expectedError, err)
		} else // Retrieve updated dial
		if updated, err := ds.GetDial(test.dial.ID); err != nil {
			t.Fatal(err)
		} else {
			// Match retreived dial status with expected status
			assert.Equal(t, test.expectedStatus, updated.Status)
		}
	}
}
