package model

import (
	"fmt"
)

type User struct {
	Id   int
	Name string
}

func (u *User) String() string {
	return fmt.Sprintf("id:%d,name:%s", u.Id, u.Name)
}
