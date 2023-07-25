.PHONY: all

all:
	protoc -I=../../../protos --go_out=plugins=grpc:. common.proto record.proto element.proto wise.proto vector.proto version.proto

clean:
	@rm -rf ./catfish/*
