CMDS = ssl-sim-client ssl-sim-client-team
DOCKER_TARGETS = $(addprefix docker-, $(CMDS))
.PHONY: all docker test install proto $(DOCKER_TARGETS)

all: install docker

docker: $(DOCKER_TARGETS)

$(DOCKER_TARGETS): docker-%:
	docker build --build-arg cmd=$* -t $*:latest .

test:
	go test ./...

install:
	go install -v ./...

proto:
	tools/generateProto.sh

update-go:
	go get -v -u all

update: update-go proto
