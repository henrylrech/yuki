package main

import (
	"fmt"
	"time"

	"github.com/henrylrech/yuki/yuki"
)

func tarefa1() {
	fmt.Println("tarefa 1")
}

func tarefa2() {
	fmt.Println("tarefa 2")
}

func tarefa3() {
	fmt.Println("tarefa 3")
}

func main() {
	tasks := []yuki.Task{
		{Name: "Tarefa 1", Func: tarefa1},
		{Name: "Tarefa 2", Func: tarefa2},
		{Name: "Tarefa 3", Func: tarefa3},
	}

	yuki.Execute(tasks, 5*time.Second, false)
}
