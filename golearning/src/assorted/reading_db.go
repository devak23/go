package assorted

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"strings"
)

var database *sql.DB
var err error

func init() {
	database, err = sql.Open("sqlite3", "/home/hades/Workspace/go/tasks.db")
	if err != nil {
		fmt.Printf("err while opening database: %s\n", err)
	}
}

func ReadingDBMain() {
	defer database.Close()
	getTaskSQL :=
		"SELECT id, title, content, created_date " +
			"FROM task " +
			"WHERE finish_date is null " +
			"AND is_deleted = 'N' " +
			"ORDER BY created_date ASC"
	rows, err := database.Query(getTaskSQL)
	if err != nil {
		fmt.Printf("err while retrieving rows %s\n", err)
	}

	for rows.Next() {
		TaskID := ""
		TaskTitle := ""
		TaskContent := ""
		TaskCreated := ""
		err := rows.Scan(&TaskID, &TaskTitle, &TaskContent, &TaskCreated)
		TaskContent = strings.Replace(TaskContent, "\n", "<BR>", -1)
		if err != nil {
			fmt.Printf("err while printing content: %s\n", err)
		}

		fmt.Println(TaskID, TaskTitle, TaskContent, TaskCreated)
	}
	defer rows.Close()

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
