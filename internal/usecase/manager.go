package usecase

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"alukart32.com/todos/internal/entity"
)

type TaskManager struct {
	repo ToDoListRepo
}

func NewManager(r ToDoListRepo) ToDoListManager {
	return &TaskManager{
		repo: r,
	}
}

func (m *TaskManager) Add(res chan<- error, todos []string) {
	log.Println(":: insert todos...")
	ch := make(chan error)
	defer close(ch)

	go m.repo.Add(ch, todos)

	err := <-ch
	res <- err
}

func (m *TaskManager) Done(res chan<- error, idx string) {
	v, err := strconv.Atoi(idx)
	if err != nil {
		res <- err
	}

	ch := make(chan error)
	defer close(ch)

	go m.repo.Done(ch, uint64(v))
	err = <-ch

	if err == nil {
		log.Printf(":: todo %s was removed from list", idx)
	}
	res <- err
}

func (m *TaskManager) List(res chan<- error) {
	ch := make(chan *ToDoRepoResult)
	defer close(ch)

	go m.repo.List(ch)
	result := <-ch
	if result.Err != nil {
		res <- result.Err
	}
	printList(result.ToDos)
	res <- nil
}

func (m *TaskManager) Clear(res chan<- error) {
	log.Println(":: clear todos list...")
	ch := make(chan error)
	defer close(ch)

	go m.repo.Clear(ch)
	err := <-ch

	res <- err
}

func printList(todos []entity.ToDo) {
	if len(todos) == 0 {
		log.Println(":: todos list is empty")
		return
	}

	var sb strings.Builder
	for _, t := range todos {
		sb.WriteString(fmt.Sprintf("\n\t- %s; id: %v", t.Desc, t.ID))
	}
	log.Println(":: print list...", "\ntodos list", sb.String())
}
