build:
	go build -o bin/hexlet-path-size	./cmd/hexlet-path-size

run:
	./bin/hexlet-path-size

lint:
	/Users/lev/go/bin/golangci-lint run	
  
lint-fix:
	/Users/lev/go/bin/golangci-lint run --fix
