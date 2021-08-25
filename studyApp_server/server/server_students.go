package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"strconv"
	"log"
)

func handleRequestStudent(w http.ResponseWriter, r *http.Request){
	//w.Header().Set("Access-Control-Allow-Origin","http:")
	//w.Header().Set("Access-Control-Allow-Methods","POST")
	//w.Header().Set("Access-Control-Allow-Credentials",true)

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

	post,err:=getOneStudent(id)
	if err !=nil{
		return
	}
	fmt.Println("chechpoint3\n")
	//output,err :=json.MarshalIndent(&post,"","\t\t")
	output,err :=json.Marshal(&post)

	if err !=nil{
		return
	}
	w.Header().Set("Content-Type","application/json")
	w.Write(output)
	//showAll()
	return
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