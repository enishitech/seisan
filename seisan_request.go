package main

import (
	"fmt"
	"io/ioutil"
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

func loadSeisanRequests(targetPath string) ([]SeisanRequest, error) {
	entries, err := ioutil.ReadDir(targetPath)
	if err != nil {
		return nil, err
	}
	requests := []SeisanRequest{}

	for _, entry := range entries {
		path := filepath.Join(targetPath, entry.Name())
		request, err := loadSeisanRequest(path)
		if err != nil {
			return nil, err
		}
		requests = append(requests, *request)
	}
	fmt.Printf("Loaded %d files\n", len(entries))

	return requests, nil
}
