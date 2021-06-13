prog-name = ./dist/tmai.server.mock

build:
	go build -o $(prog-name) -ldflags="-s -w" .

install:
	go install -o $(prog-name) -ldflags="-s -w" .

clean:
	rm $(prog-name)

run: build
	$(prog-name) tmai
