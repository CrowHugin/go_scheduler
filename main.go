package main

import (
	"fmt"
	// "errors"
	"go_scheduler/simu"
)

type givenInfos struct {
	user string
	mode string
	node int
}

func main() {
	jobs, err := simu.LaunchSimu(5)
	if err != nil {
		fmt.Printf("Error:\n%v", err)
	}
	fmt.Println(jobs)
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
