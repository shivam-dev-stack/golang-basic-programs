package main

import (
	"fmt"
)

type User interface {
	Login(username, password string) (string, bool)
	Role() string
}

type Credentials struct {
	Username string
	Password string
}

// Student struct
type Student struct {
	Username string
	Email    string
	Password string
}

func (s Student) Login(username, password string) (string, bool) {
	if s.Username == username && s.Password == password {
		return s.Email, true
	}
	return "", false
}

func (s Student) Role() string {
	return "Student"
}

// Instructor struct
type Instructor struct {
	Username  string
	Email     string
	Password  string
	Expertise string
}

func (i Instructor) Login(username, password string) (string, bool) {
	if i.Username == username && i.Password == password {
		return i.Email, true
	}
	return "", false
}

func (i Instructor) Role() string {
	return "Instructor"
}

type Admin struct {
	Username  string
	Email     string
	Password  string
	SecretKey string
}

func (a Admin) Login(username, password string) (string, bool) {
	if a.Username == username && a.Password == password && a.SecretKey == "letmein" {
		return a.Email, true
	}
	return "", false
}

func (a Admin) Role() string {
	return "Admin"
}

func Authenticate(u User, username, password string) bool {
	if email, ok := u.Login(username, password); ok {
		fmt.Printf("%s is Logged in as %s (%s)\n", username, u.Role(), email)
		return true
	}
	fmt.Println("Login failed!")
	return false
}

func Dashboard(u User) {
	switch u.Role() {
	case "Student":
		fmt.Println("Welcome to the student dashboard.")
	case "Instructor":
		fmt.Println("Welcome to the instructor dashboard.")
	case "Admin":
		fmt.Println("Admin panel access granted.")
	default:
		fmt.Println("Unknown role.")
	}
}

func main() {
	student := Student{"rahul", "rrr@go.com", "1234"}
	instructor := Instructor{"neha", "nti@go.com", "admin123", "Go Programming"}
	admin := Admin{"shiv", "admin@go.com", "adminpass", "letmein"}

	users := []User{student, instructor, admin}
	creds := []Credentials{
		{"rahul", "1234"},
		{"neha", "wrongpass"}, // Intentional wrong password
		{"shiv", "adminpass"}, // With correct SecretKey embedded
	}

	for i, u := range users {
		if Authenticate(u, creds[i].Username, creds[i].Password) {
			Dashboard(u)
		}
	}

}
