package main

import (
	"fmt"
	"time"
	"log"
)

func (task *Task) updateTask() (err error) {
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
//1件削除
func  deleteTask(Id int) (err error) {

	stmtDelete, err := db.Prepare("DELETE FROM tasks WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	defer stmtDelete.Close()

	_, err = stmtDelete.Exec(Id)
	if err != nil {
		panic(err.Error())
	}
	return
}

func (task *Task) createTask() (err error) {
	fmt.Println(task.Name,task.Content)
	var t=time.Now()

	const layout2 = "2006-01-02 15:04:05"

	//_,err=db.Query("INSERT INTO students(name,pass,created_at,updated_at) VALUES($1,$2,$3,$4)",&tas.Taskclass,&tas.Name,&tas.Content,&tas.Createdat,&tas.Updatedat,tas.Deadline,tas.Achivement)
	//_,err=db.Query("INSERT INTO students(name,pass) VALUES(?,?)",student.Name,student.Password)

	fmt.Println("create check point 3",t.Format(layout2))
	return
}
//投稿を一つだけ取り出して返す
func readOneTask(Id int) (task Task , err error) {
	//db, err = sql.Open("mysql", "root@/studyApp?parseTime=true")
	
	
	
	//参考書通りの実装
	task = Task {}
	
	err = db.QueryRow("SELECT id,content,name FROM tasks WHERE id =?", Id).Scan(&task.Id,&task.Content, &task.Name)
	if err != nil {
		log.Println(err)
	}
	//fmt.Println("retrive check 3",post.Id,post.Content)
	return
}
//投稿を一つだけ取り出して返す
func readAllTask(Id int) (task Task , err error) {
	//db, err = sql.Open("mysql", "root@/studyApp?parseTime=true")
	
	
	//参考書通りの実装
	task = Task {}
	err = db.QueryRow("SELECT id,content,name FROM tasks WHERE id =?", Id).Scan(&task.Id,&task.Content, &task.Name)
	if err != nil {
		log.Println(err)
	}
	//fmt.Println("retrive check 3",post.Id,post.Content)
	return
}

//全てのタスクを読み取り表示する
func getAllTask(Id int) (task []Task , err error) {

	
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