GOCMD=go
GOBUILD=$(GOCMD) build
GOGET=$(GOCMD) get

CLOUDBUILD_CONFIG=cloudbuild.yml
SHARED_PROJECT_ID=oa-shared-247623

ACTIVE_PROJECT_ID=$(shell gcloud info | grep -oP "(?<=Project: \[).*(?=\])")

BINARY_PATH=bin/
BINARY_NAME=aqua
BINARY_FULL_PATH=$(BINARY_PATH)$(BINARY_NAME)

all: build
build:
	$(GOGET) -d ./...
	$(MAKE) compile
compile:
	$(GOBUILD) -o $(BINARY_FULL_PATH) main.go func_config.go
gcloudbuild:
ifeq ($(ACTIVE_PROJECT_ID), $(SHARED_PROJECT_ID))
	gcloud builds submit --config=$(CLOUDBUILD_CONFIG)
else
	@echo "gcloud Project ID was not set to $(SHARED_PROJECT_ID). Please change your active gcloud config."
	@echo "Current Project ID: $(ACTIVE_PROJECT_ID)"
endif
