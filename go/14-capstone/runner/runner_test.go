package runner

import (
	"context"
	"errors"
	"sync/atomic"
	"testing"
	"time"
)

func TestRunner_RunValidatesWorkers(t *testing.T) {
	r := Runner{Workers: 0}
	if err := r.Run(context.Background(), nil); !errors.Is(err, ErrInvalidWorkers) {
		t.Fatalf("got %v; want ErrInvalidWorkers", err)
	}
}

func TestRunner_RunAggregatesErrors(t *testing.T) {
	e1 := errors.New("e1")
	e2 := errors.New("e2")
	tasks := []Task{
		func(context.Context) error { return nil },
		func(context.Context) error { return e1 },
		func(context.Context) error { return e2 },
	}
	r := Runner{Workers: 2}
	err := r.Run(context.Background(), tasks)
	if err == nil {
		t.Fatalf("expected non-nil error")
	}
	if !errors.Is(err, e1) || !errors.Is(err, e2) {
		t.Fatalf("expected joined error to match component errors via errors.Is; got %v", err)
	}
}

func TestRunner_RunStopsSchedulingOnCancel(t *testing.T) {
	var started int32
	block := make(chan struct{})
	defer close(block)

	tasks := make([]Task, 100)
	for i := range tasks {
		tasks[i] = func(ctx context.Context) error {
			atomic.AddInt32(&started, 1)
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-block:
				return nil
			}
		}
	}

	ctx, cancel := context.WithCancel(context.Background())
	cancel() // cancel immediately before Run

	r := Runner{Workers: 4}
	_ = r.Run(ctx, tasks)

	// Should NOT have started all 100 tasks (cancellation must stop scheduling)
	if atomic.LoadInt32(&started) >= 100 {
		t.Fatalf("unexpected: started=%d — cancellation was not respected", started)
	}
}

func TestRunner_RunZeroTasks(t *testing.T) {
	r := Runner{Workers: 1}
	if err := r.Run(context.Background(), nil); err != nil {
		t.Fatalf("got %v; want nil for empty task list", err)
	}
}

func TestRunner_String(t *testing.T) {
	r := Runner{Workers: 3}
	if got := r.String(); got != "Runner(workers=3)" {
		t.Fatalf("got %q; want %q", got, "Runner(workers=3)")
	}
}

func TestRunner_RunDoesNotHang(t *testing.T) {
	// Sanity test: ensure no deadlock via context timeout
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()
	tasks := []Task{
		func(context.Context) error { return nil },
		func(context.Context) error { return nil },
	}
	r := Runner{Workers: 2}
	_ = r.Run(ctx, tasks)
}
