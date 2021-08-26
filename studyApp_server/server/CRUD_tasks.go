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
	//_, err = dbUpdate.Exec(task.Content, task.Name,task.Id)
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
	
	var t=time.Now()

	const layout2 = "2006-01-02 15:04:05"

	//_,err=db.Query("INSERT INTO tasks(student_id,name,content,created_at,updated_at,deadline,achivement) VALUES(?,?,?,?,?,?,?)",&task.StudentId,&task.Taskclass,&task.Name,&task.Name,&task.Createdat,&task.Updatedat,&task.Deadline,&task.Achivement)
	//_,err=db.Query("INSERT INTO students(name,pass) VALUES(?,?)",student.Name,student.Password)

	fmt.Println("create check point 3",t.Format(layout2))
	return
}
//投稿を一つだけ取り出して返す
func readOneTask(Id int) (task Task , err error) {
	//db, err = sql.Open("mysql", "root@/studyApp?parseTime=true")
	
	
	
	//参考書通りの実装
	task = Task {}
	
	//err = db.QueryRow("SELECT id,student_id,task_class,textbook_id,startday,deadline,nowpage,allpage FROM tasks WHERE id =?",Id).Scan(&tas.Id,&tas.StudentId,&tas.Taskclass,&tas.Textbook_id,&tas.Startday,&tas.Deadline,&tas.Nowpage,&tas.Allpage)

	if err != nil {
		log.Println(err)
	}
	//fmt.Println("retrive check 3",post.Id,post.Content)
	return
}

//投稿を一つだけ取り出して返す
func readAllTask(Id int) (task Task , err error) {
	
	
	//参考書通りの実装
	task = Task {}
	err = db.QueryRow("SELECT id, FROM tasks WHERE id =?", Id).Scan(&task.Id)
	if err != nil {
		log.Println(err)
	}
	//fmt.Println("retrive check 3",post.Id,post.Content)
	return
}

//全てのタスクを読み取り表示する
func getAllTask(Id int) (task []Task,taskBox TaskBox, err error) {

	
	//参考書通りの実装
	task = []Task{}
	taskBox= TaskBox{}
	//err = db.QueryRow("SELECT id,content,name FROM tasks WHERE student_id =?", Id).Scan(&task.Id,&task.Content, &task.Name)
	rows,err:=db.Query("select * from tasks where student_id=?",Id)
	for rows.Next() {
		var tas Task //構造体Person型の変数personを定義
		err := rows.Scan(&tas.Id,&tas.StudentId,&tas.Taskclass,&tas.Textbook_id,&tas.Startday,&tas.Deadline,&tas.Nowpage,&tas.Allpage)
		
		tas.TextName,err=getTextNameFromId(tas.Textbook_id)
		//ガントチャート用の左端と右端
		sf:=float64(tas.Startday)
		sd:=float64(tas.Deadline)
		tas.Startday=int(sf*((48.0)/360.0))
		tas.Deadline=int(sd*((48.0)/360.0))
		fmt.Println(tas.Startday,tas.Deadline)
		var gnt= Gnt{}
		gnt.Start	=tas.Startday
		gnt.End		=tas.Deadline
		gnt.Name=tas.TextName
		if tas.Taskclass ==1 {
			taskBox.TaskJapa=append(taskBox.TaskJapa,gnt)
			fmt.Println(taskBox.TaskJapa)
		}else if tas.Taskclass==2{
			taskBox.TaskMath=append(taskBox.TaskMath,gnt)
			fmt.Println(taskBox.TaskMath)
		}else if tas.Taskclass==3{
			taskBox.TaskEng=append(taskBox.TaskEng,gnt)
		}else if tas.Taskclass==4{
			taskBox.TaskSci=append(taskBox.TaskSci,gnt)
		}else if tas.Taskclass==5{
			taskBox.TaskSoc=append(taskBox.TaskSoc,gnt)
		}		
		task=append(task,tas)
		if err != nil {
			panic(err.Error())
		}
	}

	if err != nil {
		log.Println(err)
	}
	//fmt.Println("retrive check 3",post.Id,post.Content)
	return
}
func getTextNameFromId(Id int)(name string,err error){
	err=db.QueryRow("select name from textbooks where id=?",Id).Scan(&name)
	if err != nil {
		log.Println(err)
	}
	return
}