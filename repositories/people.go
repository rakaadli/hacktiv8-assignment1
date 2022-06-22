package repositories

import (
	"hacktiv8-assesment1/models"
	"hacktiv8-assesment1/params"
)

func SetPerson(req *params.SetPerson) models.Person {
	var people models.Person

	people.No = req.No
	people.Name = req.Name
	people.Address = req.Address
	people.Job = req.Job
	people.Reason = req.Reason

	return people

}
