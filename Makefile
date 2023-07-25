.PHONY: rust

rust:
	cd components/rust/michelia-lib
	cargo build

go:
	make -C components/golang/michelia-lib

clean:
	@rm -rf ./components/rust/michelia-lib/src/proto-rs/v0/*
	@rm -rf ./components/golang/michelia-lib/proto-golang/v0/*
