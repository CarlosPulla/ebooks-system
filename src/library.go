package main

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

type Library struct {
	Books map[string]*Book
	Users map[string]*User
}

func NewLibrary() *Library {
	return &Library{
		Books: make(map[string]*Book),
		Users: make(map[string]*User),
	}
}

// -------------------- BOOKS --------------------

func (l *Library) AddBook(b Book) error {
	b.ID = strings.TrimSpace(b.ID)
	if b.ID == "" {
		return errors.New("el ID del libro no puede estar vacío")
	}
	if _, exists := l.Books[b.ID]; exists {
		return fmt.Errorf("ya existe un libro con ID %s", b.ID)
	}

	b.Available = true
	b.BorrowedBy = ""
	l.Books[b.ID] = &b
	return nil
}

func (l *Library) RemoveBook(bookID string) error {
	bookID = strings.TrimSpace(bookID)
	b, ok := l.Books[bookID]
	if !ok {
		return fmt.Errorf("no existe libro con ID %s", bookID)
	}
	if !b.Available {
		return fmt.Errorf("no se puede eliminar: el libro %s está prestado", bookID)
	}
	delete(l.Books, bookID)
	return nil
}

func (l *Library) FindBooksByTitle(query string) []*Book {
	query = strings.ToLower(strings.TrimSpace(query))
	var result []*Book
	for _, b := range l.Books {
		if strings.Contains(strings.ToLower(b.Title), query) {
			result = append(result, b)
		}
	}
	sort.Slice(result, func(i, j int) bool { return result[i].Title < result[j].Title })
	return result
}

func (l *Library) ListBooks() []*Book {
	var list []*Book
	for _, b := range l.Books {
		list = append(list, b)
	}
	sort.Slice(list, func(i, j int) bool { return list[i].ID < list[j].ID })
	return list
}

// -------------------- USERS --------------------

func (l *Library) AddUser(u User) error {
	u.ID = strings.TrimSpace(u.ID)
	if u.ID == "" {
		return errors.New("el ID del usuario no puede estar vacío")
	}
	if _, exists := l.Users[u.ID]; exists {
		return fmt.Errorf("ya existe un usuario con ID %s", u.ID)
	}

	l.Users[u.ID] = &u
	return nil
}

func (l *Library) RemoveUser(userID string) error {
	userID = strings.TrimSpace(userID)
	if _, ok := l.Users[userID]; !ok {
		return fmt.Errorf("no existe usuario con ID %s", userID)
	}

	// No permitir eliminar si tiene libros prestados
	for _, b := range l.Books {
		if b.BorrowedBy == userID {
			return fmt.Errorf("no se puede eliminar: el usuario %s tiene libros prestados", userID)
		}
	}
	delete(l.Users, userID)
	return nil
}

func (l *Library) ListUsers() []*User {
	var list []*User
	for _, u := range l.Users {
		list = append(list, u)
	}
	sort.Slice(list, func(i, j int) bool { return list[i].ID < list[j].ID })
	return list
}

// -------------------- LOANS --------------------

func (l *Library) BorrowBook(bookID, userID string) error {
	bookID = strings.TrimSpace(bookID)
	userID = strings.TrimSpace(userID)

	b, ok := l.Books[bookID]
	if !ok {
		return fmt.Errorf("no existe libro con ID %s", bookID)
	}
	if _, ok := l.Users[userID]; !ok {
		return fmt.Errorf("no existe usuario con ID %s", userID)
	}
	if !b.Available {
		return fmt.Errorf("el libro %s ya está prestado", bookID)
	}

	b.Available = false
	b.BorrowedBy = userID
	return nil
}

func (l *Library) ReturnBook(bookID string) error {
	bookID = strings.TrimSpace(bookID)

	b, ok := l.Books[bookID]
	if !ok {
		return fmt.Errorf("no existe libro con ID %s", bookID)
	}
	if b.Available {
		return fmt.Errorf("el libro %s ya está disponible (no estaba prestado)", bookID)
	}

	b.Available = true
	b.BorrowedBy = ""
	return nil
}
