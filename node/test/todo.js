var assert = require('assert')
var { Todo } = require('../src/todo')
var expect = require('expect.js')

describe('Todo', function() {

  var todo

  beforeEach(function() {
    todo = new Todo("Walk the dog")
  })

  describe('attributes', function() {
    it('has a message', function() {
      expect(todo.message).to.equal("Walk the dog")
    })

    it('can be marked as complete', function() {
      todo.complete()
      expect(todo.completed).to.be(true)
    })

    it('can be marked as incomplete', function() {
      todo.complete()
      todo.incomplete()
      expect(todo.completed).to.be(false)
    })
  })

})
