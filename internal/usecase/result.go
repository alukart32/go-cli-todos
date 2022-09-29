package usecase

import "alukart32.com/todos/internal/entity"

type ToDoRepoResult struct {
	ToDos []entity.ToDo
	Err   error
}
