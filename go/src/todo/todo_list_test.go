package todo

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("TodoList", func() {
	var (
		todoList todoList
		buyMilk  todo
		sellCow  todo
	)

	BeforeSuite(func() {
		buyMilk = NewTodo("Buy milk")
		sellCow = NewTodo("Sell cow")
	})

	BeforeEach(func() {
		todoList = NewTodoList()
	})

	Context("Storage of todos", func() {
		It("can add a todo", func() {
			todoList.Add(buyMilk)
			listLength := len(todoList.todos)

			Expect(listLength).To(Equal(1))
		})
	})

	Context("With multiple existing todos", func() {
		BeforeEach(func() {
			todoList.Add(buyMilk)
			todoList.Add(sellCow)
		})

		It("deletes a todo", func() {
			todoList.Delete(0)
			listLength := len(todoList.todos)

			Expect(listLength).To(Equal(1))
		})

		It("can find the second todo", func() {
			item := todoList.Find(1)
			Expect(item.Message).To(Equal("Sell cow"))
		})

		It("can find all completed todos", func() {
			completedItem := NewTodo("breathe")
			completedItem.Complete()
			todoList.Add(completedItem)

			items := todoList.AllCompleted()
			Expect(len(items)).To(Equal(1))
		})
	})

	Context("Modifying todos", func() {
		BeforeEach(func() {
			todoList.Add(buyMilk)
			todoList.Add(sellCow)
		})

		It("Completes a todo", func() {
			todoList.Complete(1)

			item := todoList.todos[1]
			Expect(item.Completed).To(BeTrue())
		})

		It("Deletes all completed todos", func() {
			todoList.Complete(0)
			todoList.Complete(1)

			todoList.DeleteAllCompleted()

			items := todoList.AllCompleted()
			Expect(len(items)).To(Equal(0))
		})
	})

})
