default: build

build: generate
	@echo "Building ..."
	@CGO_ENABLE=0 go build -o zddi 
	@echo "Build success"

generate:
	@echo "Generate api"
	@for f in $$(ls *.yaml); do echo "Generate $$f"; openapi-cli-generator generate $$f; done
