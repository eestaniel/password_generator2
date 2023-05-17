package main

import (
	"context"
	"fmt"
	"math/rand"
	"strings"
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

func (a *App) GeneratePassword(password_length int, doUpper bool, doLower bool, doNumber bool, doSpecial bool) string {
	types := []string{}
	if doUpper {
		types = append(types, "ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	}
	if doLower {
		types = append(types, "abcdefghijklmnopqrstuvwxyz")
	}
	if doNumber {
		types = append(types, "0123456789")
	}
	if doSpecial {
		types = append(types, "!@#$%^&*()_+")
	}
	if len(types) == 0 {
		return "Please select at least one character type."
	}

	alphabet := strings.Join(types, "")

	for {
		// Generate a password
		password := ""
		for i := 0; i < password_length; i++ {
			password += string(alphabet[rand.Intn(len(alphabet))])
		}

		// Check if it contains at least one character from each type
		valid := true
		for _, t := range types {
			if !strings.ContainsAny(password, t) {
				valid = false
				break
			}
		}

		// If the password is valid, return it
		if valid {
			return password
		}

		// If not, generate a new password in the next iteration
	}
}
