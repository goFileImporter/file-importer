package types

import (
	"fmt"
	"github.com/bxcodec/faker"
	"github.com/stretchr/testify/suite"
	"testing"
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

func (suite *StaffTestSuite) SetupTest() {
	err := faker.FakeData(&suite.staff)
	if err != nil {
		panic(err)
	}
}

func (suite *StaffTestSuite) TestStaffStructValid() {
	suite.True(suite.staff.Valid())
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

	for i := 0; i < 10; i++ {
		var staff Staff
		err := faker.FakeData(&staff)
		if err != nil {
			panic(err)
		}
		suite.staffManager.SetData(append(suite.staffManager.ShowData(), staff))
	}
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

func (suite *StaffManagerSuite) TestShowData() {
	suite.NotEmpty(suite.staffManager.ShowData())
}

func TestStaffManagerSuite(t *testing.T) {
	suite.Run(t, new(StaffManagerSuite))
}
