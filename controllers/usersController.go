package controllers

import (
	"fmt"
	"net/http"
	"strconv"

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
	fmt.Println("Get user function is invoked")
	userid, userErr := strconv.Atoi(c.Param("user_id"))
	if userErr != nil {
		err := errors.NewBadRequest("Enter the valid used id")
		c.JSON(err.Status, err)
		return
	}
	user, err := services.GetUser(userid)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, user)

}

//SearchUser :To search for the particaular user
func SearchUser(c *gin.Context) {

}
