package users

import (
	"fmt"
	"time"

	userdb "github.com/ankitanwar/user-api/databasource/postgres"
	"github.com/ankitanwar/user-api/utils/errors"
)

var (
	usersDB = make(map[int]*User)
)

//Save : To save the user into the database
func (user *User) Save() *errors.RestError {
	if usersDB[user.ID] != nil {
		return errors.NewBadRequest(fmt.Sprintf("User with %d already exist in the database", user.ID))
	}

	now := time.Now()
	user.DateCreated = now.Format("02-01-2006 15:05:04")

	usersDB[user.ID] = user
	return nil

}

//Get : To get the user from the database by given id
func (user *User) Get() *errors.RestError {
	if err := userdb.Client.Ping(); err != nil {
		panic(err)
	}
	result := usersDB[user.ID]
	if result == nil {
		err := errors.NewNotFound("User Not found in the database")
		return err
	}
	user.ID = result.ID
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated
	return nil
}
