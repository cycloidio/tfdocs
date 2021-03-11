.PHONY: sync
sync:
	@mkdir -p ./repositories && go run cmd/*.go sync

.PHONY: generate
generate:
	@rm -rf ./providers/* && go run cmd/*.go generate && go test ./providers/...

.PHONY: clean
clean:
	@rm -rf ./repositories/*
