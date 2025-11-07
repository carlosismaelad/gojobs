package main

import (
	"fmt"
	"log"
	"os"

	"github.com/carlosismaelad/gojobs/router"
)

func loadBanner() string {
	banner, err := os.ReadFile("banner.txt")
	if err != nil {
		log.Printf("Error loading banner: %v", err)
		return "GoJobs API"
	}
	return string(banner)
}

func main() {
	fmt.Println()
	fmt.Print(loadBanner())
	fmt.Println()
	fmt.Println("🚀 Starting the GoJobs API...")
	fmt.Println()

	// Initialize router
	router.Initialize()	
}
