SHELL=/bin/bash
SELF = $(MAKE)
include Makefile.helpers
include Makefile.debug


dep:
	go get -u
	go mod tidy

lint:
	golangci-lint run -v ./...

build: lint
	go build


## get some statistics about the repository
repo_stats:
	@echo total commits so far: `git log --oneline | wc -l`
	@echo ======
	@echo total files in repo:
	@git diff --shortstat `git hash-object -t tree /dev/null` | xargs echo
	@echo ======

init: py/install_python_packages build_java

define print_error
echo "====Error===="; \
echo "$(1)"; \
echo "============="
endef

## Release related targets

# checks if the state is dirty first, since a changelog cannot be generated if there are uncommitted changes
changelog:
	@git diff --quiet || ($(call print_error,"The working directory is not clean. Please commit all your changes then run again"); exit 1)
	@echo making changelog
	git-chglog -o CHANGELOG.md --next-tag `./semtag final -s minor -o`
	git add -A
	git commit -m "chore: updating changelog"
	git push

release: changelog
	./semtag final -s minor
