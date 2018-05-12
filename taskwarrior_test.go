// The MIT License (MIT)
// Copyright (C) 2018 Georgy Komarov <jubnzv@gmail.com>

package taskwarrior

import (
	"testing"
	"os/exec"
	"encoding/json"
)

// Helper that executes `task` with selected config path and return result as new TaskRC instances array.
func UtilTaskCmd(configPath string) ([]Task, error) {
	var out []byte
	if configPath != "" {
		rcOpt := "rc:" + configPath
		out, _ = exec.Command("task", rcOpt, "export").Output();
	} else {
		out, _ = exec.Command("task", "export").Output();
	}

	// Initialize new tasks
	tasks := []Task{}
	err := json.Unmarshal([]byte(out), &tasks); if err != nil {
		return nil, err
	}
	return tasks, nil
}

func TestNewTaskWarrior(t *testing.T) {
	config1 := "./fixtures/taskrc/simple_1"
	taskrc1 := &TaskRC{ConfigPath:config1}
	expected1 := &TaskWarrior{Config:taskrc1}
	result1, err := NewTaskWarrior(config1)
	if err != nil {
		t.Errorf("NewTaskWarrior fails with following error: %s", err)
	}
	if expected1.Config.ConfigPath != result1.Config.ConfigPath {
		t.Errorf("Incorrect taskrc path in NewTaskWarrior: expected '%s' got '%s'",
			expected1.Config.ConfigPath, result1.Config.ConfigPath)
	}

	// Incorrect config path
	config2 := "./fixtures/not_exists/33"
	_, err = NewTaskWarrior(config2)
	if err == nil {
		t.Errorf("NewTaskWarrior works with non-existent config '%s'", config2)
	}
}

func TestTaskWarrior_FetchAllTasks(t *testing.T) {
	// Looks that there are no way for use relative path in .taskrc. So I get real tasks from .taskrc and compare
	// their number.
	config1 := "" // Use default ~/.taskrc
	tasks, _ := UtilTaskCmd(config1)
	tw1, err := NewTaskWarrior(config1)
	if err != nil {
		t.Errorf("NewTaskWarrior fails with following error: %s", err)
	}
	tw1.FetchAllTasks()
	if len(tasks) != len(tw1.Tasks) {
		t.Errorf("FetchAllTasks response mismatch: expect %d for %d",
			len(tasks), len(tw1.Tasks))
	}
}
