package mysqlTest

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

const (
	USERNAME = "root"
	PASSWORD = "xx921226"
	NETWORK  = "tcp"
	SERVER   = "134.175.1.59"
	PORT     = 3306
	DATABASE = "test"
)

type customer struct {
	custId   int
	custName string
	custAddress string
	custCity string
	custState string
	custZip string
	custCountry string
	custContact string
	custEmail string
}

func Connect() *sql.DB {
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s",
		USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	DB, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("Open mysql failed, err: %v\n", err)
		return nil
	}

	DB.SetConnMaxLifetime(100*time.Second)  //最大连接周期
	DB.SetMaxOpenConns(100)  //最大连接数
	DB.SetMaxIdleConns(16)  //最大闲置连接数

	return DB
}

func QueryOne( DB *sql.DB) (*customer, error) {
	//fmt.Println("query time: ")
	cust := new(customer)
	row := DB.QueryRow("select * from customers where cust_id=?", 10001)
	if err := row.Scan(&cust.custId, &cust.custName, &cust.custAddress, &cust.custCity,
		&cust.custState, &cust.custZip, &cust.custCountry, &cust.custContact, &cust.custEmail); err != nil {
		fmt.Printf("scan failed, err:%v", err)
		return nil ,err
	}
	return cust, nil
}