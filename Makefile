# Go Prams
GOCMD=go
GOBUILD=$(GOCMD) build -trimpath
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
CURRENT_VERSION=$(shell git describe --tags --abbrev=0)
BUILD_TARGET="cmd/main.go"
BUILD_PATH="./bin/"
BUILD_BASE_NAME=registry-snitcher

build:
	@$(GOCLEAN) all
	@echo Version:$(CURRENT_VERSION)
	@mkdir -p $(BUILD_PATH)
	@echo "== Build for Windows amd64"
	@GOOS=windows GOARCH=amd64 CGO_ENABLED=0 $(GOBUILD) -o $(BUILD_PATH)$(BUILD_BASE_NAME) -ldflags "-X main.version=$(CURRENT_VERSION)" $(BUILD_TARGET)
	@tar zcvf $(BUILD_PATH)$(BUILD_BASE_NAME)-$(CURRENT_VERSION)-windows-amd64.tar.gz -C $(BUILD_PATH) $(BUILD_BASE_NAME)
	@echo "== Build for OSX amd64"
	@GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 $(GOBUILD) -o $(BUILD_PATH)$(BUILD_BASE_NAME) -ldflags "-X main.version=$(CURRENT_VERSION)" $(BUILD_TARGET)
	@tar zcvf $(BUILD_PATH)$(BUILD_BASE_NAME)-$(CURRENT_VERSION)-darwin-amd64.tar.gz -C $(BUILD_PATH) $(BUILD_BASE_NAME)
	@echo "== Build for OSX arm64"
	@GOOS=darwin GOARCH=arm64 CGO_ENABLED=0 $(GOBUILD) -o $(BUILD_PATH)$(BUILD_BASE_NAME) -ldflags "-X main.version=$(CURRENT_VERSION)" $(BUILD_TARGET)
	@tar zcvf $(BUILD_PATH)$(BUILD_BASE_NAME)-$(CURRENT_VERSION)-darwin-arm64.tar.gz -C $(BUILD_PATH) $(BUILD_BASE_NAME)
	@echo "== Build for Linux amd64"
	@GOOS=linux GOARCH=amd64 CGO_ENABLED=0 $(GOBUILD) -o $(BUILD_PATH)$(BUILD_BASE_NAME) -ldflags "-X main.version=$(CURRENT_VERSION)" $(BUILD_TARGET)
	@tar zcvf $(BUILD_PATH)$(BUILD_BASE_NAME)-$(CURRENT_VERSION)-linux-amd64.tar.gz -C $(BUILD_PATH) $(BUILD_BASE_NAME)
	@echo "== Build for Linux arm64"
	@GOOS=linux GOARCH=arm64 CGO_ENABLED=0 $(GOBUILD) -o $(BUILD_PATH)$(BUILD_BASE_NAME) -ldflags "-X main.version=$(CURRENT_VERSION)" $(BUILD_TARGET)
	@tar zcvf $(BUILD_PATH)$(BUILD_BASE_NAME)-$(CURRENT_VERSION)-linux-arm64.tar.gz -C $(BUILD_PATH) $(BUILD_BASE_NAME)
	@rm $(BUILD_PATH)$(BUILD_BASE_NAME)

docker-build:
	@echo Version:$(CURRENT_VERSION)
	@echo "== Build for Linux amd64"
	@docker buildx build --platform linux/amd64 --build-arg BIN_FILENAME=$(BUILD_BASE_NAME)-$(CURRENT_VERSION)-linux-amd64.tar.gz --no-cache=true --push -t docker.io/fideltak/$(BUILD_BASE_NAME):$(CURRENT_VERSION) .
	@echo Version:$(CURRENT_VERSION)
	@echo "== Build for Linux arm64"
	@docker buildx build --platform linux/arm64 --build-arg BIN_FILENAME=$(BUILD_BASE_NAME)-$(CURRENT_VERSION)-linux-arm64.tar.gz --no-cache=true --push -t docker.io/fideltak/$(BUILD_BASE_NAME):$(CURRENT_VERSION) .