GITTAG := $(shell git describe --tags --abbrev=0 2> /dev/null)
GITCOMMIT := $(shell git describe --always)
VERSION := $(or ${GITTAG}, v0.0.0)

.PHONY: scs installer all test clean

scs:
	env GOOS=linux go build -ldflags "-X intel/isecl/sgx-caching-service/version.Version=$(VERSION) -X intel/isecl/sgx-caching-service/version.GitHash=$(GITCOMMIT)" -o out/scs

test:
	go test ./... -coverprofile cover.out
	go tool cover -func cover.out
	go tool cover -html=cover.out -o cover.html

installer: scs
	mkdir -p out/installer
	cp dist/linux/scs.service out/installer/scs.service
	cp dist/linux/install.sh out/installer/install.sh && chmod +x out/installer/install.sh
	cp dist/linux/db_rotation.sql out/installer/db_rotation.sql
	cp out/scs out/installer/scs
	makeself out/installer out/scs-$(VERSION).bin "SGX Caching Service $(VERSION)" ./install.sh
	cp dist/linux/install_pgscsdb.sh out/install_pgscsdb.sh && chmod +x out/install_pgscsdb.sh

all: clean installer test

clean:
	rm -f cover.*
	rm -rf out/
