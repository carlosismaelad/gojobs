package main

import (
	"fmt"
	"log"
	"os"

	"github.com/carlosismaelad/gojobs/config"
	"github.com/carlosismaelad/gojobs/router"
)

var (
	logger *config.Logger
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

	logger = config.GetLogger("main")

	// initialize configs
	err := config.Init()
	if err != nil {
		logger.Errorf("Config initialization error: %v",err)
		return
	}

	// Initialize router
	router.Initialize()	
}
