package main

import "fmt"

type User struct {
	ID    string
	Name  string
	Email string
}

func (u User) String() string {
	return fmt.Sprintf("[%s] %s | %s", u.ID, u.Name, u.Email)
}
