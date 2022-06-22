package models

import "fmt"

type Person struct {
	No      int
	Name    string
	Address string
	Job     string
	Reason  string
}

func (p *Person) Print(numb int) {
	if numb == p.No {
		fmt.Println("No \t\t : ", p.No)
		fmt.Println("Name \t\t : ", p.Name)
		fmt.Println("Name \t\t : ", p.Address)
		fmt.Println("Job Position \t : ", p.Job)
		fmt.Println("Reason \t\t : ", p.Reason)
	}
}
