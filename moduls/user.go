package moduls

import (
	"errors"
	"fmt"
)

// User struct
type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users  []*User
	nextID = 1
)

// GetUsers getting all users memory location
func GetUsers() []*User {
	return users
}

// AddUser add user memory location to users array
func AddUser(u User) (User, error) {
	if u.ID != 0 {
		return User{}, errors.New("New user must not use ID")
	}
	u.ID = nextID
	nextID++
	users = append(users, &u)

	return u, nil
}

func GetUserByID(id int) (User, error) {
	for _, u := range users {
		if u.ID == id {
			return *u, nil
		}
	}

	return User{}, fmt.Errorf("User with ID '%v' not found", id)
}

func UpdateUser(u User) (User, error) {
	for i, candidate := range users {
		if candidate.ID == u.ID {
			users[i] = &u
			return u, nil
		}
	}

	return User{}, fmt.Errorf("User with ID '%v' not found", u.ID)
}

func RemoveUserbyId(id int) error {
	for i, u := range users {
		if u.ID == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("user with ID '%v' not found", id)
}
