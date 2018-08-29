package jsonify

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type response struct {
	Folders []folder `json:"folders"`
}

type folder struct {
	Name string `json:"name"`
	Size int64  `json:"size"`
}

func Do(dir string) (*response, error) {

	var folders []folder
	files, err := ioutil.ReadDir("./" + dir)
	if err != nil {
		return nil, err
	}

	for _, f := range files {
		if f.IsDir() {
			folders = append(folders, folder{
				Name: f.Name(),
				Size: f.Size(),
			})
		}

	}
	return &response{
		Folders: folders,
	}, nil
}

func Print(res *response) {
	b, err := json.Marshal(res)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(b))
}
