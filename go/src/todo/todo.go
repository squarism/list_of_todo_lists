package todo

type Todo struct {
	Message   string
	Completed bool
}

func NewTodo(message string) Todo {
	return Todo{
		Message:   message,
		Completed: false,
	}
}

func (t *Todo) Complete() {
	t.Completed = true
}

func (t *Todo) Incomplete() {
	t.Completed = false
}
