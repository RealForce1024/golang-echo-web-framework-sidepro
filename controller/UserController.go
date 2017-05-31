package controller

import (
	"github.com/labstack/echo"
	"net/http"
	"myecho/model"
	"strconv"
	"fmt"
)

var Users = make(map[int]model.User, 5)

func init() {
	for i := 0; i < 5; i++ {
		//Users = append(, model.User{Id: i, Name: "kobe" + strconv.Itoa(i+1)})
		Users[i] = model.User{Id: i, Name: "kobe" + strconv.Itoa(i+1)}
		fmt.Printf("%v\n", Users)
	}
}

func GetUser(c echo.Context) error {
	id := c.Param("id")
	result := id
	nid, _ := strconv.Atoi(id)
	user, ok := Users[nid]
	if !ok {
		result = fmt.Sprintf("not found this id:%d", nid)
		return c.String(http.StatusOK, result)
	}
	//return c.String(http.StatusOK, user.Name)
	return c.String(http.StatusOK, user.String())
}
