SHELL := /bin/bash

.PHONY : all
all : clean deps build deploy

.PHONY: deps
deps :
	go get -u github.com/julienschmidt/httprouter; go get -u github.com/BurntSushi/toml;
	go get -u go.uber.org/zap; go get -u github.com/gobasis/log;
	go get -u github.com/satori/go.uuid; go get -u github.com/gobasis/utils;
	go get -u github.com/gobasis/http; go get -u github.com/go-sql-driver/mysql;

.PHONY : init
init :
	cd ./src/github.com/greatbsky/gomvc/dao/Init; go test -v -count=1 Init_test.go;

.PHONY: test
test :
	cd src/github.com/greatbsky/gomvc/;

.PHONY: build
build :
	go build -o gomvc github.com/greatbsky/gomvc

.PHONY: clean
clean :
	rm -rf /tmp/mktarget && rm -rf ./target

.PHONY: deploy
deploy : clean
	mkdir /tmp/mktarget && mv /tmp/mktarget ./target

.PHONY: tar
tar :
	cd ./target; tar -zvcf gomvc.tar.gz *

