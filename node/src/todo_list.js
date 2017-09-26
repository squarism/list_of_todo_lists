var _ = require('lodash');

class TodoList {

  constructor () {
    this.todos = []
  }

  add(todo) {
    this.todos.push(todo)
  }

  find(index) {
    return this.todos[index]
  }

  delete(index) {
    // meh, mutation here
    this.todos = _.reject(this.todos, function(element, i) {
      return i === index
    })
  }

  deleteAllCompleted() {
    this.todos = _.reject(this.todos, function(element) {
      return element.completed === true
    })
  }
  
  complete(index) {
    this.todos[index].complete()
  }

  allCompleted() {
    return _.filter(this.todos, 'completed')
  }

}

exports.TodoList = TodoList
