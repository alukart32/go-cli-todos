package usecase

type (
	ToDoListManager interface {
		Add(chan<- error, []string)
		Done(chan<- error, string)
		List(chan<- error)
		Clear(chan<- error)
	}

	ToDoListRepo interface {
		Add(chan<- error, []string)
		Done(chan<- error, uint64)
		List(chan<- *ToDoRepoResult)
		Clear(chan<- error)
	}
)
