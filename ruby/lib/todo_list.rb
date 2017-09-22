class TodoList
  attr_reader :todos

  def initialize
    @todos = []
  end

  def add(todo) 
    @todos << todo
  end

  def find_by_index(index)
    @todos[index]
  end

  def delete(delete_at_index)
    @todos.reject.each_with_index { |e, i| i == delete_at_index }
  end

  def delete!(delete_at_index)
    @todos = @todos.reject.each_with_index { |e, i| i == delete_at_index }
  end

  def all_completed
    @todos.select { |t| t.completed == true }
  end

  # oh man, getting away with murder here because of ruby memory references
  # In other languages, we'd want to clone a new item, set the attribute
  # and then return a new list that contains the modified item
  def complete!(index)
    todo = find_by_index(index)
    todo.complete!  # modify memory that just happens to be in @todos
  end

  # this is not functional style.  :(
  def delete_all_completed!
    @todos = @todos.reject { |todo| todo.completed == true }
  end

end
