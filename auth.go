package main

import (
	"fmt"
)

// Authenticator interface
type UserType interface {
	Login(username, password string) bool
	Role() string
}

type User struct {
	Username string
	Password string
}

type Admin struct {
	Username  string
	Password  string
	SecretKey string
}

func (u User) Login(username, password string) bool {
	return u.Username == username && u.Password == password
}

func (u User) Role() string {
	return "User"
}

func (a Admin) Login(username, password string) bool {
	return a.Username == username && a.Password == password && a.SecretKey == "letmein"
}

func (a Admin) Role() string {
	return "Admin"
}

func Authenticate(u UserType, username, password string) {
	if u.Login(username, password) {
		fmt.Printf("%s logged in as: %s\n", username, u.Role())
	} else {
		fmt.Println("Invalid credentials!")
	}
}

func main() {
	user := User{Username: "user", Password: "1234"}
	admin := Admin{Username: "admin", Password: "adminpass", SecretKey: "letmein"}

	Authenticate(user, "user", "1234")        // ✅ valid user
	Authenticate(admin, "admin", "adminpass") // ✅ valid admin
	Authenticate(user, "user", "wrong")       // ❌ wrong password
}
