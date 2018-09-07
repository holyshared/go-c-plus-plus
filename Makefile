GO_C_BIN=go_c_bin

.PHONY: all test clean

all:
	go build -o $(GO_C_BIN)

test:
	./go_c_bin

clean:
	go clean
	rm $(GO_C_BIN)
