package main

import (
	"bytes"
	"fmt"
	"github.com/jubnzv/go-taskwarrior"
	"io"
	"os"
	"os/exec"
	"strings"
)

var LOCAL_MAIL = "user@localhost"

func sendMail(report string) {
	c1 := exec.Command(
		"echo", "-e",
		"Content-Type: text/plain; charset=\"utf-8\";\nSubject: Agenda report",
		"\n\n", report)
	c2 := exec.Command("/usr/sbin/sendmail", LOCAL_MAIL)

	r, w := io.Pipe()
	c1.Stdout = w
	c2.Stdin = r

	var b2 bytes.Buffer
	c2.Stdout = &b2

	c1.Start()
	c2.Start()
	c1.Wait()
	w.Close()
	c2.Wait()
	io.Copy(os.Stdout, &b2)
}

func getPendingTasks(ss []taskwarrior.Task, test func(taskwarrior.Task) bool) (ret []string) {
	for _, s := range ss {
		if test(s) {
			project := s.Project
			if len(project) == 0 {
				project = "<no project>"
			}
			entry := fmt.Sprintf("%-12.12s :: %s", project, s.Description)
			fmt.Printf("%+v\n", s)
			ret = append(ret, entry)
		}
	}
	return
}

func main() {
	tw, _ := taskwarrior.NewTaskWarrior("~/.taskrc")
	tw.FetchAllTasks()
	mytest := func(s taskwarrior.Task) bool { return s.Status == "pending" }
	pending := getPendingTasks(tw.Tasks, mytest)

	title := fmt.Sprintf("\nThere are %d pending taskwarrior tasks:\n\n", len(pending))
	tasks_str := strings.Join(pending, "\n")
	report := title + tasks_str

	sendMail(report)
}
