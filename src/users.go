package main

import (
	"errors"
	"fmt"
	"strings"
)

type User struct {
	id    string
	name  string
	email string
}

// NewUser valida y crea usuario (encapsulación + errores).
func NewUser(id, name, email string) (*User, error) {
	id = strings.TrimSpace(id)
	name = strings.TrimSpace(name)
	email = strings.TrimSpace(email)

	if id == "" || name == "" || email == "" {
		return nil, errors.New("id, nombre y email no pueden estar vacíos")
	}
	if !strings.Contains(email, "@") {
		return nil, errors.New("email inválido (debe contener @)")
	}

	return &User{id: id, name: name, email: email}, nil
}

// Getters
func (u *User) ID() string { return u.id }

func (u *User) String() string {
	return fmt.Sprintf("[%s] %s | %s", u.id, u.name, u.email)
}
