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

// API bindings to taskwarrior database

package taskwarrior

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

// Default configuration path
var TASKRC = "~/.taskrc"

// Keep configration file values
type TaskRC struct {
	Values map[string]string
}

// Represent a taskwarrior instance
type TaskWarrior struct {
	Config         *TaskRC // Configuration manager
	TasksPending   []Task  // Pending tasks
	TasksCompleted []Task  //Completed tasks
}

// Parse taskwarriror configuration file (default ~/.taskrc)
func ParseConfig(config_path string) (c *TaskRC, err error) {
	c = new(TaskRC)
	c.Values = make(map[string]string)

	// Expand tilda in filepath
	if config_path[:2] == "~/" {
		user, _ := user.Current()
		homedir := user.HomeDir
		config_path = filepath.Join(homedir, config_path[2:])
	}

	// Read the configuration
	file, err := os.Open(config_path)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	buf, err := ioutil.ReadAll(file)
	if err != nil {
		return
	}

	// Traverse line-by-line
	lines := bytes.Split(buf, []byte{'\n'})
	for _, line := range lines {
		line = bytes.TrimSpace(line)
		if len(line) == 0 {
			continue
		}

		parts := bytes.SplitN(line, []byte{'='}, 2)
		parts[0] = bytes.TrimSpace(parts[0])

		// Exclude some patterns
		switch line[0] {
		case '#': // Commented
			continue
		}
		if strings.HasPrefix(string(parts[0]), "include") { // Include additional plugins / themes
			continue
		}

		// Fill the map
		key := string(bytes.ToLower(parts[0]))
		value := ""
		if len(parts) == 2 {
			value = string(bytes.TrimSpace(parts[1]))
		} else {
			value = "true"
		}
		c.Values[key] = value
	}

	return
}

// Custom unmarshaller for task data files
func (t *Task) UnmarshalJSON(buf []byte) (err error) {
	tmp := []interface{}{&t.Description, &t.Entry, &t.Modified, &t.Project, &t.Status, &t.Uuid}
	want_len := len(tmp)
	if err = json.Unmarshal(buf, &tmp); err != nil {
		return
	}
	if g, e := len(tmp), want_len; g != e {
		return fmt.Errorf("wrong number of fields in Task: %d != %d", g, e)
	}

	return
}

// Read data file from 'data.location' filepath
// We are interested in two files in this dir: `completed.data` and `pending.data` that represents data entries as json arrays.
func ReadDataFile(filepath string) (tasks []Task, err error) {
	data_file, err := os.Open(filepath)
	if err != nil {
		return
	}
	defer data_file.Close()

	buf, err := ioutil.ReadAll(data_file)
	if err != nil {
		return
	}

	lines := bytes.Split(buf, []byte{'\n'})
	for _, line := range lines[:len(lines)-1] {
		fmt.Println(string(line))
		buf_task, e := ParseTask(string(line))
		if e != nil {
			return
		}
		tasks = append(tasks, *buf_task)
	}

	return
}

// Create new TaskWarrior instance
func NewTaskWarrior(config_path string) (tw *TaskWarrior, err error) {
	// Read the configuration file
	c, err := ParseConfig(config_path)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Initialize hashmap for active tasks
	tp, err := ReadDataFile(c.Values["data.location"] + "/pending.data")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Initialize hashmap for completed tasks
	tc, err := ReadDataFile(c.Values["data.location"] + "/completed.data")
	if err != nil {
		fmt.Println(err)
		return
	}

	tw = &TaskWarrior{
		Config:         c,
		TasksPending:   tp,
		TasksCompleted: tc,
	}

	return
}

// Fetch tasks for current taskwarrior
func (tw *TaskWarrior) FetchTasks() (tasks []Task) {
	tasks = append(tw.TasksCompleted, tw.TasksPending...)
	return
}
