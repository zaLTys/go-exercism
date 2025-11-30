package tasks

import (
    "testing"
)

func TestNewTask(t *testing.T) {
    t.Helper()

    payload := "send welcome email"
    task := NewTask(42, payload)

    if task == nil {
        t.Fatalf("NewTask returned nil")
    }

    if task.ID != 42 {
        t.Errorf("NewTask ID = %d, want %d", task.ID, 42)
    }

    if task.Payload != payload {
        t.Errorf("NewTask Payload = %q, want %q", task.Payload, payload)
    }

    if task.Processed {
        t.Errorf("NewTask Processed = true, want false")
    }
}

func TestMarkTaskProcessed(t *testing.T) {
    task := NewTask(1, "example")

    if task.Processed {
        t.Fatalf("expected Processed to be false initially")
    }

    MarkTaskProcessed(task)

    if !task.Processed {
        t.Fatalf("expected Processed to be true after MarkTaskProcessed")
    }
}

func TestProcessSequential(t *testing.T) {
    tasks := []Task{
        {ID: 1, Payload: "a", Processed: false},
        {ID: 2, Payload: "b", Processed: false},
        {ID: 3, Payload: "c", Processed: false},
    }

    got := ProcessSequential(tasks)

    if len(got) != len(tasks) {
        t.Fatalf("ProcessSequential returned %d tasks, want %d", len(got), len(tasks))
    }

    for i, task := range got {
        if !task.Processed {
            t.Errorf("task at index %d not processed (ID=%d)", i, task.ID)
        }
    }

    // ensure original ordering is preserved
    for i, task := range got {
        if task.ID != tasks[i].ID {
            t.Errorf("order changed at index %d: got ID=%d, want ID=%d", i, task.ID, tasks[i].ID)
        }
    }
}

func TestProcessConcurrentlyAllProcessed(t *testing.T) {
    tasks := []Task{
        {ID: 1, Payload: "t1", Processed: false},
        {ID: 2, Payload: "t2", Processed: false},
        {ID: 3, Payload: "t3", Processed: false},
        {ID: 4, Payload: "t4", Processed: false},
        {ID: 5, Payload: "t5", Processed: false},
    }

    workers := 3
    got := ProcessConcurrently(tasks, workers)

    if len(got) != len(tasks) {
        t.Fatalf("ProcessConcurrently returned %d tasks, want %d", len(got), len(tasks))
    }

    // Build a map of ID -> processed
    processedByID := make(map[int]bool)
    for _, task := range got {
        processedByID[task.ID] = task.Processed
    }

    for _, original := range tasks {
        processed, ok := processedByID[original.ID]
        if !ok {
            t.Errorf("task with ID=%d missing from result", original.ID)
            continue
        }
        if !processed {
            t.Errorf("task with ID=%d not processed", original.ID)
        }
    }
}

func TestProcessConcurrentlyWithSingleWorker(t *testing.T) {
    tasks := []Task{
        {ID: 10, Payload: "x", Processed: false},
        {ID: 20, Payload: "y", Processed: false},
    }

    got := ProcessConcurrently(tasks, 1)

    if len(got) != len(tasks) {
        t.Fatalf("ProcessConcurrently returned %d tasks, want %d", len(got), len(tasks))
    }

    for _, task := range got {
        if !task.Processed {
            t.Errorf("expected all tasks processed, got ID=%d with Processed=false", task.ID)
        }
    }
}
