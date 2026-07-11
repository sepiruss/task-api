package routes

import (
	"net/http"

	"github.com/sepiruss/task-api/internal/handler"
)

func RegisterRoutes(mux *http.ServeMux, taskHandler *handler.TaskHandler) {

	mux.HandleFunc("GET /tasks", taskHandler.GetAll)

	mux.HandleFunc("GET /tasks/{id}", taskHandler.GetByID)

	mux.HandleFunc("POST /tasks", taskHandler.Create)

	//mux.HandleFunc("PUT /tasks/{id}", taskHandler.Update)

	//mux.HandleFunc("DELETE /tasks/{id}", taskHandler.Delete)
}
