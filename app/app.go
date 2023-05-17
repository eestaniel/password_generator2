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

func (a *App) GeneratePassword(password_length int, doUpper bool, doLower bool, doNumber bool, doSpecial bool) string {
	alphabet := ""
	flags := make(map[string]bool)
	if doUpper == true {
		alphabet += "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		flags["Uppercase"] = false
	}
	if doLower == true {
		alphabet += "abcdefghijklmnopqrstuvwxyz"
		flags["Lowercase"] = false
	}
	if doNumber == true {
		alphabet += "0123456789"
		flags["Number"] = false
	}
	if doSpecial == true {
		alphabet += "!@#$%^&*()_+"
		flags["Special"] = false
	}
	if doUpper == false && doLower == false && doNumber == false && doSpecial == false {
		return "Please select at least one character type."
	}
	password := ""
	valid := false
	for !valid {
		for i := 0; i < password_length; i++ {
			char := string(alphabet[rand.Intn(len(alphabet))])
			password += char
			for key, _ := range flags {
				if unicode.IsUpper(rune(char[0])) && key == "Uppercase" {
					flags[key] = true
				} else if unicode.IsLower(rune(char[0])) && key == "Lowercase" {
					flags[key] = true
				} else if unicode.IsNumber(rune(char[0])) && key == "Number" {
					flags[key] = true
				} else if unicode.IsPunct(rune(char[0])) && key == "Special" {
					flags[key] = true
				}
			}
		}
		for _, value := range flags {
			if value == false {
				password = ""
				for key, _ := range flags {
					flags[key] = false
				}
				continue
			}
		}
		valid = true
	}
	return password

}
