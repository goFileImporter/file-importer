package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/catdevman/file-importer/types"
	"log"
)

var (
	filePath, fileType string
)

func main() {
	flag.StringVar(&filePath, "file", "staff.csv", "Choose path for file")
	flag.StringVar(&fileType, "type", "", "Choose type of file to process")
	flag.Parse()

	manager := types.NewManager(fileType)

	_, err := manager.LoadDataFromPath(filePath)

	if err != nil {
		log.Fatal(err)
	}

	validator, ok := manager.(types.ManagerValidator)

	if ok {
		_ = validator.ValidateCollection(manager.ShowData())
	}

	s, _ := json.Marshal(manager.ShowData())
	fmt.Printf("%s\n", s)

}
