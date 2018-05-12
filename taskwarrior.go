// The MIT License (MIT)
// Copyright (C) 2018 Georgy Komarov <jubnzv@gmail.com>
//
// API bindings to taskwarrior database.

package taskwarrior

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

// Represent a taskwarrior instance
type TaskWarrior struct {
	Config         *TaskRC // Configuration table
	TasksPending   []Task  // Pending tasks
	TasksCompleted []Task  // Completed tasks
}

// Read data file from 'data.location' filepath.
// We are interested in two files in this dir: `completed.data` and `pending.data` that represents data entries
// with format very similar to JSON arrays.
func ReadDataFile(filepath string) (tasks []Task, err error) {
	dataFile, err := os.Open(filepath)
	if err != nil {
		return
	}
	defer dataFile.Close()

	buf, err := ioutil.ReadAll(dataFile)
	if err != nil {
		return
	}

	lines := bytes.Split(buf, []byte{'\n'})
	for _, line := range lines[:len(lines)-1] {
		fmt.Println(string(line))
		bufTask, e := ParseTask(string(line))
		if e != nil {
			return
		}
		tasks = append(tasks, *bufTask)
	}

	return
}

// Create new TaskWarrior instance.
func NewTaskWarrior(configPath string) (*TaskWarrior, error) {
	// Read the configuration file.
	taskRC, err := ParseTaskRC(configPath); if err != nil {
		return nil, err
	}

	// Initialize active tasks.
	tasksPending, err := ReadDataFile(taskRC.DataLocation + "/pending.data"); if err != nil {
		return nil, err
	}

	// Initialize completed tasks.
	tasksCompleted, err := ReadDataFile(taskRC.DataLocation + "/completed.data"); if err != nil {
		return nil, err
	}

	// Create new TaskWarrior instance.
	tw := &TaskWarrior{
		Config:         taskRC,
		TasksPending:   tasksPending,
		TasksCompleted: tasksCompleted,
	}

	return tw, nil
}

// Fetch tasks for current taskwarrior.
func (tw *TaskWarrior) FetchTasks() (tasks []Task) {
	tasks = append(tw.TasksCompleted, tw.TasksPending...)
	return
}
