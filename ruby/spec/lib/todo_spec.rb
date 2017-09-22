require 'todo'

describe Todo do

  context "A single todo has attributes" do
    let(:message) { "Stare at lawn" }

    it "has a message in it" do
      subject.message = message
      expect(subject.message).to eq message
    end

    it "can be marked as completed" do
      subject.complete!
      expect(subject.completed).to be true
    end

    it "can be marked as incomplete" do
      subject.complete!
      subject.incomplete!
      expect(subject.completed).to be false
    end

  end

end
