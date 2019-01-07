PACKAGES = $(shell go list ./...)

.PHONY: default
default: build

.PHONY: install
install: build
	go install

.PHONY: build
build: check quick

.PHONY: quick
quick:
	go build

.PHONY: check
check: vet lint errcheck interfacer aligncheck structcheck varcheck unconvert staticcheck vendorcheck test

.PHONY: vet
vet:
	go vet $(PACKAGES)

.PHONY: lint
lint:
	golint -set_exit_status $(PACKAGES)

.PHONY: errcheck
errcheck:
	errcheck -exclude errcheck_excludes.txt $(PACKAGES)

.PHONY: interfacer
interfacer:
	interfacer $(PACKAGES)

.PHONY: aligncheck
aligncheck:
	aligncheck $(PACKAGES)

.PHONY: structcheck
structcheck:
	structcheck $(PACKAGES)

.PHONY: varcheck
varcheck:
	varcheck $(PACKAGES)

.PHONY: unconvert
unconvert:
	unconvert -v $(PACKAGES)

.PHONY: staticcheck
staticcheck:
	staticcheck $(PACKAGES)

.PHONY: vendorcheck
vendorcheck:
	vendorcheck $(PACKAGES)
	vendorcheck -u $(PACKAGES)

.PHONY: test
test:
	go test -cover $(PACKAGES)

.PHONY: coverage
coverage:
	echo "mode: count" > coverage-all.out
	$(foreach pkg,$(PACKAGES),\
		go test -coverprofile=coverage.out -covermode=count $(pkg);\
		tail -n +2 coverage.out >> coverage-all.out;)
	go tool cover -html=coverage-all.out

.PHONY: clean
clean:
	go clean

.PHONY: deps
deps:
	go get -u github.com/FiloSottile/vendorcheck
	go get -u github.com/golang/dep/cmd/dep
	go get -u github.com/golang/lint/golint
	go get -u github.com/kisielk/errcheck
	go get -u github.com/mdempsky/unconvert
	go get -u github.com/opennota/check/...
	go get -u honnef.co/go/tools/...
	go get -u mvdan.cc/interfacer
