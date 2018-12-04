package types

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStaffValid(t *testing.T) {
	staff := Staff{
		"John",
		"Doe",
		"john@doe.com",
		"1",
		"jdoe",
		"password",
		"123",
		"1234",
		"My Building",
		"staff",
	}
	assert.True(t, staff.Valid())
}

func TestNewStaffManager(t *testing.T) {
	var staffManager Manager
	staffManager = NewStaffManager()

	_, ok := staffManager.(Manager)

	if !ok {
		t.Fail()
	}
}

func TestStaffManagerLoadDataFromPath(t *testing.T) {
	staffManager := NewStaffManager()
	var data []Data
	var err error
	data, err = staffManager.LoadDataFromPath("../testdata/staffManager.csv")
	if err != nil {
		t.Fail()
	} else {
		assert.Equal(t, (data[1]).(Staff).FirstName, "John")
	}

}
