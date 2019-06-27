IMPORT?=""

all: test readme

test: generate
	go test ./...

generate:
	rm -f *_generated.go
	go generate
	rename 's/gotemplate_([A-Za-z0-9]+)\.go/\L$$1_generated.go/' gotemplate_*.go

readme:
	godocdown 4d63.com/optf > README.md

setup:
	go get github.com/ncw/gotemplate
	go get github.com/robertkrimen/godocdown/godocdown

replace:
	@set -e; \
	if [ -z "$(IMPORT)" ]; then \
		echo "IMPORT env variable is not set"; \
	else \
		grep -rl "github.com/hqhs/types" . | xargs sed -i '' -e 's:example.com/hqhs/types:$(IMPORT)/:g'; \
	fi
