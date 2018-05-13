// The MIT License (MIT)
// Copyright (C) 2018 Georgy Komarov <jubnzv@gmail.com>
//
// Simple example that reads tasks from system install taskwarrior.

package main

import (
	"github.com/jubnzv/go-taskwarrior"
)

func main() {
	// Initialize new TaskWarrior instance.
	tw, err := taskwarrior.NewTaskWarrior("~/.taskrc"); if err != nil {
		panic(err)
	}

	// Get all available tasks
	tw.FetchAllTasks()
	// Now you can access to their JSON entries
	//fmt.Println("Your taskwarrior tasks:\n", tw.Tasks)

	// Pretty print for current taskwarrior instance
	tw.PrintTasks()
}
