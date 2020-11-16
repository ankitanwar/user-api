package users

import (
	"fmt"
	"strings"
	"time"

	userdb "github.com/ankitanwar/user-api/databasource/postgres"
	"github.com/ankitanwar/user-api/utils/errors"
)

const (
	insertUser = "INSERT INTO users(first_name,last_name,email,date_created)VALUES(?,?,?,?) "
	getUser    = "SELECT id,first_name,last_name,email,date_created FROM users WHERE id=?;"
	errNoRows  = "no rows in result set"
	updateUser = "UPDATE users SET first_name=?,last_name=?,email=? WHERE id=?"
	deleteUser = "DELETE FROM users WHERE id=?"
)

//Save : To save the user into the database
func (user *User) Save() *errors.RestError {
	stmt, err := userdb.Client.Prepare(insertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	now := time.Now()
	user.DateCreated = now.Format("02-01-2006 15:04")
	insert, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)

	// insert, err := stmt.Exec(insertUser, user.FirstName, user.LastName, user.Email, user.DateCreated) we can also do it like this

	if err != nil {
		if strings.Contains(err.Error(), "users.email_UNIQUE") {
			return errors.NewBadRequest(fmt.Sprintf("User with %s already exist in the database", user.Email))
		}
		return errors.NewBadRequest(err.Error())
	}
	userid, err := insert.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	user.ID = int(userid)
	return nil

}

//Get : To get the user from the database by given id
func (user *User) Get() *errors.RestError {
	stmt, err := userdb.Client.Prepare(getUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	result := stmt.QueryRow(user.ID) //to query the single row
	if err := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil {
		if strings.Contains(err.Error(), errNoRows) {
			return errors.NewNotFound(fmt.Sprintf("No user with exist with id %v ", user.ID))
		}
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}

//Update : To  update the value of the existing users
func (user *User) Update() *errors.RestError {
	stmt, err := userdb.Client.Prepare(updateUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.ID)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	return nil

}

//Delete : To delete the user from the database
func (user *User) Delete() *errors.RestError {
	stmt, err := userdb.Client.Prepare(deleteUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.ID)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}
