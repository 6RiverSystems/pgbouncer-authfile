# GCE Project ID
GCLOUD_PROJECT ?= plasma-column-128721

# Image Version
# VERSION ?= $(shell cat .version)
VERSION ?= "0.0.2"

# GCR image name
GCR_SCRATCH_NAME = gcr.io/$(GCLOUD_PROJECT)/pgbouncer-authfile:$(VERSION)-scratch
GCR_ALPINE_NAME = gcr.io/$(GCLOUD_PROJECT)/pgbouncer-authfile:$(VERSION)-alpine

all: build

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o pgbouncer-authfile cmd/pgbouncer-authfile/main.go

images: build
	docker build -t $(GCR_SCRATCH_NAME) -f Dockerfile.scratch .
	docker build -t $(GCR_ALPINE_NAME) -f Dockerfile.alpine .

publish:
	docker push $(GCR_SCRATCH_NAME)
	docker push $(GCR_ALPINE_NAME)
