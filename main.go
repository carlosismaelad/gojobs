package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
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
	
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "API ok!",
    })
  })
	r.Run()
}
