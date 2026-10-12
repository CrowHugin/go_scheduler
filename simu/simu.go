package simu

import (
	"fmt"
	"errors"
	"math/rand/v2"
)

func setupJob() (string, error){
	possibleJobs := []string{"1;test;bku;16", "2;test2;lku;32", "3;test3;kulb;64"}
	if len(possibleJobs) == 0 {
		return "", errors.New("No zob in job")
	}
	idx := rand.IntN(len(possibleJobs))
	selectedJob := possibleJobs[idx]
	return selectedJob, nil
}

func LaunchSimu(nbTest int) ([]string, error) {
	var jobSlice []string
	fmt.Println("launching virtu")
	for i:=0;i<nbTest;i++{
		job, err := setupJob()
		if err != nil{
			return nil, err
		}
		jobSlice =  append(jobSlice, job)

	}
	return jobSlice, nil
}
