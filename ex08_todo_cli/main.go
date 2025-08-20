package main

import (
	"fmt"
	"os"
	"strconv"
)

type Task struct {
	ID    int
	Title string
	Done  bool
}

var tasks []Task
var nextID = 1

func add(title string) {
	t := Task{ID: nextID, Title: title}
	nextID++
	tasks = append(tasks, t)
	fmt.Printf("added %d: %s\n", t.ID, t.Title)
}

func list() {
	if len(tasks) == 0 {
		fmt.Println("(no tasks)")
		return
	}
	for _, t := range tasks {
		status := " "
		if t.Done {
			status = "x"
		}
		fmt.Printf("[%s] %d: %s\n", status, t.ID, t.Title)
	}
}

func done(id int) {
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Done = true
			fmt.Printf("done %d\n", id)
			return
		}
	}
	fmt.Printf("task %d not found\n", id)
}

func del(id int) {
	for i := range tasks {
		if tasks[i].ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			fmt.Printf("deleted %d\n", id)
			return
		}
	}
	fmt.Printf("task %d not found\n", id)
}

func usage() {
	fmt.Println("Usage:")
	fmt.Println("  add <title>")
	fmt.Println("  list")
	fmt.Println("  done <id>")
	fmt.Println("  del <id>")
}

func main() {
	if len(os.Args) < 2 {
		usage()
		return
	}
	cmd := os.Args[1]

	switch cmd {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("missing title")
			return
		}
		title := os.Args[2]
		add(title)
	case "list":
		list()
	case "done":
		if len(os.Args) < 3 {
			fmt.Println("missing id")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("invalid id")
			return
		}
		done(id)
	case "del":
		if len(os.Args) < 3 {
			fmt.Println("missing id")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("invalid id")
			return
		}
		del(id)
	default:
		usage()
	}
}
