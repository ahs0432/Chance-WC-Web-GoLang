package database

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DBConn *sql.DB

func Init(host string, port string, protocol string, user string, password string, dbname string) {
	DBConn = dbConn(host, port, protocol, user, password, dbname)

	_, table_check := DBConn.Query("SELECT * FROM WEB")

	if table_check != nil {
		fmt.Println(table_check)
		os.Exit(3)
	}

	_, table_check = DBConn.Query("SELECT * FROM WEBUPTIME")

	if table_check != nil {
		fmt.Println(table_check)
		os.Exit(3)
	}

	_, table_check = DBConn.Query("SELECT * FROM CHKRESULT")

	if table_check != nil {
		fmt.Println(table_check)
		os.Exit(3)
	}

	_, table_check = DBConn.Query("SELECT * FROM WEBGROUP")

	if table_check != nil {
		fmt.Println(table_check)
		os.Exit(3)
	}
}

func dbConn(host string, port string, protocol string, user string, password string, dbname string) *sql.DB {
	db, err := sql.Open("mysql", user+":"+password+"@"+protocol+"("+host+":"+port+")/"+dbname)

	if err != nil {
		//errCheck(err, "Database Connection")
		fmt.Println(err)
	} else {
		db.SetConnMaxLifetime(time.Minute * 5)
		db.SetConnMaxIdleTime(time.Minute * 1)
		db.SetMaxOpenConns(1024)
		db.SetMaxIdleConns(1024)

		return db
	}
	return nil
}

func dbClose(db *sql.DB) {
	db.Close()
}
