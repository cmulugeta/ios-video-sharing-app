package main

import (
	"github.com/joho/godotenv"
	"github.com/cmulugeta/ios-video-sharing-app/windmill-backend/app"
	"log"
)

// Init function
func init() {
	// Load .env file to access environment vars
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

// Main function
func main() {
	app := &app.App{}
	app.Init()
}