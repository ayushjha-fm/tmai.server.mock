prog-name = ./dist/tmai.server.mock

build:
	go build -o $(prog-name) .

install:
	go install -o $(prog-name) .

clean:
	rm $(prog-name)

run: build
	$(prog-name) tmai

docker-run:
	docker pull ayushjhafm/tmai.server.mock:latest
	docker run -p 3000:3000 ayushjhafm/tmai.server.mock
