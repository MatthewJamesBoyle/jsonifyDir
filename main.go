package main

import (
	"errors"
	"fmt"
	"github.com/mattJamesBoyle/jsonifyDir/jsonify"
	"os"
)

func main() {
	directory := os.Args[1:]
	if len(directory) < 1 {
		fmt.Println("fail")
		panic(errors.New("No directory passed in."))
	}
	result, err := jsonify.Do(directory[0])
	if err != nil {
		panic(err)
	}
	jsonify.Print(result)
}
