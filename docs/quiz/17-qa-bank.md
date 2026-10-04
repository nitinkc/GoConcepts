# 17 · Interview Q&A Bank

Rapid-fire — cover the answer, say it out loud, check yourself. Grouped by how often they actually come up.

## Tier 1 — asked in nearly every Go interview

??? question "Goroutine vs thread?"
    M:N scheduled by the Go runtime onto OS threads. ~2KB growable stacks vs ~1MB fixed. Switching is a few hundred ns, in userspace — no kernel transition. 100k+ goroutines is normal.

??? question "Channels vs mutexes — when?"
    Channels for *ownership transfer / coordination* (producer→consumer, pipelines, signaling). Mutex for *shared state with no clear owner*. Proverb: "don't communicate by sharing memory; share memory by communicating" — but it's a guideline, not law; counters prefer atomic/Mutex.

??? question "Explain `defer`."
    Deferred calls run LIFO when the enclosing function returns. Args evaluated at defer time, not execution time. Used for cleanup (`defer f.Close()`), and `recover()` only works inside a deferred func. Cost: small, used to matter more pre-1.14 open-coded defers.

??? question "Slice vs array internals?"
    Array = fixed value type, copied on assignment. Slice = header `{ptr, len, cap}` over a shared backing array — cheap to pass, but aliasing bugs (`append` reallocation, subslices sharing memory) are the classic trap.

??? question "How does Go GC work?"
    Concurrent tri-color mark-and-sweep, non-generational, non-moving. Sub-ms STW pauses. Tune with `GOGC` (default 100 = collect when heap doubles) or `GOMEMLIMIT` (1.19+). Escape analysis decides stack vs heap allocation — `go build -gcflags=-m` shows it.

??? question "`interface{}` / `any` — what is it?"
    A (type descriptor, data pointer) pair — two words. That's why `nil` interface comparisons bite: an interface holding a typed nil pointer (`(*T)(nil)`) is **not** `nil`. Top interview trap.

??? question "value vs pointer receivers?"
    Pointer receiver: mutation, avoids copy, convention: if one method needs pointer, all use pointer for consistency. Value receiver: small immutable types, goroutine-safe copies.

## Tier 2 — senior-level depth

??? question "Why is `map` not goroutine-safe?"
    Design choice: sync baked in would cost everyone. Concurrent read+write on a map = panic ("concurrent map read and map write") — runtime detects it. Use `sync.RWMutex`+map or `sync.Map` (only for write-once-read-heavy or disjoint key sets).

??? question "Memory model / happens-before?"
    Channel send happens-before the corresponding receive. `sync.Mutex` Unlock happens-before next Lock. Without these, no ordering guarantees — "if it races, anything goes" (like Java's JMM but simpler).

??? question "How would you implement a rate limiter?"
    `time.Ticker` + channel, or `golang.org/x/time/rate` (token bucket, `rate.NewLimiter(rate.Every(100*time.Millisecond), 10)`).

??? question "Graceful shutdown?"
    `signal.NotifyContext(ctx, SIGTERM)` → `srv.Shutdown(ctx)` drains in-flight requests; `errgroup` to stop workers; flush buffers before exit.

??? question "Difference between `new` and `make`?"
    `new(T)` → zeroed `*T`, only allocates. `make` → initialized slices/maps/channels (creates the internal runtime structures). `make` is only for those three types.

??? question "Embedding vs inheritance?"
    Embedding promotes the embedded type's methods — composition satisfying interfaces. No virtual dispatch: `Employee.Greet()` calls `Person.Greet` with the **embedded** Person — the outer type's overrides aren't seen by promoted methods (no polymorphism through embedding).

??? question "What is `iota`?"
    Auto-incrementing constant within a `const` block — enums: `const (Read Perm = 1 << iota; Write; Exec)`.

## Tier 3 — design questions

??? question "Design a worker pool."
    Buffered `jobs chan`, N goroutines `for j := range jobs`, results channel, `close(jobs)` to signal done, `WaitGroup` to join. Mention backpressure: bounded channel = natural backpressure.

??? question "Design a concurrent cache."
    `map` + `RWMutex` (sharded mutexes for scale — e.g., 256 shards by `fnv(key) % 256`), TTL via time heap or per-entry expiry + janitor goroutine, `singleflight` (x/sync) to prevent thundering-herd on miss.

??? question "How do you test concurrent code?"
    `-race` always on; `go test -count=100` to shake flakes; injectable `Clock` interfaces; `synctest` (1.24+ experimental) for deterministic time.

## Go proverbs to drop naturally

- "Clear is better than clever."
- "A little copying is better than a little dependency."
- "Don't panic." (return errors; panic = programmer bug)
- "Make the zero value useful."
