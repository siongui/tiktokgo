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

default: item

userpage: fmt
	@echo "\033[92mTest Download profile pic and items in user page HTML...\033[0m"
	@go test -v $(ALL_GO_SOURCES) userpage_test.go

item: fmt
	@echo "\033[92mTest Download video url in item page HTML...\033[0m"
	@go test -v $(ALL_GO_SOURCES) item_test.go

example_item: fmt
	@echo "\033[92mRun example of downloading video item by URL...\033[0m"
	@go run example/item.go

example_user: fmt
	@echo "\033[92mRun example of downloading user avatar photo and latest 5 video items...\033[0m"
	@go run example/user.go

test: fmt
	@echo "\033[92mTest ...\033[0m"
	@go test -v $(ALL_GO_SOURCES) util_test.go

fmt:
	@echo "\033[92mGo fmt source code...\033[0m"
	@go fmt *.go
	@go fmt example/*.go
