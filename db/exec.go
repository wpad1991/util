package db

import (
	"database/sql"
	"errors"
)

func Exec(db *sql.DB, cmd string, iscol bool) ([][]interface{}, error) {

	rows, err := db.Query(cmd)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	cols, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	if cols == nil {
		return nil, errors.New("Exec col is null")
	}

	colsize := len(cols)

	result := make([][]interface{}, 0)

	rcol := make([]interface{}, colsize)
	vals := make([]interface{}, colsize)

	for i := 0; i < colsize; i++ {
		vals[i] = new(interface{})
		if iscol {
			rcol[i] = cols[i]
		}
	}

	if iscol {
		result = append(result, rcol)
	}

	for rows.Next() {

		rcol = make([]interface{}, colsize)

		err = rows.Scan(vals...)

		for i := 0; i < colsize; i++ {
			rcol[i] = *vals[i].(*interface{})
		}

		result = append(result, rcol)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}
	return result, nil
}
