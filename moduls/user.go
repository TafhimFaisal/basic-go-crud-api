package moduls

import "errors"

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
