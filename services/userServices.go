package services

import (
	"github.com/ankitanwar/user-api/domain/users"
	"github.com/ankitanwar/user-api/utils/errors"
)

//CreateUser : To save the user in the database
func CreateUser(newUser users.User) (*users.User, *errors.RestError) {
	return &newUser, nil
}
