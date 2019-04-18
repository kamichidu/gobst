.PHONY: help
help:  ## show this help
	@grep -E '^[a-zA-Z_\/-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: lint
lint:  ## golangci-lint run
ifeq (, $(shell which golangci-lint 2>/dev/null))
	@echo 'no golangci-lint in $$PATH, installing...'
ifneq (, $(shell which curl 2>/dev/null))
	curl -LJ -o - https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | \
		sed -e 's/msys_nt) os="windows" ;;/msys_nt*) os="windows" ;;/g' | \
		sed -e 's/case $${ARCH} in/case $${OS} in/' | \
		sh -s -- -b $$(go env GOPATH | tr '\\\\' '/' | sed -re 's#^([A-Z]):#/\L\1#')/bin v1.15.0
else ifneq (, $(shell which wget 2>/dev/null))
	wget -O - -q https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | \
		sed -e 's/msys_nt) os="windows" ;;/msys_nt*) os="windows" ;;/g' | \
		sed -e 's/case $${ARCH} in/case $${OS} in/' | \
		sh -s -- -b $$(go env GOPATH | tr '\\\\' '/' | sed -re 's#^([A-Z]):#/\L\1#')/bin v1.15.0
else
	$(error 'no wget or curl in $$PATH')
endif
endif
	golangci-lint run
