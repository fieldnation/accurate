package accurate

import "testing"

func TestAlive(t *testing.T) {
	if !Alive() {
		t.Error("something is wrong, or accurate is down")
	}
}
