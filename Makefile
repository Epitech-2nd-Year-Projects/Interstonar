NAME	=	interstonar

RM		= 	rm -rf

all: $(NAME)

$(NAME):
	go build -buildvcs=false -o $(NAME) ./cmd/interstonar

clean:
	$(RM) $(NAME)

fclean: clean

re: fclean all

test:
	go test ./...