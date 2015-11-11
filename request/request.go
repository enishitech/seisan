package request

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type Request struct {
	Data []byte
}

func (req *Request) Unmarshal(v interface{}) error {
	return yaml.Unmarshal(req.Data, v)
}

func LoadDir(dir string) ([]Request, error) {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	requests := []Request{}

	for _, entry := range entries {
		path := filepath.Join(dir, entry.Name())
		data, err := ioutil.ReadFile(path)
		if err != nil {
			return nil, err
		}

		request := Request{Data: data}
		requests = append(requests, request)
	}
	fmt.Printf("Loaded %d files\n", len(entries))

	return requests, nil
}
