package service

import (
	"errors"
	"strings"

	"github.com/sepiruss/task-api/internal/model"
	"github.com/sepiruss/task-api/internal/repository"
)

type TaskService struct {
	repo *repository.TaskRepository
}

func NewTaskService(repo *repository.TaskRepository) *TaskService {
	return &TaskService{
		repo: repo,
	}
}

func (s *TaskService) GetAll() ([]model.Task, error) {
	return s.repo.GetAll()
}

func (s *TaskService) Create(task *model.Task) error {
	if strings.TrimSpace(task.Title) == "" {
		return errors.New("title is required")
	}
	return s.repo.Create(task)
}

func (s *TaskService) GetByID(id int) (*model.Task, error) {

	return s.repo.GetByID(id)

}
