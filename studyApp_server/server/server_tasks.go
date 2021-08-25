package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"strconv"
	
)

func handleRequestTask(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin","http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Methods","GET")
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


func handleGetTask(w http.ResponseWriter,r *http.Request)(err error){
	
	w.Header().Set("Access-Control-Allow-Origin","http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Methods","GET")

	fmt.Println("get Taskchechpoint1")
	id,err	:=strconv.Atoi(path.Base(r.URL.Path))
	if err!=nil{
		return
	}
	fmt.Println("chechpoint2 id=",id)

	post,err:=readOneTask(id)
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

func handlePostTask(w http.ResponseWriter,r *http.Request)(err error){
	// len :=r.ContentLength
	// //バイト列を作成
	// body:=make([]byte,len)
	// //バイト列にリクエストの本体を読み込み
	// r.Body.Read(body)
	// var task Task
	// //バイト列を構造体Postに組み替え
	// json.Unmarshal(body,&task)
	// //データベースのレコードを作成
	// err =student.create()
	// if err != nil{
	// 	return
	// }
	// w.WriteHeader(200)
	
	// showAll()
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

	post,err:=getAllTask(id)
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
	
	return
}
