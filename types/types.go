package types

// Manager - interface for Managers to follow
type Manager interface {
	SetData([]Data)
	LoadDataFromPath(string) ([]Data, error)
	ShowData() []Data
	ManagerValidator
}

// Data - is the underlying data struct
type Data interface {
	Valid() []error
}

type ErroredRecord struct {
	err []error
	Data
}

// ManagerValidator - This is responsible for the thing that validates the data in the manager
type ManagerValidator interface {
	ValidateCollection() []ErroredRecord
}

const staff10ColType string = "staff10col"

// NewManager - return a Manager
func NewManager(t string) Manager {
	var m Manager

	if t == staff10ColType {
		m = NewStaffManager()
	}

	return m
}

// NewValidator - return validator
// func NewValidator(t string) {
// 	var mv ManagerValidator
//
// 	if t == staff10ColType {
// 		mv = NewStaffValidator()
// 	}
//
// 	return mv
// }
