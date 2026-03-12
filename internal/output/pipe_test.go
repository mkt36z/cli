package output

import "testing"

func TestNewFormatter(t *testing.T) {
	f := NewFormatter(true)
	if !f.IsJSON() {
		t.Error("NewFormatter(true) should be JSON")
	}

	f = NewFormatter(false)
	if f.IsJSON() {
		t.Error("NewFormatter(false) should not be JSON")
	}
}
