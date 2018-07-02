DCCMD=docker-compose
DCBUILD=$(DCCMD) build
DCUP=$(DCCMD) up
DCDOWN=$(DCCMD) down
WEBAPP_DIR=./webapp

.PHONY: all docker-build docker-up docker-down webapp

all: docker-build
docker-build: webapp
	$(DCBUILD) --no-cache

docker-up:
	$(DCUP) -d

docker-down:
	$(DCDOWN)

webapp:
	make -C $(WEBAPP_DIR)
