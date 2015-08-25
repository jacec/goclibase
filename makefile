default: clean build

build: check
	godep restore
	go build -o ./bin/goclibase

clean:
	rm -f ./bin/*

check:
	./check-dep.sh
