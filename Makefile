.PHONY: sync
sync:
	git submodule update --init --remote --recursive

.PHONY: generate
generate:
	GO111MODULE=on go run cmd/*.go && go test ./providers/...
