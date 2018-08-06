SHELL := /bin/bash

.PHONY : all
all : clean deps build deploy

.PHONY: deps
deps :
	go get github.com/astaxie/beego; go get github.com/BurntSushi/toml

.PHONY: build
build :
	go build -o goadmin main.go

.PHONY: clean
clean :
	rm -rf /tmp/mktarget && rm -rf ./target

.PHONY: deploy
deploy : clean
	mkdir /tmp/mktarget; cp -rf ./* /tmp/mktarget; rm -rf /tmp/mktarget/src && rm -rf /tmp/mktarget/Makefile && rm -rf /tmp/mktarget/main.go && rm -rf /tmp/mktarget/conf/* && rm -rf /tmp/mktarget/doPages/* && mv /tmp/mktarget ./target

.PHONY: tar
tar :
	cd ./target; tar -zvcf goadmin.tar.gz *

