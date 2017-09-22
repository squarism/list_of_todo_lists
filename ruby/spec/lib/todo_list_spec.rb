require 'todo_list'
require 'todo'

describe TodoList do

  # Normally, we'd have a testing data factory or something
  # to create fake data easier than this
  let(:buy_milk) do
    Todo.new.tap { |t| t.message = "Buy milk" }
  end

  let(:sell_cow) do
    Todo.new.tap { |t| t.message = "Sell cow" }
  end

  let(:breathe) do
    Todo.new.tap { |t| t.message = "Breathe"; t.complete! }
  end


  context "storage of todos" do

    it "can add a todo" do
      subject.add buy_milk
      expect(subject.todos.first.message).to eq "Buy milk"
    end

    context "with multiple existing todos" do

      before do
        subject.add buy_milk
        subject.add sell_cow
      end

      it "deletes a todo" do
        todos = subject.delete(1)
        expect(todos.length).to eq 1
      end

      it "destructively deletes a todo" do
        subject.delete!(1)
        subject.delete!(0)
        expect(subject.todos.length).to eq 0
      end

    end

  end

  context "finding todos" do

    before do
      subject.add buy_milk
      subject.add sell_cow
    end

    it "can find the second todo" do
      todo = subject.find_by_index(1)
      expect(todo.message).to eq "Sell cow"
    end

    it "can find all completed todos" do
      subject.add breathe

      completed = subject.all_completed
      expect(completed.length).to eq 1
    end

  end

  context "modifying todos" do

    before do
      subject.add buy_milk
      subject.add sell_cow
    end

    it "completes a todo" do
      subject.complete!(1)
      expect(subject.all_completed.length).to eq 1
    end

    it "deletes all completed todos" do
      subject.complete!(0)
      subject.complete!(1)
      subject.delete_all_completed!

      expect(subject.todos.length).to eq 0
    end

  end


end
