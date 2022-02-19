package dao

import (
	"database/sql"
	"log"
)

func ModifyDB(db *sql.DB, sql string, args ...interface{}) (int64, error)  {
	result,err:=db.Exec(sql, args...)
	if err!=nil {
		log.Println(err)
		return 0, err
	}
	count, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return count, nil
}

func ModifyDBID(db *sql.DB, sql string, args ...interface{}) (int64, error)  {
	result,err:=db.Exec(sql, args...)
	if err!=nil {
		return 0, err
	}
	count, err := result.LastInsertId()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return count, nil
}