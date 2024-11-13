package db

import "time"

type Task struct {
	id   int
	Text string    `json:"text"`
	Tags []string  `json:"tags"`
	Due  time.Time `json:"deadline"` // deadline date
}
