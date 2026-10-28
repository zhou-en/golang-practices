# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## What this repo is

Personal Go practice code: course exercises (`cards`, `structs`, `map`, `interfaces`, `channels`, `http`, `printFile`, ...) LeetCode problems in `pNNNN_snake_case/` folders (`p0001_two_sum`, `p0032_longest_valid_parentheses`, ...), and `batch-job` (a concurrency exercise, not LeetCode). Each top-level directory is an independent, self-contained program or package. Nothing imports across directories.

## Module layout

One `go.mod` at the root (`github.com/zhou-en/golang-practices`). Each top-level folder is its own package, so it builds and runs on its own. A compile error in one folder does not stop `go run` of another. Do not add a `go.mod` inside a folder: a nested module drops that folder out of the root module.

```sh
go run ./p0001_two_sum               # run one problem (from the repo root)
cd p0001_two_sum && go run .         # same, from inside the folder; `go run main.go` misses other files
go test ./p0001_two_sum              # test one folder
go test -race ./batch-job
go test -run 'TestRunBatch/Empty_task_list' -v ./batch-job   # one subtest; spaces become underscores
go vet ./... && go test ./...        # whole repo
```

## Adding a new problem

1. Create a folder named `p` + the 4-digit LeetCode number + `_` + the snake_case title, for example `p0015_3sum/`.
2. Put `package main`, the solution, and a `main()` in `main.go`.
3. Put a table-driven test in `main_test.go` (`package main`, `t.Run` per case). Copy the pattern from `p0001_two_sum/main_test.go`.
4. Optional: save the problem statement as `description.md`.
5. Run `go run ./p0015_3sum` and `go test ./p0015_3sum`.

## Conventions

- Each LeetCode folder has a `main_test.go` with the same cases as its `main()`. A failing test usually means the solution is unfinished or wrong, not that the test is wrong. Do not change a solution unless asked: this is practice code.
- `p0032_longest_valid_parentheses/main.go` also holds `isValid`, a solution to #20 Valid Parentheses.
- `batch-job` is a library package (`package batchjob`, no `main`). It runs `[]Task` through a worker pool with a concurrency limit, keeps result order, and honours context cancellation.
- `cards/deck_test.go` writes and removes a temp file `_decktesting` in the package directory (`go test` runs there).
- Problem statements sit next to the code as `description.md` (for example `p0014_longest_common_prefix/description.md`).
- `.gitignore` excludes compiled binaries, including the binary that `go build` creates inside a `pNNNN_*` folder.
