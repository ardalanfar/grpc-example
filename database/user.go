package database

import (
	"database/sql"
	"log"
)

type User struct {
	ID     uint
	Name   string
	Family string
}

type SQLhandler struct {
	db *sql.DB
}

func CreateDBConnection() *SQLhandler {
	return &SQLhandler{
		db: Connect(),
	}
}

//Get users
func (handler *SQLhandler) GetUsres() ([]User, error) {
	rows, err := handler.db.Query("select * from user")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []User{}
	for rows.Next() {
		u := User{}
		err = rows.Scan(&u.ID, &u.Name, &u.Family)
		if err != nil {
			log.Println(err)
			continue
		}
		users = append(users, u)
	}
	return users, nil
}

//Get user by name
func (handler *SQLhandler) GetUserByName(name string) (User, error) {
	row := handler.db.QueryRow("select * from user where name=?", name)
	u := User{}

	err := row.Scan(&u.ID, &u.Name, &u.Family)
	if err != nil {
		return u, err
	}
	return u, nil
}
