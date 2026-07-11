package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sepiruss/task-api/internal/config"
	"github.com/sepiruss/task-api/internal/database"
	"github.com/sepiruss/task-api/internal/handler"
	"github.com/sepiruss/task-api/internal/repository"
	"github.com/sepiruss/task-api/internal/routes"
	"github.com/sepiruss/task-api/internal/service"
)

func main() {

	fmt.Println("Task API")

	cfg := config.Load()

	db := database.Connect(cfg)
	defer db.Close()

	repo := repository.NewTaskRepository(db)

	taskService := service.NewTaskService(repo)

	taskHandler := handler.NewTaskHandler(taskService)

	mux := http.NewServeMux()

	routes.RegisterRoutes(mux, taskHandler)

	fmt.Println("🚀 Server running on http://localhost:8080")

	log.Fatal(http.ListenAndServe(":8080", mux))
}
