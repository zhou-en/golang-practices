# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## What this repo is

Personal Go practice code: course exercises (`cards`, `structs`, `map`, `interfaces`, `channels`, `http`, `printFile`, ...) LeetCode problems in `pNNNN_snake_case/` folders (`p0001_two_sum`, `p0032_longest_valid_parentheses`, ...), and `batch-job` (a concurrency exercise, not LeetCode). Each top-level directory is an independent, self-contained program or package. Nothing imports across directories.

## Module layout

One `go.mod` at the root (`github.com/zhou-en/golang-practices`). Each top-level folder is its own package, so it builds and runs on its own. A compile error in one folder does not stop `go run` of another. Do not add a `go.mod` inside a folder: a nested module drops that folder out of the root module.

## Commands

The `Makefile` is the entry point for every project command; `make` lists the targets. `P` takes a LeetCode number (`P=1` finds `p0001_*`) or a folder name (`P=cards`). Without `P`, `test`, `vet`, and `fmt` run on `./...`. Each target prints the `go` command it runs. Add new project commands as Makefile targets with a `## description` comment so `make help` lists them. The Makefile must work with GNU Make 3.81 (macOS default): no nested parentheses inside `$(shell ...)`, and no `#` in variable values.

```sh
make new N=15 T=3sum                                 # create p0015_3sum/ (calls new.sh)
make test P=1                                        # go test ./p0001_two_sum
make test P=batch-job RUN='TestRunBatch/Empty_task_list'   # one subtest; spaces become underscores
make run P=1                                         # go run ./p0001_two_sum
make check                                           # fmt, vet, test on ./...
go test -race ./batch-job                            # no make target for -race
```

`make fmt` and `make check` without `P` rewrite unformatted course files (`area`, `cards`, `http`, `map`). Use `make fmt P=<folder>` to limit the change.

## Adding a new problem

Run `make new N=<number> T=<snake_case_title>` (for example `make new N=15 T=3sum`). It runs `new.sh`, which gets the real title from the LeetCode API and creates `p0015_3sum/` with `description.md`, `main.go`, and a `main_test.go` that fails until test cases are added. The full workflow for the user is in `README.md`. The test pattern to copy is `p0001_two_sum/main_test.go`.

## Conventions

- Each LeetCode folder has a `main_test.go` with the same cases as its `main()`. A failing test usually means the solution is unfinished or wrong, not that the test is wrong. Do not change a solution unless asked: this is practice code.
- `p0032_longest_valid_parentheses/main.go` also holds `isValid`, a solution to #20 Valid Parentheses.
- `batch-job` is a library package (`package batchjob`, no `main`). It runs `[]Task` through a worker pool with a concurrency limit, keeps result order, and honours context cancellation.
- `cards/deck_test.go` writes and removes a temp file `_decktesting` in the package directory (`go test` runs there).
- Problem statements sit next to the code as `description.md` (for example `p0014_longest_common_prefix/description.md`).
- `.gitignore` excludes compiled binaries, including the binary that `go build` creates inside a `pNNNN_*` folder.
