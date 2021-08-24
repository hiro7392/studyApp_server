package main

import (
	"fmt"
	"time"
)


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

func (task *Task) deleteStudent() (err error) {

	stmtDelete, err := db.Prepare("DELETE FROM tasks WHERE id=?")
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





func (student *Student) createStudent() (err error) {
	fmt.Println(student.Name,student.Password)
	var t=time.Now()

	const layout2 = "2006-01-02 15:04:05"

	_,err=db.Query("INSERT INTO students(name,pass,created_at,updated_at) VALUES($1,$2,$3,$4)",student.Name,student.Password,t.Format(layout2),t.Format(layout2))
	//_,err=db.Query("INSERT INTO students(name,pass) VALUES(?,?)",student.Name,student.Password)

	fmt.Println("create check point 3",t.Format(layout2))
	return
}