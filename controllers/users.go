package controllers

import (
	"fmt"
	"net/http"

	"github.com/ankitanwar/user-api/domain/users"
	"github.com/ankitanwar/user-api/services"
	"github.com/ankitanwar/user-api/utils/errors"
	"github.com/gin-gonic/gin"
)

//CreateUser : To create the user
func CreateUser(c *gin.Context) {
	var newUser users.User
	fmt.Print(newUser)

	// bytes, err := ioutil.ReadAll(c.Request.Body)
	// if err != nil {
	// 	return
	// }
	// if err := json.Unmarshal(bytes, &newUser); err != nil {
	// 	return
	// }

	//This line is equal to the comment code above
	if err := c.ShouldBindJSON(&newUser); err != nil {
		err := errors.NewBadRequest("Invalid Request")
		c.JSON(err.Status, err)
		return
	}

	result, saverr := services.CreateUser(newUser)
	if saverr != nil {
		c.JSON(saverr.Status, saverr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

//GetUser : To get the user from the database
func GetUser(c *gin.Context) {

}

//SearchUser :To search for the particaular user
func SearchUser(c *gin.Context) {

}
