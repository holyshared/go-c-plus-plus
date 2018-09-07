.PHONY: install all

all: cpp_bin

cpp_bin: main.go
	go install ./example1
	go build -o cpp_bin .

clean:
	rm cpp_bin
