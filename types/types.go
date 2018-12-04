package types

// Manager - interface for Managers to follow
type Manager interface {
	LoadDataFromPath(string) ([]Data, error)
	ShowData() []Data
	// SetValidator(ManagerValidator) Manager
}

// Data - is the underlying data struct
type Data interface {
	Valid() bool
}

// ManagerValidator - This is responsible for the thing that validates the data in the manager
type ManagerValidator interface {
	ValidateCollection([]Data) error
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
