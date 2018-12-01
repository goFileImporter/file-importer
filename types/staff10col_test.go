package types

import (
	"testing"
)

func TestNewStaffManager(t *testing.T) {
	var staffManager Manager
	staffManager = NewStaffManager()

	_, err := staffManager.(Manager)

	if err {
		t.Fail()
	}
}
