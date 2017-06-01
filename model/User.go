package model

import (
	"fmt"
)

type User struct {
	Id    int
	Name  string `json:"name" xml:"name" form:"name" query:"name"`
	Email string `json:"email" xml:"email" form:"email" query:"email"`
}

func (u *User) String() string {
	return fmt.Sprintf("id:%d,name:%s", u.Id, u.Name)
}
