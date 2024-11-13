package db

import (
	"database/sql"
	"fmt"
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
		var due sql.NullTime // Используют так как поддерживают пустые поля.

		if err := rows.Scan(&task.id, &task.Text, &due, &tags); err != nil {
			return nil, err
		}

		task.Due = due.Time

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

func FindTasksByTags(tags ...string) (tasks []Task, err error) {
	if len(tags) == 0 {
		return nil, fmt.Errorf("No tags provided")
	}

	var tagsNew []string

	for _, t := range tags {
		t = "'" + t + "'"
		tagsNew = append(tagsNew, t)
	}

	tagsStr := strings.Join(tagsNew, ", ")

	query := fmt.Sprintf(`
		SELECT 
			tasks.id,
			tasks.text,
			tasks.due,
			IFNULL(GROUP_CONCAT(tags.tag, '; '), '') AS tags
		FROM 
			tasks
		LEFT JOIN 
			tags ON tasks.id = tags.task_id
		WHERE 
			tags.tag IN (%v)
		GROUP BY 
			tasks.id
		HAVING 
			COUNT(DISTINCT tags.tag) = %v;
	`, tagsStr, len(tags))

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
		var tagString sql.NullString
		var due sql.NullTime

		if err := rows.Scan(&task.id, &task.Text, &due, &tagString); err != nil {
			return nil, fmt.Errorf("scan error: %v", err)
		}

		task.Due = due.Time

		if tagString.String != "" {
			task.Tags = strings.Split(tagString.String, ", ")
		}

		tasks = append(tasks, task)
	}
	defer rows.Close()

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return tasks, nil
}
