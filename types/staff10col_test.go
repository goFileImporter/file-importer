package types

import (
	"fmt"
	"testing"

	"github.com/bxcodec/faker"
	"github.com/stretchr/testify/suite"
)

type StaffTestSuite struct {
	suite.Suite
	staff Staff
}

type StaffManagerSuite struct {
	suite.Suite
	staffManager      Manager
	staffManagerFaker Manager
}

type StaffManagerSuiteWithErrs struct {
	suite.Suite
	staffManager      Manager
	staffManagerFaker Manager
}

func (suite *StaffTestSuite) SetupTest() {
	err := faker.FakeData(&suite.staff)
	if err != nil {
		panic(err)
	}
}

func (suite *StaffTestSuite) TestStaffStructValid() {
	suite.Empty(suite.staff.Valid())
}

func TestStaffSuite(t *testing.T) {
	suite.Run(t, new(StaffTestSuite))
}

func (suite *StaffManagerSuite) SetupTest() {
	var staffManager Manager
	staffManager = NewStaffManager()
	_, staffManagerOk := staffManager.(Manager)
	suite.True(staffManagerOk)
	suite.staffManager = staffManager

	var staff Staff
	for i := 0; i < 10; i++ {
		err := faker.FakeData(&staff)
		if err != nil {
			panic(err)
		}
		suite.staffManager.SetData(append(suite.staffManager.ShowData(), staff))
	}
}

func (suite *StaffManagerSuiteWithErrs) SetupTest() {
	var staffManager Manager
	staffManager = NewStaffManager()
	_, staffManagerOk := staffManager.(Manager)
	suite.True(staffManagerOk)
	suite.staffManager = staffManager

	var staff Staff
	var st []Data
	for i := 0; i < 10; i++ {
		err := faker.FakeData(&staff)
		staff.Email = "bob.bob.com"
		if err != nil {
			panic(err)
		}
		st = append(st, staff)
	}

	suite.staffManager.SetData(st)
}

func (s *StaffManagerSuite) TestLoadDataFromPath() {
	var data []Data
	var err error
	data, err = s.staffManager.LoadDataFromPath("../testdata/staffManager.csv")
	if s.Nil(err) {
		s.Equal((data[1]).(Staff).FirstName, "John")
	}
}

func (suite *StaffManagerSuite) TestStaffCollectionValid() {
	errs := (suite.staffManager).ValidateCollection()
	for _, err := range errs {
		fmt.Println(err.err)
	}
	suite.Empty(errs)
}

func (suite *StaffManagerSuiteWithErrs) TestStaffCollectionNotValid() {

	errs := (suite.staffManager).ValidateCollection()
	suite.NotEmpty(errs)
}

func (s *StaffManagerSuiteWithErrs) TestLoadDataFromPath() {
	var err error
	_, err = s.staffManager.LoadDataFromPath("../failing/file/that/will/fail/do/not/put/a/file/here.csv")
	s.NotNil(err)
}

func (suite *StaffManagerSuite) TestShowData() {
	suite.NotEmpty(suite.staffManager.ShowData())
}

func TestStaffManagerSuite(t *testing.T) {
	suite.Run(t, new(StaffManagerSuite))
	suite.Run(t, new(StaffManagerSuiteWithErrs))
}
