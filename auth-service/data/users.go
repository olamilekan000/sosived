package data

import "fmt"

var AllUsers *InmemoryDb

type User struct {
	Name     string `json:"username"`
	Password string `json:"password"`
}

type InmemoryDb struct {
	users []User
}

func (u *InmemoryDb) loadUsers() {

	u.users = append(u.users,
		User{
			Name:     "admin",
			Password: "password",
		},
		User{
			Name:     "captain",
			Password: "capitano",
		})

}

func (u *InmemoryDb) FindOne(name string) (*User, int) {
	for _, v := range u.users {
		if v.Name == name {
			return &v, 1
		}
	}
	return nil, 0
}

func init() {
	fmt.Println("Loading data...")
	AllUsers = &InmemoryDb{}
	AllUsers.loadUsers()

}
