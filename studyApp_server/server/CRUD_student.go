package main

import (
	"fmt"
	"time"
	"log"
)

//生徒データを更新
func (task *Task) updateStudent() (err error) {
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

//生徒データを一件削除
func deleteStudent(Id int) (err error) {
	fmt.Println("start delete")
	stmtDelete, err := db.Prepare("DELETE FROM students WHERE id=?")
	if err != nil {
		log.Println(err)
		panic(err.Error())
	}
	defer stmtDelete.Close()
	fmt.Println("exec delete")
	_, err = stmtDelete.Exec(Id)
	if err != nil {
		log.Println(err)
		panic(err.Error())
	}
	return
}




//生徒データを新規作成
func (student *Student) createStudent() (err error) {
	fmt.Println(student.Name,student.Password)
	var t=time.Now()

	const layout2 = "2006-01-02 15:04:05"

	_,err=db.Query("INSERT INTO students(name,pass,created_at,updated_at) VALUES(?,?,?,?)",student.Name,student.Password,t.Format(layout2),t.Format(layout2))
	//_,err=db.Query("INSERT INTO students(name,pass) VALUES(?,?)",student.Name,student.Password)

	fmt.Println("create check point 3",t.Format(layout2))
	return
}
func getOneStudent(Id int) (student Student , err error) {
	//db, err = sql.Open("mysql", "root@/studyApp?parseTime=true")
	
	
	
	//参考書通りの実装
	student = Student {}
	//var tmp_id int
	err = db.QueryRow("SELECT name,pass FROM students WHERE id =?", Id).Scan(&student.Name,&student.Password)
	if err != nil {
		log.Println(err)
	}
	//fmt.Println("retrive check 3",post.Id,post.Content)
	return
}

