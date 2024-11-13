package db

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite"
)

const (
	sqlDriver    = "sqlite"
	dataBasePath = "data/db/tasklist.db"
)

var (
	dataBaseConst *sql.DB
)

func DeployDB(dataBasePath string) error {

	// os.Remove(dataBasePath) //?

	err := os.MkdirAll(filepath.Dir(dataBasePath), os.ModePerm)
	if err != nil {
		return fmt.Errorf("- fail to creating dir %v, %v \n", filepath.Dir(dataBasePath), err)
	}

	db, err := sql.Open(sqlDriver, dataBasePath)
	if err != nil {
		return err
	}
	defer db.Close()

	{ // Creating task table

		createTaskTableStmt := `
		CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT, 
		text TEXT NOT NULL, 
		due DATETIME);`

		_, err := db.Exec(createTaskTableStmt)
		if err != nil {
			return fmt.Errorf("fail to creating table 'tasks' %v", err)
		}
	}

	{
		tagsTableStmt := `
	CREATE TABLE IF NOT EXISTS tags (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		task_id INTEGER NOT NULL,
		tag TEXT,
		FOREIGN KEY (task_id) REFERENCES tasks (id) ON DELETE CASCADE
	);`

		_, err := db.Exec(tagsTableStmt)
		if err != nil {
			return fmt.Errorf("fail to creating table 'tags' %v", err)
		}

	}

	dataBaseConst = db

	return nil
}
