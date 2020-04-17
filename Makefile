#plat ?= linux
#plats = linux darwin
#arch ?= amd64
#archs = amd64 armv7

all: cmd

define build_cmd
        @echo 'building $(1) ...'
        @GOOS=$(2) GOARCH=$(3) go build -o ./build/cmd ./$(1)
        @echo 'build $(1) done'
endef

cmd:
	$(call build_cmd,cmd,$(plat), $(arch))

.PHONY: cmd