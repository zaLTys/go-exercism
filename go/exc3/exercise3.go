package tasks

//Exercise 3 — Task Processor with Goroutines, Channels and Pointers
//------------------------------------------------------------------
//
//Real‑world scenario:
//--------------------
//Imagine a small background service that processes "tasks" one by one:
//  - Each task has an ID and some payload (string).
//  - When a task is processed, we mark it as Processed = true.
//
//In this exercise you will:
//  1. Create tasks using a constructor that returns a *Task (pointer).
//  2. Mark a task as processed by modifying it through a pointer.
//  3. Process a slice of tasks sequentially (no goroutines).
//  4. Process the same kind of slice concurrently using goroutines and channels.
//
//You will practice:
//  - Working with pointers (*Task).
//  - Sending pointers over channels.
//  - Spinning up simple worker goroutines.
//  - Collecting results from a channel.
//
//Your goals:
//-----------
//Implement the following functions:
//
//    func NewTask(id int, payload string) *Task
//        - Create a Task value.
//        - Set Processed to false.
//        - Return a pointer to it.
//
//    func MarkTaskProcessed(t *Task)
//        - Mark the given task as processed by changing the value via its pointer.
//
//    func ProcessSequential(tasks []Task) []Task
//        - Take a slice of Task values.
//        - For each one, call MarkTaskProcessed on its address.
//        - Return a new slice with all tasks marked as processed.
//        - Preserve the order of tasks.
//
//    func ProcessConcurrently(tasks []Task, workers int) []Task
//        - Use goroutines and channels to process tasks concurrently.
//        - Steps (suggested):
//            1. Create a channel of *Task for sending work (jobs).
//            2. Create a channel of *Task for receiving processed tasks (results).
//            3. Start 'workers' goroutines. Each goroutine should:
//                  - Read tasks from the jobs channel (until it is closed).
//                  - For each task pointer:
//                        * call MarkTaskProcessed(t)
//                        * send the pointer into the results channel
//            4. In the main goroutine:
//                  - Send the address of each element of the tasks slice into jobs.
//                  - Close the jobs channel after sending all tasks.
//                  - Receive exactly len(tasks) results from the results channel.
//                  - Build and return a []Task slice.
//
//        - The result slice must contain all tasks with Processed = true.
//        - The order of tasks in the result does NOT matter in the tests.
//
//Important:
//----------
//- Do NOT use time.Sleep in this exercise.
//- You do NOT need sync.WaitGroup if you:
//    - Close the jobs channel once all tasks are sent.
//    - Receive exactly len(tasks) results from the results channel.
//

// Task represents a unit of work to be processed.
type Task struct {
	ID        int
	Payload   string
	Processed bool
}

// NewTask creates a new Task and returns a pointer to it.
// Processed must be initialized to false.
func NewTask(id int, payload string) *Task {
	return &Task{ID: id, Payload: payload, Processed: false}
}

// MarkTaskProcessed marks the given task as processed using its pointer.
func MarkTaskProcessed(t *Task) {
	t.Processed = true
}

// ProcessSequential processes all tasks one by one (no goroutines).
// It returns a new slice with all tasks marked as processed.
func ProcessSequential(tasks []Task) []Task {
	result := make([]Task, len(tasks))
	copy(result, tasks)

	for i := range result {
		MarkTaskProcessed(&result[i])
	}

	return result
}

// ProcessConcurrently processes all tasks using 'workers' goroutines and channels.
// It returns a slice containing all processed tasks (order does not matter).
func ProcessConcurrently(tasks []Task, workers int) []Task {
	jobs := make(chan *Task)
	results := make(chan *Task)

	// Start workers
	for i := 0; i < workers; i++ {
		go func() {
			for task := range jobs {
				MarkTaskProcessed(task)
				results <- task
			}
		}()
	}

	// Send jobs
	go func() {
		for i := range tasks {
			jobs <- &tasks[i]
		}
		close(jobs)
	}()

	// Collect results
	out := make([]Task, 0, len(tasks))
	for i := 0; i < len(tasks); i++ {
		processed := <-results
		out = append(out, *processed)
	}

	return out
}
