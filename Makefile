CC=go build
CFLAGS=-ldflags "-s -w"
EXEC=uuid-bin
SRC=main.go

all: prod

prod: $(SRC)
	$(CC) $(CFLAGS) -o $(EXEC) $<
clean:
	@rm $(EXEC)