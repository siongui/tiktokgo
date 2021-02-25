# cannot use relative path in GOROOT, otherwise 6g not found. For example,
#   export GOROOT=../go  (=> 6g not found)
# it is also not allowed to use relative path in GOPATH
ifndef GITHUB_ACTIONS
	export GOROOT=$(realpath ../../go)
	export PATH := $(GOROOT)/bin:$(PATH)
endif
#export GOCACHE=off
export GO111MODULE=on

ALL_GO_SOURCES=$(shell /bin/sh -c "find *.go | grep -v _test.go")

userpage: fmt
	@echo "\033[92mTest Download __NEXT_DATA__ url in user page HTML...\033[0m"
	@go test -v $(ALL_GO_SOURCES) userpage_test.go

test: fmt
	@echo "\033[92mTest ...\033[0m"
	@go test -v $(ALL_GO_SOURCES) util_test.go

fmt:
	@echo "\033[92mGo fmt source code...\033[0m"
	@go fmt *.go