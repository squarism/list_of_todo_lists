class Todo
  attr_accessor :message
  attr_reader :completed

  def initialize
    @completed = false
  end

  def complete!
    @completed = true
  end

  def incomplete!
    @completed = false
  end

end
