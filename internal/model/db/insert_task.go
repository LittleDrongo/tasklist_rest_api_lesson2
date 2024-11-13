package db

import "database/sql"

func InsertTask(db *sql.DB, t task) error {

	lastId, err := insertIntoTasks(db, t)
	if err != nil {
		return err
	}

	err = insertIntoTags(lastId, db, t)
	if err != nil {
		return err
	}

	return nil
}

func insertIntoTasks(db *sql.DB, t task) (lastId int64, err error) {
	insert := `
	INSERT INTO tasks (text, due) VALUES (?, ?)
	;
	`
	result, err := db.Exec(insert, t.Text, t.Due)
	if err != nil {
		return -1, err
	}

	lastId, err = result.LastInsertId()
	if err != nil {
		return -1, err
	}

	return lastId, nil
}

func insertIntoTags(taskId int64, db *sql.DB, t task) error {
	insert := `
		INSERT INTO tags (task_id, tag) VALUES (?, ?);
		`
	for _, tag := range t.Tags {
		_, err := db.Exec(insert, taskId, tag)
		if err != nil {
			return err
		}

	}
	return nil
}
