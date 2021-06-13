prog-name = ./dist/tmai.server.mock

build:
	go build -o $(prog-name) -ldflags="-s -w" -gcflags=all="-l -B -wb=false" .

install:
	go install -o $(prog-name) -ldflags="-s -w" .

clean:
	rm $(prog-name)

run: build
	$(prog-name) tmai
