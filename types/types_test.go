package types

import (
	"reflect"
	"testing"
)

func TestNewManagerStaff(t *testing.T) {
	var staffManager Manager
	staffManager = NewManager("staff10col")

	_, err := staffManager.(Manager)

	if err {
		t.Fail()
	}

	if reflect.TypeOf(staffManager) != reflect.TypeOf(StaffManager{}) {
		t.Fail()
	}

}
