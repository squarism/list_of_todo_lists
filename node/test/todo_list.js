var assert = require('assert')
var expect = require('expect.js')

var { Todo } = require('../src/todo')
var { TodoList } = require('../src/todo_list')


describe('TodoList', function() {

  var list
  var buyMilk, sellCow, completedItem

  beforeEach(function() {
    todoList = new TodoList()

    buyMilk = new Todo("Buy milk")
    sellCow = new Todo("Sell cow")
  })

  context('Storage of items', function() {
    it('can add a todo', function() {
      todoList.add(buyMilk)
      expect(todoList.todos.length).to.equal(1)
    })
  })

  context('With multiple existing todos', function() {
    beforeEach(function() {
      todoList.add(buyMilk)
      todoList.add(sellCow)
    })

    it('deletes a todo', function() {
      todoList.delete(0)
      expect(todoList.todos[0].message).to.equal("Sell cow")
    })

    it('finds the second todo', function() {
      todo = todoList.find(1)
      expect(todo.message).to.equal("Sell cow")
    })

    context('With one completed todo', function() {
      beforeEach(function() {
        completedItem = new Todo("Breathe")
        completedItem.complete()
        todoList.add(completedItem)
      })

      it('can find all completed todos', function() {
        completedTodos = todoList.allCompleted()
        expect(completedTodos.length).to.equal(1)
      })
    })

  })

  context('When modifying todos', function() {
    beforeEach(function() {
      todoList.add(buyMilk)
      todoList.add(sellCow)
    })

    it('completes a todo', function() {
      todoList.complete(1)
      expect(todoList.todos[1]).to.have.property('completed', true)
    })

    it('deletes all completed todos', function() {
      todoList.complete(0)
      todoList.complete(1)

      todoList.deleteAllCompleted()

      expect(todoList.allCompleted().length).to.equal(0)
    })

  })

})
