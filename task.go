// The MIT License (MIT)
// Copyright (C) 2018 Georgy Komarov <jubnzv@gmail.com>
//
// Implementation for taskwarrior's Task entries.

package taskwarrior

// Task representation.
type Task struct {
	Id          int32   `json:"id"`
	Description string  `json:"description"`
	Project     string  `json:"project"`
	Status      string  `json:"status"`
	Uuid        string  `json:"uuid"`
	Urgency     float32 `json:"urgency"`
	Priority    string  `json:"priority"`
	Due         string  `json:"due"`
	End         string  `json:"end"`
	Entry       string  `json:"entry"`
	Modified    string  `json:"modified"`
}
