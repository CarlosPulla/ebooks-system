package main

import (
	"errors"
	"fmt"
	"strings"
)

type Book struct {
	id         string
	title      string
	author     string
	year       int
	genre      string
	available  bool
	borrowedBy string // userID si está prestado
}

// NewBook aplica encapsulación: el libro solo se crea si es válido.
func NewBook(id, title, author string, year int, genre string) (*Book, error) {
	id = strings.TrimSpace(id)
	title = strings.TrimSpace(title)
	author = strings.TrimSpace(author)
	genre = strings.TrimSpace(genre)

	if id == "" || title == "" || author == "" {
		return nil, errors.New("id, título y autor no pueden estar vacíos")
	}
	if year < 0 {
		return nil, errors.New("año inválido")
	}

	return &Book{
		id:         id,
		title:      title,
		author:     author,
		year:       year,
		genre:      genre,
		available:  true,
		borrowedBy: "",
	}, nil
}

// Getters (encapsulación)
func (b *Book) ID() string    { return b.id }
func (b *Book) Title() string { return b.title }

func (b *Book) IsAvailable() bool { return b.available }

// BorrowTo cambia el estado del libro de forma controlada.
func (b *Book) BorrowTo(userID string) error {
	userID = strings.TrimSpace(userID)
	if userID == "" {
		return errors.New("userID no puede estar vacío")
	}
	if !b.available {
		return fmt.Errorf("el libro %s ya está prestado", b.id)
	}

	b.available = false
	b.borrowedBy = userID
	return nil
}

// Return devuelve el libro a disponible.
func (b *Book) Return() error {
	if b.available {
		return fmt.Errorf("el libro %s ya está disponible", b.id)
	}
	b.available = true
	b.borrowedBy = ""
	return nil
}

func (b *Book) String() string {
	status := "Disponible"
	if !b.available {
		status = fmt.Sprintf("Prestado (Usuario: %s)", b.borrowedBy)
	}
	return fmt.Sprintf("[%s] %s - %s (%d) | Género: %s | %s",
		b.id, b.title, b.author, b.year, b.genre, status)
}
