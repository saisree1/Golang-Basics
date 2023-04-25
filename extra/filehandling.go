package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	goPath := os.Getenv("GOPATH")
	fmt.Println("GOPATH: ", goPath)
	// filePathToRead := filepath.Join(goPath, )
	file, err := os.Open("helloworld.txt")
	if err != nil {
		fmt.Println("error while opening file : ", err)
	}
	defer file.Close()
	fileDataByteArray, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("error while reading file : ", err)
	}
	fmt.Println("File Name:", file.Name(), ", Contents:", string(fileDataByteArray))
}
