prog-name = ./dist/tmai.server.mock

build:
	go build -o $(prog-name) -ldflags="-s -w" -gcflags=all="-l -B -wb=false" .

install:
	go install -o $(prog-name) -ldflags="-s -w" .

clean:
	rm $(prog-name)

run: build
	$(prog-name) tmai

docker-run:
	docker pull ayushjhafm/tmai.server.mock:latest
	docker run -p 3000:3000 ayushjhafm/tmai.server.mock
