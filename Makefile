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
		echo "  build                 构建 ${APPNAME}"
		echo "  install               安装 ${APPNAME} 到 /usr/local/bin"
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
		rm -rf "bin/${APPNAME}"
		gox -osarch="darwin/amd64" \
		--output="bin/${APPNAME}" \
		-ldflags="-X 'kubecm/cmd.Version=1.0.0' \
		-X 'kubecm/cmd.BuildAt=$$(date)' \
		-X 'kubecm/cmd.OS_Arch=darwin/amd64' \
		-X 'kubecm/cmd.GitHASH=$$(git rev-parse HEAD)'"
		echo
	}

.ONESHELL:
.PHONY: install
install:
	@{
		mv -f "bin/${APPNAME}" "/usr/local/bin/${APPNAME}"
	}
