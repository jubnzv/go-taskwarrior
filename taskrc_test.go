// The MIT License (MIT)
// Copyright (C) 2018 Georgy Komarov <jubnzv@gmail.com>

package taskwarrior

import (
	"os/user"
	"testing"
)

func TestPathExpandTilda(t *testing.T) {
	userDir, _ := user.Current()
	expectOrig := userDir.HomeDir
	result := PathExpandTilda("~/")
	if result != expectOrig {
		t.Errorf("ExpectOrig: %s\tGot: %s", expectOrig, result)
	}

	expect1 := expectOrig + "/something/else"
	result = PathExpandTilda("~/something/else")
	if result != expect1 {
		t.Errorf("Expect1: %s\tGot: %s", expect1, result)
	}

	expect2 := expectOrig + "/something/else"
	result = PathExpandTilda("~/something/else/")
	if result != expect2 {
		t.Errorf("Expect2: %s\tGot: %s", expect2, result)
	}
}

func TestGetAvailableKeys(t *testing.T) {
	expected := []string{"DataLocation"}
	result := GetAvailableKeys()
	found := false
	for _, vE := range expected {
		found = false
		for _, vR := range result {
			if vR == vE {
				found = true
				break
			}
		}
		if found == false {
			t.Errorf("Key %s can not be found in avaialable keys!", vE)
		}
	}
}

func TestStripComments(t *testing.T) {
	orig1 := "Foo#123"
	expected1 := "Foo"
	result1 := StripComments(orig1)
	if result1 != expected1 {
		t.Errorf("Incorrect strip comment: expected '%s' got '%s'", expected1, result1)
	}

	orig2 := "#123Foo"
	expected2 := ""
	result2 := StripComments(orig2)
	if result2 != expected2 {
		t.Errorf("Incorrect strip comment: expected '%s' got '%s'", expected2, result2)
	}
}

func TestTaskRC_MapTaskRC(t *testing.T) {
	// Simple configuration entry
	orig1 := "data.location=/home/tester/data"
	expected1 := "/home/tester/data"
	taskrc1 := &TaskRC{DataLocation: orig1}
	taskrc1.MapTaskRC(orig1)
	if taskrc1.DataLocation != expected1 {
		t.Errorf("Incorrect map for DataLocation: expected '%s' got '%s'", expected1, taskrc1.DataLocation)
	}
}

func TestParseTaskRC(t *testing.T) {
	// Simple configuration file
	config1 := "./fixtures/taskrc/simple_1"
	orig1 := "./fixtures/data_1"
	expected1 := &TaskRC{DataLocation: orig1}
	result1, err := ParseTaskRC(config1)
	if err != nil {
		t.Errorf("Can't parse configuration file %s with following error: %v",
			config1, err)
	}
	if expected1.DataLocation != result1.DataLocation {
		t.Errorf("There are some problems to set DataLocation: expected '%s' got '%s'",
			expected1.DataLocation, result1.DataLocation)
	}

	// Use default config location (~/.taskrc)
	config2 := ""
	result2, err := ParseTaskRC(config2)
	if err != nil {
		t.Errorf("Can't parse configuration file %s with following error: %v",
			result2.ConfigPath, err)
	}
	if result2.ConfigPath != TASKRC {
		t.Errorf("There are some problems to set ConfigPath: expected '%s' got '%s'",
			result2.ConfigPath, TASKRC)
	}

	// Incorrect permissions
	config3 := "./fixtures/taskrc/err_permissions_1"
	_, err = ParseTaskRC(config3)
	if err == nil {
		t.Errorf("Read configuration file '%s' content without permissions?", config3)
	}
}
