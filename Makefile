#plat ?= linux
#plats = linux darwin

all: cmd

define build_cmd
        @echo 'building $(1) ...'
        @GOOS=$(2) GOARCH=amd64 go build -o ./build/cmd ./$(1)
        @echo 'build $(1) done'
endef

cmd:
	$(call build_cmd,cmd,$(plat))

.PHONY: cmd