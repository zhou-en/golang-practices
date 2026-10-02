package batchjob

import (
	"context"
	"fmt"
	"sync"
)

type Task func(ctx context.Context) (string, error)

type Result struct {
	Value string
	Err   error
}

type job struct {
	index int
	task  Task
}

func RunBatch(ctx context.Context, tasks []Task, maxConcurrency int) []Result {
	fmt.Printf("Hello")
	if len(tasks) == 0 {
		return []Result{}
	}

	if maxConcurrency <= 0 {
		maxConcurrency = 1
	}

	if maxConcurrency > len(tasks) {
		maxConcurrency = len(tasks)
	}

	results := make([]Result, len(tasks))

	jobs := make(chan job, len(tasks))
	for i, task := range tasks {
		jobs <- job{index: i, task: task}
	}
	close(jobs)

	var wg sync.WaitGroup
	for w := 0; w < maxConcurrency; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := range jobs {
				if err := ctx.Err(); err != nil {
					results[j.index] = Result{Err: err}
					continue
				}
				val, err := j.task(ctx)
				results[j.index] = Result{
					Value: val,
					Err:   err,
				}
			}
		}()
	}
	wg.Wait()
	return results
}
