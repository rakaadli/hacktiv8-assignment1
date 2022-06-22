package main

import (
	"fmt"
	"hacktiv8-assesment1/models"
	"hacktiv8-assesment1/params"
	"hacktiv8-assesment1/repositories"
	"os"
	"strconv"
)

func main() {
	var people []models.Person

	if len(os.Args) > 1 {

		RequestedNumber, err := strconv.Atoi(os.Args[1])

		if err != nil {
			fmt.Println("Please make sure you set the number not with string! example: `go run biodata.go 1`")
			return
		}

		person1 := params.SetNewPerson(1, "Raka Adli Pramudita", "Jakarta Timur", "Back-End Engineer", "Improve skill for coding")
		person2 := params.SetNewPerson(2, "The Strongest Man", "Jakarta Barat", "Devops", "learning go routine")
		person3 := params.SetNewPerson(3, "Bejo", "Jakarta Pusat", "Platform", "to elaborate go with platform")

		people = append(people, repositories.SetPerson(person1))
		people = append(people, repositories.SetPerson(person2))
		people = append(people, repositories.SetPerson(person3))

		for _, person := range people {
			person.Print(RequestedNumber)
		}
	} else {
		fmt.Println("Please set the number! example: `go run biodata.go 1`")
	}
}
