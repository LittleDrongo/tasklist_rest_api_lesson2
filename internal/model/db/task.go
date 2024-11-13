package db

import "time"

type task struct {
	id   int
	Text string
	Tags []string
	Due  time.Time // deadline date
}
