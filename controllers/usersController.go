package controllers

import (
	"net/http"
	"strconv"

	"github.com/ankitanwar/user-api/domain/users"
	"github.com/ankitanwar/user-api/services"
	"github.com/ankitanwar/user-api/utils/errors"
	"github.com/gin-gonic/gin"
)

func getUserid(userIDParam string) (int, *errors.RestError) {
	userID, userErr := strconv.Atoi(userIDParam)
	if userErr != nil {
		err := errors.NewBadRequest("Enter the valid used id")
		return 0, err
	}
	return userID, nil
}

//CreateUser : To create the user
func CreateUser(c *gin.Context) {
	var newUser users.User
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
	userid, userErr := getUserid(c.Param("user_id"))
	if userErr != nil {
		c.JSON(userErr.Status, userErr)
		return
	}
	user, err := services.GetUser(userid)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, user)

}

//UpdateUser :To Update the value of particaular user
func UpdateUser(c *gin.Context) {
	var user = users.User{}
	userid, userErr := getUserid(c.Param("user_id"))
	if userErr != nil {
		c.JSON(userErr.Status, userErr)
		return
	}
	user.ID = userid
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequest("Invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	isPartial := c.Request.Method == http.MethodPatch

	updatedUser, err := services.UpdateUser(isPartial, user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, updatedUser)
}

//DeleteUser :To Delete the user with given id
func DeleteUser(c *gin.Context) {
	userid, userErr := getUserid(c.Param("user_id"))
	if userErr != nil {
		c.JSON(userErr.Status, userErr)
		return
	}
	if err := services.DeleteUser(userid); err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"Status": "User Deleted"})
}
