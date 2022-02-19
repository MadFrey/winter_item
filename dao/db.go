package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)


var DB *sql.DB

func Init(dns string) (err error)  {
	database,err:=sql.Open("mysql",dns)
	if err != nil {
		return err
	}
	DB=database
	err=DB.Ping()
	if err != nil {
		return err
	}

	return nil
}

func QueryRowsDB(DB *sql.DB, sqlstr string) (*sql.Rows, error) {
	return DB.Query(sqlstr)
}