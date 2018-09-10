# g++ lintest.cpp libexample.a -o prog
CC=g++
SRCS=./example/math/math.cpp
HDRS=./example/math/math.h
OBJS=./example.o
LIBS=./libexample.a

.PHONY: all install uninstall clean

all:
	$(CC) -c $(SRCS) -o $(OBJS)
	ar rcs $(LIBS) $(OBJS)

install:
	mkdir -p /usr/local/include/example/math
	cp $(LIBS) /usr/local/lib/$(LIBS)
	cp $(HDRS) /usr/local/include/$(HDRS)

uninstall:
	rm /usr/local/lib/$(LIBS)
	rm -rf /usr/local/include/$(HDRS)

clean:
	-rm *.a *.o prog
