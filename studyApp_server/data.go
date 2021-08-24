package main

import (
	"database/sql" //ここでパッケージをimport
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" //コード内で直接参照するわけではないが、依存関係のあるパッケージには最初にアンダースコア_をつける
)

var db *sql.DB

func init() {
	var err error
	//mysqlへ接続。ドライバ名（mysql）と、ユーザー名・データソース(ここではgosample)を指定。
	db, err = sql.Open("mysql", "root@/studyApp")
	if err != nil {
		panic(err.Error())
	}
	
	log.Println("init success!\nConnected to mysql.")
	defer db.Close()
}

//引っ張ってきたデータを割り当てるための構造体を用意
type Person struct {
	Id   int
	Name string
}

// func judgeYesNo(purpose string) bool {

// 	var addNewuser string
// 	fmt.Printf("Do you want to %s ?: (Y/n)", purpose)
// 	fmt.Scan(&addNewuser)
// 	if addNewuser == "Y" || addNewuser == "y" {
// 		return true
// 	} else {
// 		return false
// 	}
// }

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
func showAll() {
	showAll_taskMath()
	showAll_taskEng()
	showAll_taskJapa()

}

//readして全てのカラムを表示
func showAll_taskMath() {
	fmt.Println("show now table")

	//データベースへクエリを送信。引っ張ってきたデータがrowsに入る。
	rows, err := db.Query("SELECT * FROM task_math")
	defer rows.Close()
	if err != nil {
		panic(err.Error())
	}

	//レコード一件一件をあらかじめ用意しておいた構造体に当てはめていく。
	for rows.Next() {
		var task Task //構造体Person型の変数personを定義
		err := rows.Scan(&task.Id, &task.Content, &task.StudentId)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(task.Id, task.Content, task.StudentId) //結果　1 yamada 2 suzuki
	}
}

//readして全てのカラムを表示
func showAll_taskEng() {
	fmt.Println("show now table")

	//データベースへクエリを送信。引っ張ってきたデータがrowsに入る。
	rows, err := db.Query("SELECT * FROM task_eng")
	defer rows.Close()
	if err != nil {
		panic(err.Error())
	}

	//レコード一件一件をあらかじめ用意しておいた構造体に当てはめていく。
	for rows.Next() {
		var task Task //構造体Person型の変数personを定義
		err := rows.Scan(&task.Id, &task.Content, &task.StudentId)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(task.Id, task.Content, task.StudentId) //結果　1 yamada 2 suzuki
	}
}

//readして全てのカラムを表示
func showAll_taskJapa() {
	fmt.Println("show now table")

	//データベースへクエリを送信。引っ張ってきたデータがrowsに入る。
	rows, err := db.Query("SELECT * FROM task_japa")
	defer rows.Close()
	if err != nil {
		panic(err.Error())
	}

	//レコード一件一件をあらかじめ用意しておいた構造体に当てはめていく。
	//レコード一件一件をあらかじめ用意しておいた構造体に当てはめていく。
	for rows.Next() {
		var task Task //構造体Person型の変数personを定義
		err := rows.Scan(&task.Id, &task.Content, &task.StudentId)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(task.Id, task.Content, task.StudentId) //結果　1 yamada 2 suzuki
	}
}

//投稿を一つだけ取り出して返す
func retrieve(Id int) (task Task , err error) {
	db, err = sql.Open("mysql", "root@/studyApp?parseTime=true")
	if err != nil {
		panic(err.Error())
	}
	
	//参考書通りの実装
	task = Task {}
	err = db.QueryRow("SELECT id,content,name FROM tasks WHERE id =?", Id).Scan(&task.Id,&task.Content, &task.Name)
	
	if err != nil {
		log.Println(err)
	}
	//fmt.Println("retrive check 3",post.Id,post.Content)
	return
}
func getAllTask(Id int) (task []Task , err error) {

	db, err = sql.Open("mysql", "root@/studyApp?parseTime=true")
	if err != nil {
		panic(err.Error())
	}
	
	//参考書通りの実装
	task = []Task{}
	//err = db.QueryRow("SELECT id,content,name FROM tasks WHERE student_id =?", Id).Scan(&task.Id,&task.Content, &task.Name)
	rows,err:=db.Query("select * from tasks where student_id=?",Id)
	for rows.Next() {
		var tas Task //構造体Person型の変数personを定義
		err := rows.Scan(&tas.Id,&tas.StudentId,&tas.Taskclass,&tas.Name,&tas.Content,&tas.Createdat,&tas.Updatedat,&tas.Deadline,&tas.Achivement)
		task=append(task,tas)
		if err != nil {
			panic(err.Error())
		}

		//fmt.Println(task.Id, task.Content, task.StudentId) //結果　1 yamada 2 suzuki
	}

	if err != nil {
		log.Println(err)
	}
	//fmt.Println("retrive check 3",post.Id,post.Content)
	return
}

func (task *Task) update() (err error) {
	fmt.Println("update checkPoint 1",task.Id)
	dbUpdate, err := db.Prepare("UPDATE task_nath SET content =? , name =? WHERE id =?")
	if err != nil {
		panic(err.Error())
	}
	defer dbUpdate.Close()
	fmt.Println("update checkPoint 2",task.Id)

	//更新を実行
	_, err = dbUpdate.Exec(task.Content, task.Name,task.Id)
	return
}

func (task *Task) delete() (err error) {

	stmtDelete, err := db.Prepare("DELETE FROM task_math WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	defer stmtDelete.Close()

	_, err = stmtDelete.Exec(task.Id)
	if err != nil {
		panic(err.Error())
	}
	return
}

func (task *Task) create() (err error) {

	
	_,err=db.Query("INSERT INTO task_math(student_id,name,content,achivement) VALUES(?,?,?,?)",task.StudentId,task.Name,task.Content,task.Achivement)
	fmt.Println("create check point 3")
	return
}
