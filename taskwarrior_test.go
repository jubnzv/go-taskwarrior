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

// Read configuration file for taskwarrior
func TestParseConfig(t *testing.T) {
	// Try to read something non-existent
	fixture_nonexistent := "./fixtures/19_random/taskrc"
	dummy, err := ParseConfig(fixture_nonexistent)
	if err == nil {
		t.Errorf("ParseCondig reads something from unexistent %s: %v", fixture_nonexistent, dummy)
	}
	t.Logf("Handle unexistent config file correctly. Error message: %s", err)

	// Read correct data
	fixture_1 := "./fixtures/1/taskrc"
	config1, err := ParseConfig(fixture_1)
	if err != nil {
		t.Errorf("ParseCondig can't read %s with eror: %s", fixture_1, err)
	}
	t.Logf("Successfully read %s: %v", fixture_1, config1)

}

// Read sample data that contains json arrays of tasks
func TestReadDataFile(t *testing.T) {
	// Try to read something non-existent
	fixture_unexistent := "./fixtures/19/33_something.data"
	dummy, err := ReadDataFile(fixture_unexistent)
	if err == nil {
		t.Errorf("ReadDataFile reads something from unexistent %s: %v", fixture_unexistent, dummy)
	}
	t.Logf("Handle unexistent data file correctly. Error message: %s", err)

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

// Create new TaskWarrior instance and perform basic checks.
func TestTaskWarriorTasks(t *testing.T) {
	// Try to initialize with non-existent config
	fixture1 := "./fixtures/19_random/taskrc"
	tw1, err := NewTaskWarrior(fixture1)
	if err == nil {
		t.Errorf("Initialize TaskWarrior with unexistent taskrc (%s): %v", fixture1, tw1)
	}
	t.Logf("Handle unexistent config file correctly. Error message: %s", err)

	// Try to initialize with config contains incorrect data paths
	fixture2 := "./fixtures/taskrc/err_paths_1"
	tw2, err := NewTaskWarrior(fixture2)
	if err == nil {
		t.Errorf("Initialize TaskWarrior with incorrect data path (see %s): %v", fixture2, tw2)
	}
	t.Logf("Handle unexistent data path correctly. Error message: %s", err)

	// Try to initialize with config contains incorrect pending.data path
	fixture3 := "./fixtures/taskrc/err_paths_2"
	tw3, err := NewTaskWarrior(fixture3)
	if err == nil {
		t.Errorf("Initialize TaskWarrior with path without pending.data (see %s): %v",
			fixture3, tw3)
	}
	t.Logf("Handle unexistent path without pending.data correctly. Error message: %s", err)

	// Try to initialize with config contains incorrect completed.data path
	fixture4 := "./fixtures/taskrc/err_paths_3"
	tw4, err := NewTaskWarrior(fixture4)
	if err == nil {
		t.Errorf("Initialize TaskWarrior with path without completed.data (see %s): %v",
			fixture4, tw4)
	}
	t.Logf("Handle unexistent path without completed.data correctly. Error message: %s", err)

	// Read correct values
	fixture5 := "./fixtures/2/taskrc"
	lenP := 2 // Desired number of pending tasks
	lenC := 1 // Desired number of completed tasks
	tw5 , _ := NewTaskWarrior(fixture5)

	ptasks := tw5.TasksPending
	if len(ptasks) != lenP {
		t.Errorf("Got %d pending tasks (want %d) in %s", len(ptasks), lenP, "fixtures/2/pending.data")
	} else {
		t.Logf("Found %d pending tasks in %s", len(ptasks), "fixtures/2/pending.data")
	}

	ctasks := tw5.TasksCompleted
	if len(ctasks) != lenC {
		t.Errorf("Got %d pending tasks (want %d) in %s", len(ctasks), lenC, "fixtures/2/completed.data")
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
