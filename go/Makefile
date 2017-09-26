watch:
	@ginkgo watch -r --succinct --failFast

build:
	@gb build

run:
	@if [ -e ./bin/todo ]; then rm ./bin/todo; fi
	@gb build && if [ $$? -eq 0 ]; then ./bin/todo; fi

clean:
	if [ -d ./bin ]; then rm -rf ./bin; fi
	if [ -d ./pkg ]; then rm -rf ./pkg; fi
