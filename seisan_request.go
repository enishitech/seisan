package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type SeisanRequest struct {
	Applicant string
	Expenses  []Expense `yaml:"expense"`
}

func loadSeisanRequests(targetPath string) []SeisanRequest {
	entries, err := ioutil.ReadDir(targetPath)
	if err != nil {
		log.Fatal("ERROR: ", err.Error())
	}
	requests := []SeisanRequest{}

	for _, entry := range entries {
		entryPath := filepath.Join(targetPath, entry.Name())
		buf, err := ioutil.ReadFile(entryPath)
		if err != nil {
			log.Fatal("ERROR: ", err.Error())
		}
		var req SeisanRequest
		err = yaml.Unmarshal(buf, &req)
		if err != nil {
			log.Fatal("ERROR: ", err.Error())
		}
		requests = append(requests, req)
	}
	fmt.Printf("Loaded %d files\n", len(entries))

	return requests
}
