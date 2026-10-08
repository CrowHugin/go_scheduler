package main

import (
		"fmt"
		"errors"
		"os"
		"io"
)

type givenInfos struct {
	user string
	mode string
	node int
}

func main() {
	fileName := "test.txt"
	file, err := os.Open(fileName)

	if err != nil {
		fmt.Printf("Cannot open file:\n%v", err)
		return
	}
	defer file.Close()

	bytes, err := ReadFile(file)
	if err != nil {
		fmt.Println("Error while reading file")
		return
	}
	fmt.Println(string(bytes))
	// test, err := takeInfos()
	// if err != nil {
	// 	fmt.Printf("ERROR:\n%v", err)
	// 	return
	// }
	// fmt.Println(test.user)
	// fmt.Println(test.mode)
	// fmt.Println(test.node)
}

func ReadFile(file *os.File) ([]byte, error) {
  buffer := make([]byte, 0, 512)
  for {
    if len(buffer) == cap(buffer) {
      buffer = append(buffer, 0)[:len(buffer)]
    }
    offset, err := file.Read(buffer[len(buffer):cap(buffer)])
    buffer = buffer[:len(buffer)+offset]
    if err != nil {
      if errors.Is(err, io.EOF) {
        err = nil
      }
      return buffer, err
    }
  }
}
// func takeInfos() (givenInfos, error){
// 	fileError := errors.New("cannot create struct")
// 	test := givenInfos {
// 		user: "",
// 		mode: "string",
// 		node: 18,
// 	}
// 	if test.user == ""{
// 		return givenInfos{}, fileError
// 	}
// 	return test, nil
// }
