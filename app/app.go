package main

import (
	"context"
	"fmt"
	"math/rand"
	"unicode"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) GeneratePassword(password_length int) string {
	alphabet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()_+"
	uppercase := false
	lowercase := false
	number := false
	special := false
	valid := false
	password := ""

	for !valid {
		for i := 0; i < password_length; i++ {
			char := string(alphabet[rand.Intn(len(alphabet))])
			password += char
			if unicode.IsUpper(rune(char[0])) {
				uppercase = true
			} else if unicode.IsLower(rune(char[0])) {
				lowercase = true
			} else if unicode.IsNumber(rune(char[0])) {
				number = true
			} else if unicode.IsPunct(rune(char[0])) {
				special = true
			}
		}
		if uppercase && lowercase && number && special {
			valid = true
		} else {
			password = ""
			uppercase = false
			lowercase = false
			number = false
			special = false
		}
	}
	return password
}
