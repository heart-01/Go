package test1

type Profile struct {
	Firstname string
	Lastname  string
}

func (p Profile) GetFirstname() string {
	return p.Firstname
}
