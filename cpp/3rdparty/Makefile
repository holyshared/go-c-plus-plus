CC=g++
TDIR=example
OBJS=math.o graphic.o
LIB=libexample.a

.PHONY: all clean test

all:
	$(CC) -c $(TDIR)/math/math.cpp -o math.o
	$(CC) -c $(TDIR)/graphic/graphic.cpp -o graphic.o
	ar rcs $(LIB) $(OBJS)

clean:
	-rm *.a *.o prog

test:
	g++ lintest.cpp -L./ libexample.a -o prog
	./prog
