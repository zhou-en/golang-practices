# Go Practices

Go practice: LeetCode problems and Go course exercises. Each top-level folder is one program. The repo is one Go module, so you run every command from the repo root.

## Practice a LeetCode problem

1. Find a problem on [LeetCode](https://leetcode.com/problemset/). Note its number and title, for example `15. 3Sum`.
2. Create the folder and template files:

   ```sh
   make new N=15 T=3sum
   ```

   The script creates `p0015_3sum/` with three files:

   | File | What to put in it |
   |---|---|
   | `description.md` | The problem statement. The script gets the title from LeetCode. If the script prints a warning, fix the title and link by hand. |
   | `main.go` | The LeetCode Go starter function, then your solution. |
   | `main_test.go` | One test case per LeetCode example. |

3. Copy the problem statement from LeetCode into `description.md`.
4. On LeetCode, select **Go** as the language. Copy the starter function into `main.go`.
5. In `main_test.go`, add `input` and `want` fields to the struct. Add one case for each LeetCode example. Change the commented call to use your function. `p0001_two_sum/main_test.go` is a complete example.
6. Run the tests. They fail until the solution is correct:

   ```sh
   make test P=15
   ```

7. Solve the problem in `main.go`. Run the tests again until they pass.
8. Optional: paste your solution into LeetCode and submit it to test the hidden cases.
9. Commit:

   ```sh
   git add p0015_3sum && git commit -m "Solve 15. 3Sum" && git push
   ```

## Commands

Run `make` to list the targets. `P` is a problem number or a folder name.

```sh
make new N=15 T=3sum                       # create p0015_3sum/
make test P=1                              # test one problem (same as P=p0001_two_sum)
make test P=1 RUN='TestTwoSum/duplicates'  # one test case (spaces become _)
make run P=1                               # run main()
make test P=cards                          # course folders use the folder name
make test                                  # test every folder
make vet                                   # go vet; add P= for one folder
make fmt P=1                               # gofmt one folder
make check                                 # fmt, vet, and test for every folder
```

Each target prints the `go` command it runs. To run `go` directly from inside a folder, use `go run .`. The command `go run main.go` fails when the folder has more than one `.go` file.

## Folder names

LeetCode folders are `p` + the 4-digit problem number + `_` + the title in snake_case, for example `p0032_longest_valid_parentheses`. The other folders (`cards`, `structs`, `channels`, `batch-job`, ...) are course exercises.
