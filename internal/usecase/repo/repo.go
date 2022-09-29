package repo

import (
	"encoding/json"
	"fmt"

	"alukart32.com/todos/internal/entity"
	"alukart32.com/todos/internal/usecase"
	"alukart32.com/todos/pkg/errorx"
	"github.com/boltdb/bolt"
)

type ToDoListRepo struct {
	DB     *bolt.DB
	Bucket string
}

func NewToDoListRepo(db *bolt.DB, bucket string) usecase.ToDoListRepo {
	return &ToDoListRepo{
		DB:     db,
		Bucket: bucket,
	}
}

func (r *ToDoListRepo) Add(res chan<- error, todos []string) {
	res <- r.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(r.Bucket))
		for _, v := range todos {
			id, _ := b.NextSequence()
			todo := &entity.ToDo{
				ID:   id,
				Desc: v,
			}
			buf, err := json.Marshal(todo)
			if err != nil {
				return err
			}

			if err = b.Put(itob(id), buf); err != nil {
				return err
			}
		}
		return nil
	})
}
func (r *ToDoListRepo) Done(res chan<- error, idx uint64) {
	res <- r.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(r.Bucket))

		todo := b.Get(itob(idx))
		if len(todo) == 0 {
			return errorx.GetErr(fmt.Sprintf("the todo with id: %d wasn't found", idx))
		}
		return b.Delete(itob(idx))
	})
}
func (r *ToDoListRepo) List(res chan<- *usecase.ToDoRepoResult) {
	var list []entity.ToDo

	err := r.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(r.Bucket))
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			todo := entity.ToDo{}
			if err := json.Unmarshal(v, &todo); err != nil {
				return err
			}
			list = append(list, todo)
		}

		return nil
	})
	res <- &usecase.ToDoRepoResult{
		ToDos: list,
		Err:   err,
	}
}

func (m *ToDoListRepo) Clear(res chan<- error) {
	res <- m.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(m.Bucket))

		c := b.Cursor()
		for k, _ := c.First(); k != nil; k, _ = c.First() {
			if err := c.Delete(); err != nil {
				return err
			}
		}
		b.SetSequence(0)
		return nil
	})
}
