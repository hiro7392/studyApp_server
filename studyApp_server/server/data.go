package main

import (
	"database/sql" //ここでパッケージをimport
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql" //コード内で直接参照するわけではないが、依存関係のあるパッケージには最初にアンダースコア_をつける
)

var db *sql.DB

func init() {
	var err error
	//mysqlへ接続。ドライバ名（mysql）と、ユーザー名・データソース(ここではgosample)を指定。
	//db, err = sql.Open("mysql", "root@/product_studyApp?parseTime=true")
	db, err = sql.Open("mysql", "root@/product_studyApp?parseTime=true")

	if err != nil {
		panic(err.Error())
	}
	//defer db.Close()

	log.Println("init success!\nConnected to mysql.")
	//defer db.Close()
}

//引っ張ってきたデータを割り当てるための構造体を用意
type Person struct {
	Id   int
	Name string
}

func getUserName() string {
	var str string

	fmt.Scan(&str)
	return str
}

func getuserId() int {
	var editUserId int
	fmt.Scan(&editUserId)
	return editUserId
}

//readして全てのカラムを表示

func retrieveUser(Id int) (student Student, err error) {
	db, err = sql.Open("mysql", "root@/product_studyApp?parseTime=true")
	if err != nil {
		panic(err.Error())
	}

	//参考書通りの実装
	student = Student{}
	err = db.QueryRow("SELECT id,content,name FROM students WHERE id =?", Id).Scan(&student.Password, &student.Name)

	if err != nil {
		log.Println(err)
	}
	return
}

func (student *Student) create() (err error) {
	fmt.Println(student.Name, student.Password)
	var t = time.Now()

	const layout2 = "2006-01-02 15:04:05"

	_, err = db.Query("INSERT INTO students(name,pass,created_at,updated_at) VALUES($1,$2,$3,$4)", student.Name, student.Password, t.Format(layout2), t.Format(layout2))

	fmt.Println("create check point 3", t.Format(layout2))
	return
}

