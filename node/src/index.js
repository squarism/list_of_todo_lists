var _ = require('lodash');
var { table, getBorderCharacters } = require('table')

var { Todo } = require('../src/todo')
var { TodoList } = require('../src/todo_list')

todoList = new TodoList()
todoList.add(new Todo("Walk the dog"))
todoList.add(new Todo("Pet the dog"))
todoList.add(new Todo("Forget it, I don't have a dog"))
todoList.complete(2)

data = _.map(todoList.todos, function(todo) {
  let completedSymbol
  if (todo.completed === true) {
    completedSymbol = '[x]'
  } else {
    completedSymbol = '[ ]'
  }

  return [todo.message, completedSymbol]
})


config = {
    border: getBorderCharacters(`norc`)
};

console.log(table(data, config))
