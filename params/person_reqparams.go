package params

type SetPerson struct {
	No      int
	Name    string
	Address string
	Job     string
	Reason  string
}

func SetNewPerson(no int, name string, address string, job string, reason string) *SetPerson {
	return &SetPerson{
		No:      no,
		Name:    name,
		Address: address,
		Job:     job,
		Reason:  reason,
	}
}
