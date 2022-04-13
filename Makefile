WORKSPACE:=$(CURDIR)
SHELL:=/bin/bash
APPNAME:=kubecm

.ONESHELL:
.PHONY: help
help:
	@{
		echo "指令列表"
		echo
		echo "  tag                   创建 tag"
		echo "  tag.force             创建 tag (强制，忽略未提交的内容)"
		echo
	}

.ONESHELL:
.PHONY: tag.link
tag.link:
	@{
		mkdir -p scripts
		ln -sf "$${HOME}/workspace/workspace.bin/tag.sh" scripts/tag.sh
	}

.ONESHELL:
.PHONY: tag
tag:
	@{
		bash "$${HOME}/workspace/workspace.bin/tag.sh"
	}

.ONESHELL:
.PHONY: tag.force
tag.force:
	@{
		bash "$${HOME}/workspace/workspace.bin/tag.sh" force
	}

.ONESHELL:
.PHONY: build
build:
	@{
		echo rm -rf "bin/${APPNAME}"
		echo gox -osarch="darwin/amd64" \
		--output="bin/${APPNAME}" \
		-ldflags="-X 'dino/cmd/version.Version=1.0.1' \
		-X 'dino/cmd/version.BuildAt=$$(date)' \
		-X 'dino/cmd/version.OS_Arch=darwin/amd64' \
		-X 'dino/cmd/version.GitHASH=$$(git rev-parse HEAD)'"
	}

.ONESHELL:
.PHONY: install
install:
	@{
		mv "bin/${APPNAME}" "/usr/local/bin/${APPNAME}"
	}
