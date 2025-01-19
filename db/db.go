package db

import(
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var DB *sql.DB
func InitDB()  {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")
	if err != nil {
		panic("error occured while opening database")
	}
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	createTable()

}

// func createTable(){
// 	createEventsTable := `
// 	CREATE TABLE IF NOT EXISTS events (
// 		id INTEGER PRIMARY KEY AUTOINCREMENT,
// 		name TEXT NOT NULL,
// 		description TEXT NOT NULL,
// 		location TEXT NOT NULL,
// 		date_time DATETIME NOT NULL,
// 		user_id INTEGER NOT NULL
// 	)
// 	`

// 	log.Println("Executing query:", createEventsTable)

// 	_,err:=DB.Exec(createEventsTable)
// 	if err != nil {
// 		panic("error occured while creating events table")	
// 	}


// }

func createTable() {
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		date_time DATETIME NOT NULL,
		user_id INTEGER NOT NULL
	)
	`

	log.Println("Executing query:", createEventsTable)

	_, err := DB.Exec(createEventsTable)
	if err != nil {
		// Log the actual error instead of a generic message
		log.Fatalf("Error occurred while creating events table: %v", err)
	}

	log.Println("Table 'events' created or already exists.")
}
