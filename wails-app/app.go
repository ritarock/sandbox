package main

import (
	"context"
	"fmt"
	"strconv"
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

func (a *App) Generate(digits string) (string, error) {
	fmt.Println(digits)
	i, err := strconv.Atoi(digits)
	if err != nil {
		return "", err
	}
	fmt.Println(i)
	return fmt.Sprintf("digits: %s", digits), nil
}
