class Todo {

  constructor (message) {
    this.message = message
    this.completed = false
  }

  complete() {
    this.completed = true
  }

  incomplete() {
    this.completed = false
  }

}

exports.Todo = Todo
