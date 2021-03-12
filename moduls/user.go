package moduls

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
	u.ID = nextID
	nextID++
	users = append(users, &u)

	return u, nil
}
