default: build

build: generate
	@echo "Building ..."
	@CGO_ENABLE=0 go build -o zddi
	@CGO_ENABLE=0 go build -o zddi_prompt prompt/main.go
	@echo "Build success"

generate:
	@echo "Generate api"
	@for f in $$(ls *.yaml); do echo "Generate $$f"; openapi-cli-generator generate $$f; done

init:
	go get -u github.com/danielgtaylor/openapi-cli-generator
	#go get -u github.com/go-bindata/go-bindata/...
