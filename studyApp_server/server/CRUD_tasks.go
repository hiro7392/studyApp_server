package main

import (
	"fmt"
	"log"
	"strconv"
)

func (task *Task) updateTask() (err error) {
	fmt.Println("update checkPoint 1",task.Id)
	dbUpdate, err := db.Prepare("UPDATE tasks SET content =? , name =? WHERE id =?")
	if err != nil {
		panic(err.Error())
	}
	defer dbUpdate.Close()
	

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

	
	
	task = []Task{}
	taskBox= TaskBox{}
	//err = db.QueryRow("SELECT id,content,name FROM tasks WHERE student_id =?", Id).Scan(&task.Id,&task.Content, &task.Name)
	rows,err:=db.Query("select * from tasks where student_id=?",Id)
	for rows.Next() {
		var tas Task //構造体Person型の変数personを定義
		err := rows.Scan(&tas.Id,&tas.StudentId,&tas.Taskclass,&tas.Textbook_id,&tas.Startday,&tas.Deadline,&tas.Nowpage,&tas.Allpage)
		
		tas.TextName,err=getTextNameFromId(tas.Textbook_id)
		var endpage int
		err=db.QueryRow("select page from textbooks where id=?",tas.Textbook_id).Scan(&endpage)
		//fmt.Println("textbook_id",tas.Textbook_id,"endpage",endpage)
		tas.Allpage=endpage
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
		gnt.Endpage=tas.Allpage
		//今週やるページ数を求める
		resPage:=tas.Allpage-tas.Nowpage
		thisWeek:=resPage/(gnt.End-gnt.Start)
		tas.EndpageThisweek=tas.Nowpage+thisWeek
		//タスクのtextbook_idからそのタスクの平均を取得
		var ave int
		err= db.QueryRow("SELECT aves FROM aves WHERE text_id =?", tas.Textbook_id).Scan(&ave)
		
		var mes string
		//タスクが平均より進んでいるか否かで表示するメッセージを変更
		if ave<tas.Nowpage{
			mes="平均より"+strconv.Itoa(tas.Nowpage-ave)+"進んでいます!　いいペースです。"
			tas.Good=true
		}else{
			mes="平均より"+strconv.Itoa(ave-tas.Nowpage)+"遅れています!　少しペースを上げましょう"
			tas.Good=false
		}
		tas.Message=mes;
		
		if tas.Taskclass ==1 {
			taskBox.TaskJapa=append(taskBox.TaskJapa,gnt)
			fmt.Println("taskJapa",taskBox.TaskJapa)
		}else if tas.Taskclass==2{
			taskBox.TaskMath=append(taskBox.TaskMath,gnt)
			fmt.Println("taskMath",taskBox.TaskMath)
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