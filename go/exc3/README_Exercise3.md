# 📘 Exercise 3 — Task Processor with Goroutines, Channels and Pointers

## 🎯 Goal

This exercise slows things down a bit and focuses on **understanding**:

- How to use **pointers** to modify values.
- How to send **pointers over channels**.
- How to spin up **simple worker goroutines**.
- How to collect results from **channels** in a safe, controlled way.

All of this is wrapped in a small, realistic scenario: processing background tasks.

---

## 🧩 Scenario

You are building a tiny background service that processes tasks.

Each **Task** has:

```go
type Task struct {
    ID        int
    Payload   string
    Processed bool
}
```

- `ID` — a unique identifier.
- `Payload` — what the task is about (e.g. "send email", "resize image").
- `Processed` — whether the task has been handled yet.

You will implement four functions in `tasks.go`.

---

## 1️⃣ Step 1 — Creating Tasks with Pointers

Function:

```go
func NewTask(id int, payload string) *Task
```

Requirements:

- Create a `Task` value.
- Set:
  - `ID` = given `id`
  - `Payload` = given `payload`
  - `Processed` = `false`
- Return a **pointer** to the created task (`*Task`).

Example usage:

```go
t := NewTask(1, "send welcome email")
// t is a *Task
```

This is your first small practice with constructors that return pointers.

---

## 2️⃣ Step 2 — Modifying a Task via Pointer

Function:

```go
func MarkTaskProcessed(t *Task)
```

Requirements:

- Use the pointer to change the actual Task it points to.
- Set `t.Processed = true`.
- Optionally, you can also modify `t.Payload` (e.g. add " (done)") — the tests only care about `Processed`.

Example mental model:

```go
t := NewTask(1, "job")
MarkTaskProcessed(t)
// Now t.Processed should be true.
```

This step reinforces how **pointers let you modify values in-place**.

---

## 3️⃣ Step 3 — Sequential Processing (No Goroutines Yet)

Function:

```go
func ProcessSequential(tasks []Task) []Task
```

Requirements:

- Take a slice of `Task` **values** (not pointers).
- For each task:
  - Use its address with `MarkTaskProcessed(&task)` **or**
  - Copy into a new variable and then mark processed.
- Return a **new slice** where all tasks have `Processed = true`.
- Preserve the **original order**.

Suggested approach:

```go
result := make([]Task, len(tasks))
for i, t := range tasks {
    temp := t
    MarkTaskProcessed(&temp)
    result[i] = temp
}
return result
```

Check the tests to see what behavior is expected.

---

## 4️⃣ Step 4 — Concurrent Processing with Goroutines & Channels

Function:

```go
func ProcessConcurrently(tasks []Task, workers int) []Task
```

Now we bring goroutines and channels into the picture.

### High-level idea

You will:

1. Create:
   - a **jobs** channel: `chan *Task`
   - a **results** channel: `chan *Task`
2. Start `workers` goroutines.
   - Each goroutine:
     - Reads `*Task` values from `jobs`.
     - Calls `MarkTaskProcessed` on each pointer.
     - Sends the pointer into `results`.
     - Exits when `jobs` is closed and no more tasks are available.
3. In the main goroutine:
   - Send the address of each element in the `tasks` slice (`&tasks[i]`) into the `jobs` channel.
   - Close the `jobs` channel when done sending.
   - Receive exactly `len(tasks)` results from the `results` channel.
   - Build and return a `[]Task` slice (you can copy the values from the pointers).

### Why we don’t need WaitGroup here

Because you:

- Know **exactly** how many tasks there are: `len(tasks)`.
- Know that each task yields exactly **one** result.
- Close the `jobs` channel when you're done sending work.

So:

- Each worker goroutine will eventually see the channel closed and stop.
- You can read exactly `len(tasks)` values from `results` and then return.

No `time.Sleep`, no busy waiting, and no `WaitGroup` required.

---

## 🧪 What the Tests Check

The tests will verify:

1. `NewTask`
   - Returns a non-nil pointer.
   - Correct ID, Payload, and `Processed == false`.

2. `MarkTaskProcessed`
   - Correctly flips `Processed` to true.

3. `ProcessSequential`
   - Returns same number of tasks.
   - All tasks are processed.
   - Order is preserved.

4. `ProcessConcurrently`
   - Works with multiple workers (e.g., 3).
   - Works with a single worker (workers = 1).
   - All tasks are present and marked as processed.
   - Order is not important here.

---

## 💡 Tips

- Start small: implement `NewTask` and run only the first test.
- Then implement `MarkTaskProcessed`, run tests again.
- Move on to `ProcessSequential`.
- Only then tackle `ProcessConcurrently`.

Try to reason about how data flows:

```text
tasks slice --> jobs channel --> worker goroutines --> results channel --> final slice
```

Each arrow is a chance to practice reading/writing channels and dealing with pointers.

---

When all tests pass, you’ll have a much stronger intuition for:

- when and why to use pointers,
- how to move pointers through channels,
- how to structure simple concurrent workers in Go.
