

.PHONY: setup-tools adr commit pull-sub

setup-tools:
	@if [ -z `command -v air` ]; then go install github.com/cosmtrek/air@latest ; fi
	@if ! [ -x ${GOPATH}/bin/buf ]; then go install github.com/bufbuild/buf/cmd/buf@latest  ; fi


ADR_COUNT:=$(shell find docs/adr -type f | wc -l | tr -d ' ')
adr:
	npx scaffdog generate ADR --output 'docs/adr' --answer 'number:${ADR_COUNT}'

commit:
	npx git-cz

pull-sub:
	git submodule update --remote

