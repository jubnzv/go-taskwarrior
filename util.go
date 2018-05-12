// The MIT License (MIT)
// Copyright (C) 2018 Georgy Komarov <jubnzv@gmail.com>

package taskwarrior

import (
	"github.com/mitchellh/mapstructure"
	"regexp"
)

// Parse single task entry from taskwarrior database
// Tasks are represented in their own format, for example:
// [description:"Foo", project:"Bar"]
func ParseTask(line string) (t *Task, err error) {
	values := make(map[string]string)
	re := regexp.MustCompile("(:?[[:alpha:]]+):\"(:?[- a-zA-Z0-9]+)\"")
	match := re.FindAllStringSubmatch(line, -1)
	for _, val := range match {
		values[string(val[1])] = string(val[2])
	}
	mapstructure.Decode(values, &t)
	return
}
