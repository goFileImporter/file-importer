package types

import (
	"github.com/yunabe/easycsv"
	"io"
	"reflect"
)

// Staff - this is the struct for a staff
type Staff struct {
	FirstName    string     `index:"0" json:"firstName"`
	LastName     string     `index:"1" json:"lastName"`
	Email        StaffEmail `index:"2" json:"email"`
	Level        string     `index:"3" json:"level"`
	Username     string     `index:"4" json:"username"`
	Password     string     `index:"5" json:"-"`
	SPN          string     `index:"6" json:"spn"`
	BuildingCode string     `index:"7" json:"buildingCode"`
	BuildingName string     `index:"8" json:"buildingName"`
	Role         string     `index:"9" json:"role"`
}
type StaffEmail string

func (se StaffEmail) Valid() bool {
	return true
}

// StaffManager - this will house the configuration and the methods for working with a staff of many staff
type StaffManager struct {
	data   []Data
	header bool
	reader io.Reader
	errors []error
}

func (s Staff) Valid() bool {
	// Go through validation process here
	if !s.Email.Valid() {
		return false
	}
	return true
}

func (sm StaffManager) ValidateCollection(staff []Data) error {
	// Loop over
	return nil
}

// NewStaffManager - Constructor method for StaffManager
func NewStaffManager() *StaffManager {
	return &StaffManager{
		header: true,
	}
}

func (sm *StaffManager) LoadDataFromPath(filePath string) ([]Data, error) {
	var rows []Data

	r := easycsv.NewReaderFile(filePath,
		easycsv.Option{
			TypeDecoders: map[reflect.Type]interface{}{
				reflect.TypeOf((StaffEmail)("")): func(s string) (StaffEmail, error) {
					return StaffEmail(s), nil
				},
			},
		},
	)
	err := r.Loop(func(row Staff) error {
		rows = append(rows, row)
		return nil
	})
	if err != nil {
		//We could do something but better to let it trickle up
		return rows, err
	}
	err = sm.ValidateCollection(rows)
    sm.data = rows;
	if err != nil {
		// Do something
	}
	return rows, err
}

// ShowData - return data structure
func (sm StaffManager) ShowData() []Data {
	return sm.data
}
