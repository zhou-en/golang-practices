# Run `make` or `make help` to list the targets.
#
# P is a folder name or a LeetCode number:
#   make test P=p0001_two_sum
#   make test P=1

.DEFAULT_GOAL := help
.PHONY: help new test run vet fmt check

# Shell code that sets $pkg from P: "1" -> ./p0001_two_sum, "cards" -> ./cards, empty -> ./...
RESOLVE = p='$(P)'; \
	if [ -z "$$p" ]; then pkg=./...; \
	elif echo "$$p" | grep -Eq '^[0-9]+$$'; then \
		d=$$(ls -d p$$(printf '%04d' "$$(expr "$$p" + 0)")_* 2>/dev/null | head -1); \
		if [ -z "$$d" ]; then echo "error: no folder for problem $$p" >&2; exit 1; fi; \
		pkg=./$$d; \
	else pkg=./$${p%/}; fi

help: ## List the targets
	@grep -E '^[a-z]+:.*## ' $(MAKEFILE_LIST) | awk -F':.*## ' '{printf "  make %-6s %s\n", $$1, $$2}'
	@echo
	@echo "  P=<folder or number>  choose one problem; without P, test/vet/fmt use every folder"
	@echo "  RUN=<regex>           with test: run only matching tests, for example RUN='TestTwoSum/duplicates'"

new: ## Create a problem folder: make new N=15 T=3sum
	@if [ -z "$(N)" ] && [ -z "$(T)" ]; then \
		echo "usage:   make new N=<number> T=<snake_case_title>"; \
		echo "example: make new N=15 T=3sum   (LeetCode 15. 3Sum -> p0015_3sum/)"; \
	elif [ -z "$(N)" ] || [ -z "$(T)" ]; then \
		echo "error: give both N and T, for example: make new N=15 T=3sum" >&2; exit 1; \
	else ./new.sh "$(N)" "$(T)"; fi

test: ## Run tests: make test [P=15] [RUN=regex]
	@$(RESOLVE); set -x; go test $(if $(RUN),-v -run '$(RUN)') $$pkg

run: ## Run main(): make run P=15
	@if [ -z "$(P)" ]; then \
		echo "usage:   make run P=<number or folder>"; \
		echo "example: make run P=1       (runs main() in p0001_two_sum/)"; \
		echo "example: make run P=cards   (runs main() in cards/)"; \
	else $(RESOLVE); set -x; go run $$pkg; fi

vet: ## Run go vet: make vet [P=15]
	@$(RESOLVE); set -x; go vet $$pkg

fmt: ## Format code with gofmt: make fmt [P=15]
	@$(RESOLVE); set -x; go fmt $$pkg

check: fmt vet test ## Run fmt, vet, and test
