package yuki

import (
	"fmt"
	"time"
)

type Task struct {
	Name string
	Func func()
}

func yuki_log(message string, should_log bool) {
	if should_log {
		fmt.Println("yuki: " + message)
	}
}

func run_tasks(tasks []Task, should_log bool) {
	start := time.Now()
	yuki_log("starting tasks at "+start.Format("15:04:05"), should_log)

	for _, task := range tasks {
		msg := fmt.Sprintf("executing %s", task.Name)
		yuki_log(msg, true)
		task.Func()
	}

	end := time.Now()
	duration := end.Sub(start)
	yuki_log(fmt.Sprintf("ending tasks at %s (duration: %s)", end.Format("15:04:05"), duration), should_log)
}

// Execute runs a given list of tasks sequentially at a fixed time interval.
//
// It starts by executing the tasks immediately in a separate goroutine, then
// continues to execute them periodically based on the provided interval.
//
// Parameters:
//   - tasks: A slice of Task structs, each containing a name and a function to execute.
//   - time_interval: The duration between each execution cycle.
//   - verbose: If true, logs the start and end of each task.
//
// Example:
//
//	Execute(tasks, time.Second*10, true)
func Execute(tasks []Task, time_interval time.Duration, verbose bool) {

	ticker := time.NewTicker(time_interval)
	defer ticker.Stop()

	go run_tasks(tasks, verbose)

	for range ticker.C {
		go run_tasks(tasks, verbose)
	}
}
