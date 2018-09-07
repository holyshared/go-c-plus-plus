GO_C_BIN=go_c_bin

ENTRY_CLANG=clang/main.go

.PHONY: all test clean

all:
	go build -o $(GO_C_BIN) $(ENTRY_CLANG)

test:
	./go_c_bin

clean:
	go clean
	rm $(GO_C_BIN)
