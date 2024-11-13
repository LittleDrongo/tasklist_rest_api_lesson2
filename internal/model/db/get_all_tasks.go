package db

import (
	"database/sql"
	"strings"
)

func GetAllTasks() (tasks []Task, err error) {
	query := `
		SELECT 
			tasks.id,
			tasks.text,
			tasks.due,
			IFNULL(GROUP_CONCAT(tags.tag, '; '), '') AS tags
		FROM 
			tasks
		LEFT JOIN 
			tags ON tasks.id = tags.task_id
		GROUP BY 
			tasks.id;
	`

	db, err := sql.Open(sqlDriver, dataBasePath)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var task Task
		var tags sql.NullString
		var due sql.NullTime

		if err := rows.Scan(&task.id, &task.Text, &due, &tags); err != nil {
			return nil, err
		}

		if tags.Valid {
			task.Tags = strings.Split(tags.String, "; ")
		}

		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}
