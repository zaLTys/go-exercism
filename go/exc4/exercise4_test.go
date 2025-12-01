package concurpatterns

import (
    "context"
    "errors"
    "testing"
    "time"
)

func successSource(msg string, delay time.Duration) Source {
    return func(ctx context.Context) (string, error) {
        select {
        case <-ctx.Done():
            return "", ctx.Err()
        case <-time.After(delay):
            return msg, nil
        }
    }
}

func errorSource(err error, delay time.Duration) Source {
    return func(ctx context.Context) (string, error) {
        select {
        case <-ctx.Done():
            return "", ctx.Err()
        case <-time.After(delay):
            return "", err
        }
    }
}

func TestFetchAllSuccess(t *testing.T) {
    ctx := context.Background()

    sources := []Source{
        successSource("A", 10*time.Millisecond),
        successSource("B", 5*time.Millisecond),
        successSource("C", 1*time.Millisecond),
    }

    got, err := FetchAll(ctx, sources)
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }

    if len(got) != 3 {
        t.Fatalf("expected 3 results, got %d", len(got))
    }
}

func TestFetchAllError(t *testing.T) {
    ctx := context.Background()

    e := errors.New("boom")

    sources := []Source{
        successSource("A", 5*time.Millisecond),
        errorSource(e, 3*time.Millisecond),
        successSource("B", 10*time.Millisecond),
    }

    _, err := FetchAll(ctx, sources)
    if err == nil {
        t.Fatal("expected error, got nil")
    }
}

func TestFetchAllContextCancelled(t *testing.T) {
    ctx, cancel := context.WithCancel(context.Background())
    cancel()

    sources := []Source{
        successSource("A", 5*time.Millisecond),
        successSource("B", 5*time.Millisecond),
    }

    _, err := FetchAll(ctx, sources)
    if err == nil {
        t.Fatal("expected error due to cancellation")
    }
}
