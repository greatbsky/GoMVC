SHELL := /bin/bash

.PHONY : all
all : clean deps build deploy

.PHONY: deps
deps :
	go get github.com/julienschmidt/httprouter; go get github.com/BurntSushi/toml;
	go get -u go.uber.org/zap; go get -u github.com/gobasis/log;
	go get github.com/satori/go.uuid;

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

