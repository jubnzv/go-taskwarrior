# go-taskwarrior

[![Build Status](https://travis-ci.org/jubnzv/go-taskwarrior.svg?branch=master)](https://travis-ci.org/jubnzv/go-taskwarrior)
[![codecov](https://codecov.io/gh/jubnzv/go-taskwarrior/branch/master/graph/badge.svg)](https://codecov.io/gh/jubnzv/go-taskwarrior)
[![GoDoc](https://godoc.org/github.com/jubnzv/go-taskwarrior?status.svg)](https://godoc.org/github.com/jubnzv/go-taskwarrior)

Golang API for [taskwarrior](https://taskwarrior.org/) database.

## Features

* Custom parser for `.taskrc` configuration files
* Read access to taskwarrior database
* Adding/modifying existing tasks

## Quickstart

Read all tasks for current user:

```
tw, err := taskwarrior.NewTaskWarrior("~/.taskrc")
tw.FetchAllTasks()
tw.PrintTasks()
```

To add new task initialize `Task` object with desired values and use:

```
tw.Add(newTask)
tw.Commit() // Save changes
```

For more samples see `examples` directory and package tests.
