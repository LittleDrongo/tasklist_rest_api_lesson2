package db

import (
	"database/sql"
	"fmt"
	"strings"
	"time"
)

func FindTaskById(id int) (task Task, err error, rowExists bool) {

	template := `SELECT 
    tasks.id,
    tasks.text,
    tasks.due,
    GROUP_CONCAT(tags.tag, '; ') AS tags
FROM 
    tasks
LEFT JOIN 
    tags ON tasks.id = tags.task_id
WHERE
    tasks.id = %v
GROUP BY 
    tasks.id;`

	db, err := sql.Open(sqlDriver, dataBasePath)
	if err != nil {
		return task, err, rowExists
	}
	defer db.Close()

	query := fmt.Sprintf(template, id)

	rows, err := db.Query(query)
	if err != nil {
		return task, err, rowExists
	}
	defer rows.Close()

	for rows.Next() {
		rowExists = true
		var (
			id      int
			text    string
			tagsStr string
			due     time.Time
		)
		// use pointers to get data
		err = rows.Scan(&id, &text, &due, &tagsStr)
		if err != nil {
			return task, err, rowExists
		}

		task.id = id
		task.Text = text
		task.Tags = strings.Split(tagsStr, "; ")
		task.Due = due
	}

	err = rows.Err()
	if err != nil {
		return task, err, false
	}

	return task, nil, rowExists
}
