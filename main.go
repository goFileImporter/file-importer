package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/rihoj/file-importer/types"
	"log"
)

var (
	filePath, fileType string
)

func main() {
	flag.StringVar(&filePath, "file", "staff.csv", "Choose path for file")
	flag.StringVar(&fileType, "type", "", "Choose type of file to process")
	flag.Parse()
	var rows []types.Data
	manager := types.NewManager(fileType)

	rows, err := manager.GetData(filePath)

	if err != nil {
		log.Fatal(err)
	}
	manager = manager.SetData(rows)

	validator, ok := manager.(types.ManagerValidator)

	if ok {
		_ = validator.ValidateCollection(manager.ShowData())
	}

	s, _ := json.Marshal(manager.ShowData())
	str := fmt.Sprintf("%s", s)
	fmt.Println(str)

}
