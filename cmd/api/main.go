package main

import (
	"fmt"

	"github.com/sepiruss/task-api/internal/config"
	"github.com/sepiruss/task-api/internal/database"
)

func main() {

	fmt.Println("Task API")

	cfg := config.Load()

	db := database.Connect(cfg)
	defer db.Close()

	fmt.Println("Application started successfully.")
}
