GO_PACKAGE_PATHS := $(shell glide novendor)

tests:
	go test $(GO_PACKAGE_PATHS)
