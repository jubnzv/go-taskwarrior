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

package taskwarrior_test

import (
	tw "../go-taskwarrior"
	"testing"
)

func TestParseTask(t *testing.T) {
	var parseTests = []struct {
		in   string
		want *tw.Task
	}{
		{"[description:\"foo\"]", &tw.Task{Description: "foo"}},
		{"[description:\"Write tests\" entry:\"1515066136\" modified:\"1515066136\" project:\"go-taskwarrior\" status:\"pending\" uuid:\"1793f808-0d06-4e9f-95a9-f8ac50ba5c03\"]", &tw.Task{Description: "Write tests", Entry: "1515066136", Modified: "1515066136", Project: "go-taskwarrior", Status: "pending", Uuid: "1793f808-0d06-4e9f-95a9-f8ac50ba5c03"}},
	}

	for _, test := range parseTests {
		got, _ := tw.ParseTask(test.in)
		if *got != *test.want {
			t.Errorf("Error while parsing %s:\n\tgot\t%s\n\twant\t%s\n", test.in, got, test.want)
		} else {
			t.Logf("Parse successful %s:\n\tgot\t%s\n\twant\t%s\n", test.in, got, test.want)
		}
	}
}