package services

import (
	"github.com/ankitanwar/user-api/domain/users"
	"github.com/ankitanwar/user-api/utils/errors"
)

//CreateUser : To save the user in the database
func CreateUser(newUser users.User) (*users.User, *errors.RestError) {
	if err := newUser.Validate(); err != nil {
		return nil, err
	}
	if err := newUser.Save(); err != nil {
		return nil, err
	}
	return &newUser, nil
}

//GetUser : To get the detail of the user with given id
func GetUser(userid int) (*users.User, *errors.RestError) {
	if userid < 0 {
		return nil, errors.NewBadRequest("Enter the valied user id")
	}
	result := &users.User{
		ID: userid,
	}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}
