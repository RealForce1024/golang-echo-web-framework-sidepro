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
	for i := 0; i < 3; i++ {
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
		return c.String(http.StatusBadRequest, result)
	}
	return c.String(http.StatusOK, user.String())
}

func SaveUser(c echo.Context) error {
	v, ok := Users[3]
	if !ok {
		v = model.User{Id: 3, Name: "YaoMing"}
		//v = user
		Users[3] = v
	} else {
		v.Name = "updateYaoMing"
	}
	return c.String(http.StatusBadRequest, v.String())
}

func UpdateUser(c echo.Context) error {
	result := ""
	id := c.Param("id")
	nid, err := strconv.Atoi(id)
	if err != nil {
		result = fmt.Sprintf("参数有误，请校对:%v", id)
		return c.String(http.StatusBadRequest, result)
	}

	v, ok := Users[nid]
	if ok {
		v.Name += "update"
	} else {
		v = model.User{Id: nid, Name: "YaoMingAddAdd"}
	}
	Users[nid] = v
	return c.String(http.StatusOK, v.String())
}

func DeleteUser(c echo.Context) error {
	result := ""
	id := c.Param("id")
	nid, err := strconv.Atoi(id)

	if err != nil {
		result = fmt.Sprintf("参数有误，请校对:%v", id)
		return c.String(http.StatusBadRequest, result)
	}

	v, ok := Users[nid]
	if id == "" || !ok {
		return c.String(http.StatusBadRequest, fmt.Sprintf("id:%s not found", id))
	}

	result = fmt.Sprintf("delete %v ok", v)

	delete(Users,nid)
	return c.String(http.StatusOK, result)
}
