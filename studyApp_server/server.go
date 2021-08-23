package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"strconv"
)

type Post struct{
	Id 		int		`json:"id"`
	Content	string	`json:"content"`
	Author	string	`json:"author"`
}

type Task struct{
	Id 			int		`json:"id"`
	StudentId 	int		`json:"student_id"`
	Content		string	`json:"content"`
	Name		string	`json:"name"`
	Achivement 	int		`json:"achivement"`
}

func main(){
	server :=http.Server{
		Addr:"127.0.0.1:8081",
	}
	
	http.HandleFunc("/post/",handleRequest)
	server.ListenAndServe()
}
//リクエストを正しい関数に振り分けるためのハンドル関数
func handleRequest(w http.ResponseWriter, r *http.Request){
	var err error
	switch	r.Method{
		case "GET":
			err=handleGet(w,r)
		case "POST":
			err=handlePost(w,r)
		case "PUT":
			err=handlePut(w,r)
		case "DELETE":
			err=handleDelete(w,r)
	}
	if err!=nil{
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}
	
}

func handleGet(w http.ResponseWriter,r *http.Request)(err error){

	fmt.Println("chechpoint1")
	id,err	:=strconv.Atoi(path.Base(r.URL.Path))
	if err!=nil{
		return
	}
	//fmt.Println("chechpoint2 id=",id)
	post,err:=retrieve(id)
	if err !=nil{
		return
	}
	//fmt.Println("chechpoint3\n")
	output,err :=json.MarshalIndent(&post,"","\t\t")
	if err !=nil{
		return
	}
	w.Header().Set("Content-Type","application/json")
	w.Write(output)
	//showAll()
	return
}

func handlePost(w http.ResponseWriter,r *http.Request)(err error){
	len :=r.ContentLength
	//バイト列を作成
	body:=make([]byte,len)
	//バイト列にリクエストの本体を読み込み
	r.Body.Read(body)
	var task Task
	//バイト列を構造体Postに組み替え
	json.Unmarshal(body,&task)
	//データベースのレコードを作成
	err =task.create()
	if err != nil{
		return
	}
	w.WriteHeader(200)
	
	showAll()
	return
}

func handlePut(w http.ResponseWriter,r *http.Request)(err error){
	id,err:=strconv.Atoi(path.Base(r.URL.Path))
	if err != nil{
		return
	}
	//データベースから構造体Postにデータを格納
	task,err :=retrieve(id)
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
	err =task.update()
	if err !=nil{
		return
	}
	w.WriteHeader(200)
	showAll()
	return
}

func handleDelete(w http.ResponseWriter,r *http.Request)(err error){
	id,err :=strconv.Atoi(path.Base(r.URL.Path))
	if err !=nil{
		return
	}
	task,err :=retrieve(id)
	if err !=nil{
		return
	}
	err =task.delete()
	if err !=nil{
		return
	}
	w.WriteHeader(200)
	showAll()
	return
}