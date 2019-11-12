package databaselayer

import (
	"database/sql"
	"log"
)

type SQLHandler struct {
	*sql.DB
}
func(handler *SQLHandler) GetAllBooks() ([]Book,error){
	Books := []Book{}
	rows, err := handler.Query("select * from bookslist")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		a := Book{}
		err := rows.Scan(&a.Id, &a.Name, &a.Author)
		if err != nil {
			log.Println(err)
			continue
		}
		Books = append(Books, a)
	}
	return Books, rows.Err()
}
