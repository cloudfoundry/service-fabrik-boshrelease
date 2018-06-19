#!/bin/bash

BIN = ./bin
API_SERVER_CONFIG = ./apiserver.local.config
ETCD_WORK_DIR = ./default.etcd

all:	build

clean:
	rm -rf ${BIN} ${API_SERVER_CONFIG} ${ETCD_WORK_DIR} kubeconfig

init:
	[ -d ${BIN} ] || mkdir ${BIN}
	[ -d ${API_SERVER_CONFIG} ] || mkdir ${API_SERVER_CONFIG}
	[ -d ${ETCD_WORK_DIR} ] || mkdir ${ETCD_WORK_DIR}

regenerate:	init
	apiserver-boot build generated

build:	build-local

build-local:	init
	go build -o ${BIN}/apiserver cmd/apiserver/main.go
	go build -o ${BIN}/controller-manager cmd/controller-manager/main.go

run-local:	build-local
	apiserver-boot run local --build=false
