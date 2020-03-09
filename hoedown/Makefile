CC = gcc 
CFLAGS = -O3

HOEDOWN_SRC = \
		src/autolink.o \
		src/buffer.o \
		src/document.o \
		src/escape.o \
		src/html.o \
		src/html_blocks.o \
		src/html_smartypants.o \
		src/stack.o \
		src/version.o

MAIN = main.c

.PHONY: all clean

all: testHtml

# Executables
testHtml: main.c $(HOEDOWN_SRC)
	$(CC) $^ $(CFLAGS) -o $@

# Generate .o object
$(HOEDOWN_SRC): %.o: %.c
	$(CC) $(CFLAGS) -c $< -o $@

# Clean
clean:
	$(RM) *.o src/*.o