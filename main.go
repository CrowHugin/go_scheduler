package main

import (
	"fmt"
	// "errors"

	"go_scheduler/virt"
)

type givenInfos struct {
	user string
	mode string
	node int
}

func main() {
	err := virt.LaunchSimu()
	if err != nil {
		fmt.Printf("Error:\n%v", err)
	}
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
