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
