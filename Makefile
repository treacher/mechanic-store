GO_PACKAGE_PATHS := $(shell glide novendor)

test:
	go test $(GO_PACKAGE_PATHS)
