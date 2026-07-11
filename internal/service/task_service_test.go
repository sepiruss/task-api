package service

import (
	"testing"

	"github.com/sepiruss/task-api/internal/model"
)

type FakeRepository struct {
	CreateCalled bool
}

func (f *FakeRepository) GetAll() ([]model.Task, error) {
	return nil, nil
}

func (f *FakeRepository) GetByID(id int) (*model.Task, error) {
	return nil, nil
}

func (f *FakeRepository) Update(task *model.Task) error {
	return nil
}

func (f *FakeRepository) Delete(id int) error {
	return nil
}

func (f *FakeRepository) Create(task *model.Task) error {
	f.CreateCalled = true
	return nil
}

func TestCreateTaskSuccess(t *testing.T) {

	repo := &FakeRepository{}

	service := NewTaskService(repo)

	task := &model.Task{
		Title: "Learn Go",
	}

	err := service.Create(task)

	if err != nil {
		t.Errorf("expected nil error, got %v", err)
	}

	if !repo.CreateCalled {
		t.Error("expected Create repository method to be called")
	}
}

func TestCreateTaskWithEmptyTitle(t *testing.T) {

	repo := &FakeRepository{}

	service := NewTaskService(repo)

	task := &model.Task{
		Title: "",
	}

	err := service.Create(task)

	if err == nil {
		t.Error("expected error but got nil")
	}
}
