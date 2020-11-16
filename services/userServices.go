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

//UpdateUser : To update the values of the existing user
func UpdateUser(partial bool, user users.User) (*users.User, *errors.RestError) {
	current := users.User{
		ID: user.ID,
	}
	if err := current.Get(); err != nil {
		return nil, err
	}
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if partial {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}
		if user.LastName != "" {
			current.LastName = user.LastName
		}
		if user.Email != "" {
			current.Email = user.Email
		}
	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
	}
	if err := current.Update(); err != nil {
		return nil, err
	}
	return &current, nil
}

//DeleteUser : to delete the given user
func DeleteUser(userID int) *errors.RestError {
	user := &users.User{
		ID: userID,
	}
	return user.Delete()
}
