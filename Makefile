GOPATH=$(shell pwd)/../../../../
export GOPATH

.PHONY : clean

clean :
	rm -Rf $(GOPATH)bin/*
	rm -Rf bin

build :
	go install

docker-clean:
	docker rm greenhome
	docker rmi alivinco/greenhome

dist-docker :
	mkdir -p bin
	env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -v -o bin/greenhome
	echo $(shell ls -a bin/)
	docker build -t alivinco/greenhome .

docker-publish : dist-docker
	docker push alivinco/greenhome
