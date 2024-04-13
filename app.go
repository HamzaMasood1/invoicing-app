package main

import (
	"changeme/backend"
	"context"
	"database/sql"
	"fmt"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called at application startup
func (a *App) startup(ctx context.Context) {
	// Perform your setup here
	a.ctx = ctx
}

// domReady is called after the front-end dom has been loaded
func (a *App) domReady(ctx context.Context) {
	// Add your action here
}

// beforeClose is called when the application is about to quit,
// either by clicking the window close button or calling runtime.Quit.
// Returning true will cause the application to continue,
// false will continue shutdown as normal.
func (a *App) beforeClose(ctx context.Context) (prevent bool) {
	return false
}

// shutdown is called at application termination
func (a *App) shutdown(ctx context.Context) {
	// Perform your teardown here
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

// Saves to the DB
func (a *App) Save(firstName string, lastName string) string {
	newName := backend.Name{
		FirstName: firstName,
		LastName:  lastName,
	}
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		fmt.Println(err)
	}
	backend.AddName(db, newName)
	defer db.Close()
	return fmt.Sprintf("User has been saved")
}
