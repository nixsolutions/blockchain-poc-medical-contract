package model

type User struct {
	Id   string
	Org  string
}


const HOSPITAL_ORG = "hospital"
const PARENTS_ORG = "parents"


func NewUser(id, org string) *User {
	return &User{Id: id, Org: org}
}
func (user *User) IsHospitalWorker() bool {
	return user.Org == HOSPITAL_ORG
}

func (user *User) IsParent() bool {
	return user.Org == PARENTS_ORG
}
