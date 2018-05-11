//The MIT License (MIT)
//Copyright (C) 2018 Georgy Komarov <jubnzv@gmail.com>

//Permission is hereby granted, free of charge, to any person obtaining a copy
//of this software and associated documentation files (the "Software"), to deal
//in the Software without restriction, including without limitation the rights
//to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
//copies of the Software, and to permit persons to whom the Software is
//furnished to do so, subject to the following conditions:

//The above copyright notice and this permission notice shall be included in all
//copies or substantial portions of the Software.

//THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
//EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
//MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
//IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM,
//DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR
//OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE
//OR OTHER DEALINGS IN THE SOFTWARE.

package taskwarrior

import (
	"testing"
)

// Read sample data that contains json arrays of tasks
func TestReadDataFile(t *testing.T) {
	fixture_pending := "./fixtures/2/pending.data"
	ptasks, err := ReadDataFile(fixture_pending)
	if err != nil {
		t.Errorf("Error while reading %s: %s", fixture_pending, err)
	}
	t.Logf("Found %d pending tasks in %s", len(ptasks), fixture_pending)
}

// Create taskwarrior with custom config and test it values
func TestTaskWarriorConfiguration(t *testing.T) {
	fixture_1 := "./fixtures/1/taskrc"
	warrior, _ := NewTaskWarrior(fixture_1)
	config := warrior.Config
	config_values := config.Values

	t.Logf("%s configuration values: %s", fixture_1, config_values)

	if config_values["data.location"] != "./fixtures/1/" {
		t.Errorf("Incorrect data.location value from %s: got %s, want %s", fixture_1, config_values["data.location"], "./fixtures/1/")
	}

	if config_values["confirmation"] != "no" {
		t.Errorf("Incorrect data.location value from %s: got %s, want %s", fixture_1, config_values["confirmation"], "no")
	}

	if config_values["weekstart"] != "monday" {
		t.Errorf("Incorrect data.location value from %s: got %s, want %s", fixture_1, config_values["weekstart"], "monday")
	}
}

// Check tasks for created taskwarrior
func TestTaskWarriorTasks(t *testing.T) {
	fixture_2 := "./fixtures/2/taskrc"
	len_p_2 := 2 // Desired number of pending tasks
	len_c_2 := 1 // Desired number of completed tasks
	warrior, _ := NewTaskWarrior(fixture_2)

	ptasks := warrior.TasksPending
	if len(ptasks) != len_p_2 {
		t.Errorf("Got %d pending tasks (want %d) in %s", len(ptasks), len_p_2, "fixtures/2/pending.data")
	} else {
		t.Logf("Found %d pending tasks in %s", len(ptasks), "fixtures/2/pending.data")
	}

	ctasks := warrior.TasksCompleted
	if len(ctasks) != len_c_2 {
		t.Errorf("Got %d pending tasks (want %d) in %s", len(ctasks), len_c_2, "fixtures/2/completed.data")
	} else {
		t.Logf("Found %d completed tasks in %s", len(ctasks), "fixtures/2/completed.data")
	}
}

// Try to fetch all tasks
func TestFetchTasks(t *testing.T) {
	fixture_2 := "./fixtures/2/taskrc"
	len_2 := 3 // Desired number of tasks
	warrior, _ := NewTaskWarrior(fixture_2)

	tasks := warrior.FetchTasks()
	if len(tasks) != len_2 {
		t.Errorf("Got %d tasks (want %d) from %s", len(tasks), len_2, fixture_2)
	} else {
		t.Logf("Successfully fetch %d tasks from %s", len(tasks), fixture_2)
	}
}
