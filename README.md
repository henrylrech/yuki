# Yuki ❄️

Yuki is a small Go library for scheduled execution of tasks with support for intervals and custom functions.

It allows you to define a list of tasks with names and associated functions, and execute them sequentially at a fixed time interval.

## ⚙️ Installation

```bash
go get github.com/henrylrech/yuki@v1.0.0
```

## ⚡️ Usage
```go title="Example"
package main

import (
	"fmt"
	"time"

	"github.com/henrylrech/yuki/yuki"
)

func main() {
	tasks := []yuki.Task{
		{
			Name: "Task 1",
			Func: func() {
				fmt.Println("executing task 1")
			},
		},
		{
			Name: "Task 2",
			Func: SomeTask,
		},
	}

	yuki.Execute(tasks, time.Second*20, false)
}

func SomeTask() {
	x := 16 + 47
	fmt.Printf("%d\n", x)
}

```
