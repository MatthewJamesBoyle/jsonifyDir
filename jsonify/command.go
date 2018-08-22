package jsonify

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func Do() {
	directory := os.Args[1:]
	fmt.Println(directory)
	if len(directory) < 1 {
		fmt.Println("fail")
	}

	files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		fmt.Println(f.Name())
	}
}
