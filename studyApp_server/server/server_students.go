package main

import (
	//"encoding/json"
	"fmt"
	"net/http"
	"path"
	"strconv"
	"log"
)

func handleRequestStudent(w http.ResponseWriter, r *http.Request){
	
	var err error
	switch	r.Method{
		case "GET":
			err=handleGetStudent(w,r)
		case "POST":
			err=handleCreateStudent(w,r)
		case "PUT":
			err=handlePutStudent(w,r)
		case "DELETE":
			err=handleDeleteStudent(w,r)
	}
	if err!=nil{
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}
}

//生徒テーブルに関する処理


func handleCreateStudent(w http.ResponseWriter,r *http.Request)(err error){
	len :=r.ContentLength
	//バイト列を作成
	body:=make([]byte,len)
	//バイト列にリクエストの本体を読み込み
	r.Body.Read(body)
	var student Student
	student=Student{}
	//バイト列を構造体studentに組み替え
	//json.Unmarshal(body,&student)

	student.Name=r.FormValue("name")
	student.Password=r.FormValue("pass")
	fmt.Println(student)
	//データベースのレコードを作成
	err =student.createStudent()
	if err != nil{
		return
	}
	w.WriteHeader(200)
	
	return
}

func handleGetStudent(w http.ResponseWriter,r *http.Request)(err error){
	
	w.Header().Set("Access-Control-Allow-Origin","http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Methods","GET")

	fmt.Println("handleGetStudent called")
	id,err	:=strconv.Atoi(path.Base(r.URL.Path))
	if err!=nil{
		return
	}
	fmt.Println("chechpoint2 id=",id)

	post,err:=getOneStudentById(id)
	if err !=nil{
		return
	}
	fmt.Println("chechpoint3\n")

	//output,err :=json.MarshalIndent(&post,"","\t\t")
	//output,err :=json.Marshal(&post)

	data :=struct{
		Name string
		Password string
	}{
        Name:   post.Name,
        Password: post.Password,
    }
	if err := templates["onestudent"].Execute(w, data); err != nil {
        log.Printf("failed to execute template: %v", err)
		//return
    }
	// if err !=nil{
	// 	return
	// }
	//w.Header().Set("Content-Type","application/json")
	//w.Write(output)


	return
}

func handleGetStudentByTeacher(w http.ResponseWriter,r *http.Request){
	
	fmt.Println("handleGetStudentByTeachert called")
	student_name:=r.FormValue("student_name")
	fmt.Println("chechpoint2 id=",student_name)

	post,err:=getOneStudentByName(student_name)
	if err !=nil{
		return
	}
	teacher_name,err:=getOneTeacherNameById(post.Teacher_id)
	data :=struct{
		Id			int
		Name 		string 
		Password 	string 
		NowSchool 	string
		WantSchool	string
		Grade 		int
		Teacher_name string
	}{
		Id	:	post.Id,
        Name:   post.Name,
        Password: post.Password,
		NowSchool: post.NowSchool,
		WantSchool: post.WantSchool,
		Grade:post.Grade,
		Teacher_name:teacher_name,
    }
	if err := templates["search_student"].Execute(w, data); err != nil {
        log.Printf("failed to execute template: %v", err)
    }
	
}

//教師Aの担当する生徒を全て取得する
func handleGetStudentByTeacherId(w http.ResponseWriter,r *http.Request){
	
	fmt.Println("handleGetStudentByTeacherId called")
	
	teacher_name:=r.FormValue("teacher_name")

	Id,err:=getOneTeacherByName(teacher_name)

	//教師id=Id を満たす生徒データが配列で帰ってくる
	post,err:=getStudentsByTeacherId(Id)
	if err !=nil{
		return
	}
	data:=[]Student{}

	for _,v:=range post{
		var stu Student
		stu.Id=v.Id
		stu.Name=v.Name
		stu.Password=v.Password
		stu.NowSchool=v.NowSchool
		stu.WantSchool=v.WantSchool
		stu.Grade=v.Grade
		data=append(data,stu)
	}

	if err := templates["show_all_student"].Execute(w, data); err != nil {
        log.Printf("failed to execute template: %v", err)
    }
	
}
func handlePutStudent(w http.ResponseWriter,r *http.Request)(err error){
	// id,err:=strconv.Atoi(path.Base(r.URL.Path))
	// if err != nil{
	// 	return
	// }
	// //データベースから構造体Postにデータを格納
	// task,err :=readOneTask(id)
	// if err !=nil{
	// 	return
	// }
	// len :=r.ContentLength
	// body :=make([]byte,len)
	// //リクエスト本体からJSONデータを読み出し
	// r.Body.Read(body)
	// //JSONデータを構造体Postに組み替え
	// json.Unmarshal(body,&task)
	// //データベースを更新
	// err =task.update()
	// if err !=nil{
	// 	return
	// }
	// w.WriteHeader(200)
	
	return
}


func handleDeleteStudent(w http.ResponseWriter,r *http.Request)(err error){
	w.Header().Set("Access-Control-Allow-Methods","DELETE")

	id,err :=strconv.Atoi(path.Base(r.URL.Path))
	if err !=nil{
		log.Println(err)
		return
	}
	
	err =deleteStudent(id)
	if err !=nil{
		log.Println(err)
		return
	}
	w.WriteHeader(200)
	
	return
}

func handleStudentDetail(w http.ResponseWriter,r *http.Request){
	id,err :=strconv.Atoi(path.Base(r.URL.Path))
	if err !=nil{
		log.Println(err)
	}
	stu,err:=getOneStudentById(id)
	if err !=nil{
		log.Println(err)
	}
	tas,taskBox,err:=getAllTask(id)
	if err !=nil{
		log.Println(err)
	}
	var student_data Student
	student_data.Id=stu.Id
	student_data.Name=stu.Name
	student_data.NowSchool=stu.NowSchool
	student_data.WantSchool=stu.WantSchool
	
	
	data:=struct {
		Tas []Task
		St Student
		Taskbox TaskBox
	}{
		Tas: tas,
		St: stu,
		Taskbox: taskBox,
	}

	if err := templates["teacher_student"].Execute(w, data); err != nil {
        log.Printf("failed to execute template: %v", err)
    }
}
func handleStudentDetailForStu(w http.ResponseWriter,r *http.Request){
	id,err :=strconv.Atoi(path.Base(r.URL.Path))
	if err !=nil{
		log.Println(err)
	}
	stu,err:=getOneStudentById(id)
	if err !=nil{
		log.Println(err)
	}
	tas,taskBox,err:=getAllTask(id)
	if err !=nil{
		log.Println(err)
	}
	var student_data Student
	student_data.Id=stu.Id
	student_data.Name=stu.Name
	student_data.NowSchool=stu.NowSchool
	student_data.WantSchool=stu.WantSchool
	
	
	data:=struct {
		Tas []Task
		St Student
		Taskbox TaskBox
	}{
		Tas: tas,
		St: stu,
		Taskbox: taskBox,
	}

	if err := templates["teacher_student_for_stu"].Execute(w, data); err != nil {
        log.Printf("failed to execute template: %v", err)
    }
}