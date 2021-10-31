package week02

import (
	"database/sql"
	"errors"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

func newDB() *sql.DB{
	db, err := sql.Open("mysql",
		"root:@tcp(127.0.0.1:3306)/twi")
	if err != nil {
		log.Fatal("open: ", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("ping: ", err)
	}

	return db
}

type Author struct {
	id int
	name string
}

func GetAuthor(id int) (*Author, error) {
	db := newDB()
	var (
		// result *Author
		name string
	)


	err := db.QueryRow("select id, name from authors where id = ?", id).Scan(&id, &name)
	log.Printf("queryrow: ", err)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		// ErrNoRows is special
		log.Printf("record %d not found", id)
		return &Author{}, nil
	case err != nil:
		return nil, err
	}
	return &Author{id: id, name: name}, nil
}
