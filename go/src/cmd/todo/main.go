package main

import (
	"os"
	"todo"

	"github.com/olekukonko/tablewriter"
)

func main() {
	// fake UI interaction
	list := todo.NewTodoList()

	list.Add(todo.NewTodo("Do a little dance"))
	list.Add(todo.NewTodo("Make a little love"))
	list.Add(todo.NewTodo("Get down tonight"))

	list.Complete(2)

	// table printing
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Todo", "Completed"})
	table.SetBorder(false)

	AppendToTable(table, list.AllIncomplete())
	AppendToTable(table, list.AllCompleted())

	table.Render()
}

func AppendToTable(table *tablewriter.Table, todos []todo.Todo) {
	for _, todo := range todos {
		var completedSymbol string
		if todo.Completed == true {
			completedSymbol = "[x]"
		} else {
			completedSymbol = "[ ]"
		}
		row := []string{todo.Message, completedSymbol}
		table.Append(row)
	}
}
