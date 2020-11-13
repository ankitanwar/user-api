package users

import (
	"github.com/ankitanwar/user-api/utils/errors"
)

//Save : To save the user into the database
func (user *User) Save() *errors.RestError {
	return nil
}

//Get : To get the user from the database by given id
func (user *User) Get() (*User, *errors.RestError) {
	return nil, nil
}
