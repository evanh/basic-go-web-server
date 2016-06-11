package db

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type AugmentedDB struct {
	db *sql.DB
}

func NewDB() (*AugmentedDB, error) {
	db, err := sql.Open("postgres", "user=postgres password=password dbname=postgres sslmode=disable")
	if err != nil {
		return nil, err
	}

	return &AugmentedDB{db}, nil
}

func (ag *AugmentedDB) MapQuery(query string, args ...interface{}) ([]map[string]interface{}, error) {
	rows, err := ag.db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	length := len(columns)

	results := []map[string]interface{}{}
	for rows.Next() {
		v := make([]interface{}, length)
		values := make([]interface{}, length)
		for i := range v {
			values[i] = &(v[i])
		}
		rows.Scan(values...)
		res := make(map[string]interface{}, length)
		for i := 0; i < length; i++ {
			res[columns[i]] = values[i]
		}
		results = append(results, res)
	}

	return results, nil
}

func (ag *AugmentedDB) Exec(query string, args ...interface{}) (sql.Result, error) {
	return ag.db.Exec(query, args...)
}
