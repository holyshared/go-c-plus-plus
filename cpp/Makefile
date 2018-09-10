.PHONY: install all

all: cpp_bin

cpp_bin: main.go
	make -C $(PWD)/3rdparty
	go install ./example1
	go install ./example2
	go install ./example3
	go build -o cpp_bin .

clean:
	make -C $(PWD)/3rdparty clean
	-rm cpp_bin
