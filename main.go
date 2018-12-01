package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/yunabe/easycsv"
	"log"
)

// Staff - this is the struct for a staff
type Staff struct {
	FirstName    string `index:"0" json:"firstName"`
	LastName     string `index:"1" json:"lastName"`
	Email        string `index:"2" json:"email"`
	Level        string `index:"3" json:"level"`
	Username     string `index:"4" json:"username"`
	Password     string `index:"5" json:"-"`
	SPN          string `index:"6" json:"spn"`
	BuildingCode string `index:"7" json:"buildingCode"`
	BuildingName string `index:"8" json:"buildingName"`
	Role         string `index:"9" json:"role"`
}

var filePath string

func main() {
	flag.StringVar(&filePath, "file", "staff.csv", "Choose path for file")
	flag.Parse()
	var rows []Staff
	r := easycsv.NewReaderFile(filePath)
	err := r.Loop(func(row Staff) error {
		rows = append(rows, row)
		return nil
	})
	if err != nil {
		log.Fatalf("Failed to read a CSV file: %v", err)
	}

	s, _ := json.Marshal(rows)
	str := fmt.Sprintf("%s", s)
	fmt.Println(str)

}
