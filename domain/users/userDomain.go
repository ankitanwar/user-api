package users

import "github.com/ankitanwar/user-api/utils/errors"

//User : User and its values
type User struct {
	ID          int    `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
}

//Validate : To validate the users
func (user *User) Validate() *errors.RestError {
	if user.Email == "" {
		err := errors.NewBadRequest("Please enter the valid mail address")
		return err
	}
	return nil
}
