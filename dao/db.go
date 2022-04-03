package dao

import (
	"database/sql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *sql.DB

func Init(dns string) (err error) {
	database, err := sql.Open("mysql", dns)
	if err != nil {
		return err
	}
	DB = database
	err = DB.Ping()
	if err != nil {
		return err
	}

	return nil
}

func QueryRowsDB(DB *sql.DB, sqlstr string) (*sql.Rows, error) {
	return DB.Query(sqlstr)
}

func OpenTransaction() (*sql.Tx, error) {
	return DB.Begin()
}

func OpenGormLink() (*gorm.DB, error) {
	dns := "root:sjk123456@tcp(localhost:3306)/blog?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dns,   // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{NamingStrategy: schema.NamingStrategy{
		SingularTable: true,
	}})
	if err != nil {
		return db, err
	}
	return db, nil
}
