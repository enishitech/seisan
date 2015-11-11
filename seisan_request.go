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

func loadSeisanRequest(srcPath string) (*SeisanRequest, error) {
	buf, err := ioutil.ReadFile(srcPath)
	if err != nil {
		return nil, err
	}
	var req SeisanRequest
	err = yaml.Unmarshal(buf, &req)
	if err != nil {
		return nil, err
	}
	return &req, nil
}

func loadSeisanRequests(targetPath string) []SeisanRequest {
	entries, err := ioutil.ReadDir(targetPath)
	if err != nil {
		log.Fatal("ERROR: ", err.Error())
	}
	requests := []SeisanRequest{}

	for _, entry := range entries {
		path := filepath.Join(targetPath, entry.Name())
		request, err := loadSeisanRequest(path)
		if err != nil {
			log.Fatal("ERROR: ", err.Error())
		}
		requests = append(requests, *request)
	}
	fmt.Printf("Loaded %d files\n", len(entries))

	return requests
}
