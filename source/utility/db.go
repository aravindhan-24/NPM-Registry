package utility

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type SQLiteDBConnector struct {
	Connector *sql.DB
}

func (db *SQLiteDBConnector) CreateTables() bool {
	_, err := db.Connector.Exec("CREATE TABLE IF NOT EXISTS credentials (id INTEGER PRIMARY KEY AUTOINCREMENT, user_name TEXT NOT NULL, password TEXT NOT NULL)")
	if err != nil {
		log.Println("Unable to create credentials table")
		return false
	}
	return true
}

func (db *SQLiteDBConnector) AddAdminCredentials() bool {
	var id int
	var userName, password string
	err := db.Connector.QueryRow(
		"SELECT id, user_name,password FROM credentials WHERE user_name = ? AND password = ?",
		"admin",
		"admin").Scan(&id, &userName, &password)
	if err != nil {

		if err == sql.ErrNoRows {
			hashedPassword, err := HashPassword("admin")
			if err != nil {
				log.Println("Unable to hash password")
				return false
			}
			_, err = db.Connector.Exec(
				"INSERT INTO credentials (user_name, password) VALUES(?, ?)",
				"admin",
				hashedPassword,
			)
			if err != nil {
				return false
			}
		}
	}
	return true
}

func NewSQLiteDB(dbFilePath string) (*SQLiteDBConnector, error) {
	con, err := sql.Open("sqlite3", dbFilePath)
	if err != nil {
		return nil, err
	}
	if err := con.Ping(); err != nil {
		return nil, err
	}
	return &SQLiteDBConnector{Connector: con}, nil
}

func (db *SQLiteDBConnector) Exec(query string, args ...any) (sql.Result, error) {
	return db.Connector.Exec(query, args...)
}

func (db *SQLiteDBConnector) Query(query string, args ...any) (*sql.Rows, error) {
	return db.Connector.Query(query, args...)
}

func (db *SQLiteDBConnector) Close() error {
	return db.Connector.Close()
}
