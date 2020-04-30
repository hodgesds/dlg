DATE := $(shell date --iso-8601=seconds)

build/dlg: test | vendor
	@go build -o ./build/dlg \
		-ldflags "-X main.BuildDate=${DATE}" \
		./dlg

go.mod:
	@GO111MODULE=on go mod tidy

go.sum: | go.mod
	@GO111MODULE=on go mod verify

vendor: | go.sum
	@GO111MODULE=on go mod vendor 

.PHONEY: install
install: | vendor
	@go install -v -ldflags "-X main.BuildDate=${DATE}" ./dlg

test: | vendor
	@go test -v -race -cover ./...

docs: | build/dlg
	@./build/dlg doc > dlg.md

.PHONEY: clean
clean:
	rm -rf build vendor
