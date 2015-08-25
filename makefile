default: clean build

build:
	godep restore
	go build -o ./bin/goclibase

clean:
	rm -f ./bin/*
