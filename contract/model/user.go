package model

type User struct {
	Id   string
	Role string
	Org  string
}

const PEDIATRICIAN_ROLE = "pediatrician"
const NEUROPATHOLOGIST_ROLE = "neuropathologist"

const HOSPITAL_ORG = "hospital"
const PARENTS_ORG = "parents"

//TODO: change  to CID
func NewUser() *User {
	return &User{Id: "max-parent-1", Role: "Parent", Org: "Parents"}
}

func (user *User) IsPediatrician() bool {
	return user.Role == PEDIATRICIAN_ROLE
}

func (user *User) IsNeuropathologist() bool {
	return user.Role == NEUROPATHOLOGIST_ROLE
}

func (user *User) IsHospitalWorker() bool {
	return user.Org == HOSPITAL_ORG
}

func (user *User) IsParent() bool {
	return user.Org == PARENTS_ORG
}
