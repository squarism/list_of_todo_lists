package todo

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Go Suite")
}

var _ = Describe("Todo", func() {
	var (
		mowLawn todo
	)

	BeforeEach(func() {
		mowLawn = NewTodo("Mow the lawn")
	})

	Context("Todo items have attributes", func() {
		It("has a message in it", func() {
			Expect(mowLawn.Message).To(Equal("Mow the lawn"))
		})

		It("can be marked as complete", func() {
			mowLawn.Complete()
			Expect(mowLawn.Completed).To(BeTrue())
		})

		It("can be marked as incomplete", func() {
			mowLawn.Completed = true
			mowLawn.Incomplete()
			Expect(mowLawn.Completed).To(BeFalse())
		})
	})
})
