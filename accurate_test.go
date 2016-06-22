package accurate

import "testing"

var (
	testClientID     = "your_clientId_here"
	testClientSecret = "your_clientSecret_here"
)

func TestSetClientID(t *testing.T) {
	SetClientID(testClientID)
	if clientID != testClientID {
		t.Errorf("expected %q got %q", testClientID, clientID)
	}
}

func TestSetClientSecret(t *testing.T) {
	SetClientID(testClientID)
	if clientID != testClientID {
		t.Errorf("expected %q got %q", testClientID, clientID)
	}
}
