#plat ?= linux
#plats = linux darwin
#arch ?= amd64
#archs = amd64 armv7

all: ddns

define build_ddns
        @echo 'building $(1) ...'
        @GOOS=$(2) GOARCH=$(3) go build -o ./build/ddns ./$(1)
        @echo 'build $(1) done'
endef

ddns:
	$(call build_ddns,cmd,$(plat), $(arch))

.PHONY: ddns
