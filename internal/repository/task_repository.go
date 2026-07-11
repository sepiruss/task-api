package repository

import (
	"database/sql"

	"github.com/sepiruss/task-api/internal/model"
)

type TaskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) *TaskRepository {
	return &TaskRepository{
		db: db,
	}
}

func (r *TaskRepository) GetAll() ([]model.Task, error) {
	rows, err := r.db.Query(`
		SELECT id, title, description, completed, created_at, updated_at
		FROM tasks
		ORDER BY id DESC
	`)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var tasks []model.Task

	for rows.Next() {
		var task model.Task

		err := rows.Scan(
			&task.ID,
			&task.Title,
			&task.Description,
			&task.Completed,
			&task.CreatedAt,
			&task.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (r *TaskRepository) Create(task *model.Task) error {
	query := `
		INSERT INTO tasks (title, description, completed)
		VALUES ($1, $2, $3)
		RETURNING id, created_at, updated_at
	`

	return r.db.QueryRow(
		query,
		task.Title,
		task.Description,
		task.Completed,
	).Scan(
		&task.ID,
		&task.CreatedAt,
		&task.UpdatedAt,
	)
}

func (r *TaskRepository) GetByID(id int) (*model.Task, error) {

	var task model.Task

	query := `
		SELECT id, title, description, completed, created_at, updated_at
		FROM tasks
		WHERE id = $1
	`

	err := r.db.QueryRow(query, id).Scan(
		&task.ID,
		&task.Title,
		&task.Description,
		&task.Completed,
		&task.CreatedAt,
		&task.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (r *TaskRepository) Update(task *model.Task) error {

	query := `
		UPDATE tasks
		SET
			title = $1,
			description = $2,
			completed = $3,
			updated_at = CURRENT_TIMESTAMP
		WHERE id = $4
		RETURNING updated_at
	`

	return r.db.QueryRow(
		query,
		task.Title,
		task.Description,
		task.Completed,
		task.ID,
	).Scan(&task.UpdatedAt)
}

func (r *TaskRepository) Delete(id int) error {

	query := `
		DELETE FROM tasks
		WHERE id = $1
	`

	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
