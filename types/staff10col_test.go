package types

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewStaffManager(t *testing.T) {
	var staffManager Manager
	staffManager = NewStaffManager()

	_, ok := staffManager.(Manager)

	if !ok {
		t.Fail()
	}
}

func TestStaffManagerGetData(t *testing.T) {
	staffManager := NewStaffManager()
	var data []Data
	var err error
	data, err = staffManager.GetData("../testdata/staffManager.csv")
	if err != nil {
		t.Fail()
	} else {
		assert.Equal(t, (data[1]).(Staff).FirstName, "John")
	}

}
