package task3

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

var sqlxDB *sqlx.DB

func init() {
	dsn := "root:123456@tcp(localhost:3306)/gin?charset=utf8mb4&parseTime=true&loc=Local"
	sqlxDB, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Println(err)
	}
}

// 题目1
type Employee struct {
	Id         int
	Name       string
	Department string
	Salary     int
}

func (e Employee) TableName() string {
	return "employees "
}

func Query1() []Employee {
	var arr []Employee
	sqlxDB.Select(&arr, "select * from employees where department = ?", "技术部")
	return arr
}

func Query2() Employee {
	employee := Employee{}
	sqlxDB.Get(&employee, "select * from employees order by salary desc limit 1")
	return employee
}

// 题目2
type Book struct {
	Id     int
	Title  string
	Author string
	Price  int
}

func (b Book) TableName() string {
	return "books"
}

func QueryBook() []Book {
	var arr []Book
	sqlxDB.Select(&arr, "select * from books where price > ? ", 50)
	return arr
}
