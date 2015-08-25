default: clean build

build:
	godep restore
	go build -o ./bin/jazz

clean:
	rm -f ./bin/*
