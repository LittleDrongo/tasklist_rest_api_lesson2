package db

import "database/sql"

func FindTaskById(id int) (t Task, err error) {

	db, err := sql.Open(sqlDriver, "data/db/tasklist.db")
	if err != nil {
		return -1, err
	}
	defer db.Close()

	lastId, err := insertIntoTasks(db, t)
	if err != nil {
		return -1, err
	}

	err = insertIntoTags(lastId, db, t)
	if err != nil {
		return -1, err
	}
	id = int(lastId)

	return id, nil
}
