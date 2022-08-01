package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	goPath := os.Getenv("GOPATH")
	fmt.Println(goPath)

	filePathToRead := filepath.Join(goPath, "/src/session/extra/helloworld.txt")
	file, err := os.Open(filePathToRead)
	if err != nil {
		fmt.Println("error while opening file : ", err)
	}
	defer file.Close()
	fileDataByteArray, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("error while reading file : ", err)
	}
	fmt.Println(string(fileDataByteArray), file.Name())
}
