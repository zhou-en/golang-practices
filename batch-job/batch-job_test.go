package batchjob

import (
	"context"
	"errors"
	"sync/atomic"
	"testing"
	"time"
)

func TestRunBatch(t *testing.T) {
	type testCase struct {
		name           string
		tasks          []Task
		maxConcurrency int
		ctxFn          func() (context.Context, context.CancelFunc)
		validate       func(t *testing.T, results []Result, duration time.Duration)
	}

	tests := []testCase{
		{
			name:           "Empty task list",
			tasks:          []Task{},
			maxConcurrency: 5,
			ctxFn:          func() (context.Context, context.CancelFunc) { return context.WithCancel(context.Background()) },
			validate: func(t *testing.T, results []Result, duration time.Duration) {
				if len(results) != 0 {
					t.Errorf("expected 0 results, got %d", len(results))
				}
			},
		},
		{
			name: "Order preservation and error propagation",
			tasks: []Task{
				func(ctx context.Context) (string, error) {
					time.Sleep(50 * time.Millisecond) // finishes second
					return "task-0", nil
				},
				func(ctx context.Context) (string, error) {
					time.Sleep(10 * time.Millisecond) // finishes first
					return "task-1", nil
				},
				func(ctx context.Context) (string, error) {
					return "", errors.New("task-2-failed")
				},
			},
			maxConcurrency: 3,
			ctxFn:          func() (context.Context, context.CancelFunc) { return context.WithCancel(context.Background()) },
			validate: func(t *testing.T, results []Result, duration time.Duration) {
				if len(results) != 3 {
					t.Fatalf("expected 3 results, got %d", len(results))
				}

				if results[0].Value != "task-0" || results[0].Err != nil {
					t.Errorf("result[0] mismatch: got val=%q, err=%v", results[0].Value, results[0].Err)
				}
				if results[1].Value != "task-1" || results[1].Err != nil {
					t.Errorf("result[1] mismatch: got val=%q, err=%v", results[1].Value, results[1].Err)
				}
				if results[2].Err == nil || results[2].Err.Error() != "task-2-failed" {
					t.Errorf("result[2] error mismatch: got %v", results[2].Err)
				}
			},
		},
		{
			name:           "Enforce maximum concurrency cap",
			maxConcurrency: 2,
			ctxFn:          func() (context.Context, context.CancelFunc) { return context.WithCancel(context.Background()) },
			tasks: func() []Task {
				var activeCount int32
				var peakCount int32
				tasks := make([]Task, 6)

				for i := 0; i < 6; i++ {
					tasks[i] = func(ctx context.Context) (string, error) {
						curr := atomic.AddInt32(&activeCount, 1)
						defer atomic.AddInt32(&activeCount, -1)

						// Track peak concurrency using CAS
						for {
							p := atomic.LoadInt32(&peakCount)
							if curr <= p || atomic.CompareAndSwapInt32(&peakCount, p, curr) {
								break
							}
						}

						time.Sleep(30 * time.Millisecond)
						return "done", nil
					}
				}
				return tasks
			}(),
			validate: func(t *testing.T, results []Result, duration time.Duration) {
				if len(results) != 6 {
					t.Fatalf("expected 6 results, got %d", len(results))
				}
				// 6 tasks with concurrency limit of 2 and 30ms sleep should take at least 90ms total
				if duration < 80*time.Millisecond {
					t.Errorf("batch executed too fast (%v), concurrency cap was likely ignored", duration)
				}
			},
		},
		{
			name:           "Context timeout skips pending tasks",
			maxConcurrency: 1, // Sequential execution to force queuing
			ctxFn: func() (context.Context, context.CancelFunc) {
				return context.WithTimeout(context.Background(), 25*time.Millisecond)
			},
			tasks: []Task{
				func(ctx context.Context) (string, error) {
					time.Sleep(40 * time.Millisecond) // Exceeds timeout during execution
					return "task-0", nil
				},
				func(ctx context.Context) (string, error) {
					return "task-1", nil
				},
			},
			validate: func(t *testing.T, results []Result, duration time.Duration) {
				if len(results) != 2 {
					t.Fatalf("expected 2 results, got %d", len(results))
				}
				// Task 1 should have been skipped due to context timeout
				if !errors.Is(results[1].Err, context.DeadlineExceeded) {
					t.Errorf("expected context.DeadlineExceeded for task[1], got %v", results[1].Err)
				}
			},
		},
		{
			name:           "Invalid maxConcurrency defaults safely to 1",
			maxConcurrency: -5,
			ctxFn:          func() (context.Context, context.CancelFunc) { return context.WithCancel(context.Background()) },
			tasks: []Task{
				func(ctx context.Context) (string, error) { return "ok", nil },
			},
			validate: func(t *testing.T, results []Result, duration time.Duration) {
				if len(results) != 1 || results[0].Value != "ok" {
					t.Errorf("unexpected results for negative concurrency: %v", results)
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := tt.ctxFn()
			defer cancel()

			start := time.Now()
			results := RunBatch(ctx, tt.tasks, tt.maxConcurrency)
			duration := time.Since(start)

			tt.validate(t, results, duration)
		})
	}
}
