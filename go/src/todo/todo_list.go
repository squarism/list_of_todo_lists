package todo

type todoList struct {
	todos []Todo
}

func NewTodoList() todoList {
	t := todoList{
		todos: []Todo{},
	}

	return t
}

func (t *todoList) Add(todo Todo) {
	t.todos = append(t.todos, todo)
}

func (t *todoList) Delete(index int) {
	// you can't delete at a slice index but you can make a new slice
	// that skips over the item at the index
	t.todos = append(t.todos[:index], t.todos[index+1:]...)
}

func (t *todoList) Find(index int) Todo {
	return t.todos[index]
}

// Oh man, I wish I had higher order functions!
func (t *todoList) AllCompleted() []Todo {
	var items []Todo
	for _, todo := range t.todos {
		if todo.Completed == true {
			items = append(items, todo)
		}
	}
	return items
}

// this wasn't in the ruby version, probably should be?
func (t *todoList) AllIncomplete() []Todo {
	var items []Todo
	for _, todo := range t.todos {
		if todo.Completed == false {
			items = append(items, todo)
		}
	}
	return items
}

func (t *todoList) Complete(index int) {
	t.todos[index].Complete()
}

func (t *todoList) DeleteAllCompleted() {
	// we delete all completed by only saving the incomplete items
	// this ... isn't obvious I think, could be better from an API
	// and naming standpoint
	t.todos = t.AllIncomplete()
}
