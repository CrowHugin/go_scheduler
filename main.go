package main

import (
	"fmt"
	"go_scheduler/file"
	"os"
)

type givenInfos struct {
	user string
	mode string
	node int
}

func main() {
	fileName := "test.txt"
	file_, err := os.Open(fileName)

	if err != nil {
		fmt.Printf("Cannot open file:\n%v", err)
		return
	}
	defer file_.Close()

	bytes, err := file.ReadFile(file_)
	if err != nil {
		fmt.Println("Error while reading file")
		return
	}
	err = file.CheckContent(string(bytes))
	if err != nil {
		fmt.Printf("Error:\n%v", err)
	}
	// fmt.Println(string(bytes))
	// test, err := takeInfos()
	// if err != nil {
	// 	fmt.Printf("ERROR:\n%v", err)
	// 	return
	// }
	// fmt.Println(test.user)
	// fmt.Println(test.mode)
	// fmt.Println(test.node)
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
