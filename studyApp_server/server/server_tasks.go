package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"strconv"
	"time"
	"log"
)

func handleRequestTask(w http.ResponseWriter, r *http.Request){
	var err error
	switch	r.Method{
		case "GET":
			err=handleGetTask(w,r)
		case "POST":
			err=handlePostTask(w,r)
		case "PUT":
			err=handlePutTask(w,r)
		case "DELETE":
			err=handleDeleteTask(w,r)
	}
	if err!=nil{
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}
	
}
func calday(month int,day int)(sum int){
	sum=30*(month-1)+day
	if sum <=90{
		sum+=270
	}else{
		sum-=90
	}
	return sum
}

func create_task(w http.ResponseWriter, r *http.Request){

	fmt.Println("createTask called")
	stu_id,err :=strconv.Atoi(path.Base(r.URL.Path))
	//Formから値を取得
	name:=r.FormValue("text_title")
	// subject_class:=r.FormValue("subject")
	// pages:=r.FormValue("pages")
	endday:=r.FormValue("endday")
	fmt.Println("endday",endday)
	month,err:=strconv.Atoi(endday[5:7])
	if err != nil {
		log.Println(err)
	}
	fmt.Println(month)
	day,err:=strconv.Atoi(endday[8:10])
	if err != nil {
		log.Println(err)
	}
	fmt.Println(day)
	startday:=calday(8,27)
	deadline:=calday(month,day)
	//fmt.Println("deadline",deadline,"startday",startday)
	var subject_class int
	var pages int
	var text_id int
	err = db.QueryRow("SELECT id,page,subject_id FROM textbooks WHERE name=?", name).Scan(&text_id,&pages,&subject_class)
	if err != nil {
		log.Println(err)
	}

	_,err=db.Query("INSERT INTO tasks(student_id,task_class,textbook_id,startday,deadline,nowpage,allpage) values(?,?,?,?,?,?,?)",stu_id,subject_class,text_id,startday,deadline,0,pages)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("create task successful!")
	//fmt.Println("name",name,"subject",subject_class,"page",pages,"endday",endday,"text_id",text_id,"stu_id",stu_id)
	http.Redirect(w, r, "../teacher_student/"+strconv.Itoa(stu_id), 301)
	//_,err:=db.Query("INSERT INTO tasks(student_id,task_class,textbook_id,startday,deadline,nowpage,allpage) values(?,?,?,?,?,?,?)",stu_id,subject_class,text_id,
}

//タスクの更新
func update_task(w http.ResponseWriter, r *http.Request){
	//生徒Idの取得
	fmt.Println("createTask called")
	stu_id,err :=strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		log.Println(err)
	}

	//新しい現在のページ数
	nowpage,err:=strconv.Atoi(r.FormValue("pages"))
	if err != nil {
		log.Println(err)
	}
	name:=r.FormValue("text_title")
	

	var text_id int
	err=db.QueryRow("select id from textbooks where name=?",name).Scan(&text_id)
	if err != nil {
		log.Println(err)
	}
	
	
	dbUpdate, err := db.Prepare("UPDATE tasks SET nowpage=? WHERE textbook_id =? && student_id=? ")
	if err != nil {
		panic(err.Error())
	}

	

	//更新を実行
	_, err = dbUpdate.Exec(nowpage,text_id,stu_id)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("update task successful!")
	http.Redirect(w, r, "../teacher_student/"+strconv.Itoa(stu_id), 301)
}



func handleGetTask(w http.ResponseWriter,r *http.Request)(err error){
	
	fmt.Println("called handle getTask")
	id,err	:=strconv.Atoi(path.Base(r.URL.Path))
	if err!=nil{
		return
	}
	

	post,err:=readOneTask(id)
	if err !=nil{
		return
	}
	
	data := struct {
        id int
        
    }{
        id: post.Id,
    }
    if err := templates["onetask"].Execute(w, data); err != nil {
        log.Printf("failed to execute template: %v", err)
    }
	
	// w.Header().Set("Content-Type","application/json")
	// w.Write(output)
	
	return
}

func handlePostTask(w http.ResponseWriter,r *http.Request)(err error){
	len :=r.ContentLength
	//バイト列を作成
	body:=make([]byte,len)
	//バイト列にリクエストの本体を読み込み
	r.Body.Read(body)
	var task Task
	task=Task{}
	//バイト列を構造体studentに組み替え
	//json.Unmarshal(body,&student)

	//task.Taskclass=r.FormValue("taskclass")
	//task.Name=r.FormValue("name")
	//task.StudentId=r.FormValue("studentId")
	//task.Content=r.FormValue("content")
	
	//とりあえず仮決め
	const layout2 = "2006-01-02 15:04:05"

	//var t=time.Now()
	task.Taskclass=1
	// task.Name="テキストA 80pまで"
	
	// task.Createdat=t//.Format(layout2)
	// task.Updatedat=t//.Format(layout2)
	// task.Deadline=t//.Format(layout2)
	// task.Achivement=0
	
	//データベースのレコードを作成
	err =task.createTask()
	if err != nil{
		return
	}
	w.WriteHeader(200)
	
	return
}

func handlePutTask(w http.ResponseWriter,r *http.Request)(err error){
	id,err:=strconv.Atoi(path.Base(r.URL.Path))
	if err != nil{
		return
	}
	//データベースから構造体Postにデータを格納
	task,err :=readOneTask(id)
	if err !=nil{
		return
	}
	len :=r.ContentLength
	body :=make([]byte,len)
	//リクエスト本体からJSONデータを読み出し
	r.Body.Read(body)
	//JSONデータを構造体Postに組み替え
	json.Unmarshal(body,&task)
	//データベースを更新
	err =task.updateStudent()
	if err !=nil{
		return
	}
	w.WriteHeader(200)
	
	return
}

func handleDeleteTask(w http.ResponseWriter,r *http.Request)(err error){
	id,err :=strconv.Atoi(path.Base(r.URL.Path))
	if err !=nil{
		return
	}
	
	if err !=nil{
		return
	}
	err =deleteTask(id)
	if err !=nil{
		return
	}
	w.WriteHeader(200)
	
	return
}

type taskData struct{
	Name  string	`json:"name"`
	Content string  `json:"content"`
	Deadline time.Time 
}
//全てのタスクを表示
func showAll_task(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin","http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Methods","GET")
	fmt.Println("chechpoint1")
	id,err	:=strconv.Atoi(path.Base(r.URL.Path))
	if err!=nil{
		return
	}
	//fmt.Println("chechpoint2 id=",id)

	post,_,err:=getAllTask(id)
	if err !=nil{
		return
	}
	fmt.Println("chechpoint3\n")
	//output,err :=json.MarshalIndent(&post,"","\t\t")
	//output,err :=json.Marshal(&post)
	
	data :=[]taskData{}
	
	//要素を挿入
	for _,v:=range post{
		var task taskData
		fmt.Println(v)
		//task.Id=v.Id
		//task.task_class=v.task_class
		//task.Deadline=v.Deadline
		data=append(data,task)
		//fmt.Println(v.Name)
		//fmt.Println(v.Content)
		//fmt.Println(v)
	}
	sub:=struct {
		D []taskData
		Id  int
	}{
		D: 	data,
		Id: id,
	}


	//fmt.Println(data)
	if err := templates["alltask"].Execute(w, sub); err != nil {
        log.Printf("failed to execute template: %v", err)
    }


	if err !=nil{
		return
	}
	//w.Header().Set("Content-Type","application/json")
	//w.Write(output)
	
	return
}
