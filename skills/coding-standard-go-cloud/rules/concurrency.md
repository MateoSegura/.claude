# Concurrency Rules (GO-CON-*)

Go's concurrency model is built on goroutines and channels. These rules ensure safe and effective concurrent programming.

## Philosophy

> **Do not communicate by sharing memory; instead, share memory by communicating.**

## Channels vs Mutexes

| Use Channels When | Use Mutexes When |
|-------------------|------------------|
| Passing ownership of data | Protecting internal state |
| Distributing work | Caching |
| Communicating async results | Simple counters |

---

## GO-CON-001: Always Pass Context to Long-Running Operations :red_circle:

**Tier**: Critical

**Rationale**: Context enables cancellation, timeouts, and deadline propagation.

```go
// Correct - Accept and check context
func (s *Service) Process(ctx context.Context, items []Item) error {
    for _, item := range items {
        select {
        case <-ctx.Done():
            return ctx.Err()
        default:
        }

        if err := s.processItem(ctx, item); err != nil {
            return err
        }
    }
    return nil
}

// Incorrect - No context, can't cancel
func (s *Service) Process(items []Item) error {
    for _, item := range items {
        if err := s.processItem(item); err != nil {
            return err
        }
    }
    return nil
}
```

---

## GO-CON-002: Never Store Context in Structs :red_circle:

**Tier**: Critical

**Rationale**: Context is request-scoped. Storing it in structs leads to stale contexts and unclear ownership.

```go
// Correct - Context as first parameter
type Server struct {
    db *sql.DB
}

func (s *Server) GetUser(ctx context.Context, id string) (*User, error) {
    row := s.db.QueryRowContext(ctx, "SELECT ...", id)
    // ...
}

// Incorrect - Context stored in struct
type Server struct {
    ctx context.Context  // Don't do this
    db  *sql.DB
}
```

---

## GO-CON-003: Ensure Goroutines Can Exit :red_circle:

**Tier**: Critical

**Rationale**: Leaked goroutines consume resources indefinitely. Every goroutine must have a clear exit path.

```go
// Correct - Goroutine exits on context cancellation
func (s *Service) StartWorker(ctx context.Context) {
    go func() {
        ticker := time.NewTicker(time.Second)
        defer ticker.Stop()

        for {
            select {
            case <-ctx.Done():
                return  // Clean exit
            case <-ticker.C:
                s.doWork()
            }
        }
    }()
}

// Incorrect - Goroutine can never exit
func (s *Service) StartWorker() {
    go func() {
        for {
            time.Sleep(time.Second)
            s.doWork()
        }
    }()
}
```

---

## GO-CON-004: Use sync.WaitGroup for Goroutine Coordination :yellow_circle:

**Tier**: Required

**Rationale**: WaitGroup provides a clean way to wait for multiple goroutines. Always call `Add` before starting the goroutine.

```go
// Correct - WaitGroup with Add before goroutine
func ProcessAll(ctx context.Context, items []Item) error {
    var wg sync.WaitGroup
    errCh := make(chan error, len(items))

    for _, item := range items {
        wg.Add(1)  // Add before starting goroutine
        go func(item Item) {
            defer wg.Done()
            if err := process(ctx, item); err != nil {
                errCh <- err
            }
        }(item)
    }

    wg.Wait()
    close(errCh)

    for err := range errCh {
        if err != nil {
            return err
        }
    }
    return nil
}

// Incorrect - Add inside goroutine (race condition)
for _, item := range items {
    go func(item Item) {
        wg.Add(1)  // Race condition!
        defer wg.Done()
        process(item)
    }(item)
}
```

---

## GO-CON-005: Prefer Mutex for Simple State Protection :green_circle:

**Tier**: Recommended

**Rationale**: Channels are for communication; mutexes are for protecting state.

```go
// Correct - Mutex for state protection
type Counter struct {
    mu    sync.Mutex
    count int
}

func (c *Counter) Inc() {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.count++
}

func (c *Counter) Value() int {
    c.mu.Lock()
    defer c.mu.Unlock()
    return c.count
}
```

---

## GO-CON-006: Use errgroup for Concurrent Operations with Error Handling :yellow_circle:

**Tier**: Required

**Rationale**: `errgroup` from `golang.org/x/sync/errgroup` provides a cleaner pattern than manual `sync.WaitGroup` with error channels.

```go
// Correct - errgroup for concurrent operations with errors
import "golang.org/x/sync/errgroup"

func FetchAllData(ctx context.Context, urls []string) ([][]byte, error) {
    g, ctx := errgroup.WithContext(ctx)
    results := make([][]byte, len(urls))

    for i, url := range urls {
        i, url := i, url  // Capture loop variables (not needed in Go 1.22+)
        g.Go(func() error {
            data, err := fetchURL(ctx, url)
            if err != nil {
                return fmt.Errorf("fetching %s: %w", url, err)
            }
            results[i] = data
            return nil
        })
    }

    if err := g.Wait(); err != nil {
        return nil, err
    }
    return results, nil
}

// Correct - errgroup with concurrency limit
func ProcessWithLimit(ctx context.Context, items []Item) error {
    g, ctx := errgroup.WithContext(ctx)
    g.SetLimit(10)  // Max 10 concurrent goroutines

    for _, item := range items {
        item := item
        g.Go(func() error {
            return processItem(ctx, item)
        })
    }

    return g.Wait()
}
```
