package test2

import "struct-embed/test1"

type Profile struct {
	*test1.Profile
}

func (p Profile) GetLastname() string {
	return p.Lastname
}
