package main

import (
	
	"net/http"
	"strconv"
	"html/template"
	"log"
	"fmt"
)


func main(){
	server :=http.Server{
		Addr:"127.0.0.1:8080",
	}
	http.Handle("/client/", http.StripPrefix("/client/", http.FileServer(http.Dir("client/"))))

	//テンプレート(gtpl)をロード
    templates["login"] = loadTemplate("login")
	http.HandleFunc("/login/", login)
	http.HandleFunc("/",first)

	//ログインから生徒用画面or教師用画面への遷移
	http.HandleFunc("/auth/", auth)

	//生徒登録情報修正画面(register_student.html)

	//講師用画面
	templates["teacher"] = loadTemplate("teacher")	
	http.HandleFunc("/teacher", teacher)

	//講師用生徒検索画面
	//templates["teacher_student"] = loadTemplate("teacher_student")

	//講師用
	templates["search_student"] = loadTemplate("search_student")
	http.HandleFunc("/search_student", handleGetStudentByTeacher)

	
	//講師用　生徒一覧表示
	templates["show_all_student"] = loadTemplate("show_all_student")
	http.HandleFunc("/show_all_student/", handleGetStudentByTeacherId)

	//講師用　教材一覧表示
	templates["show_all_textbook"] = loadTemplate("show_all_textbook")
	http.HandleFunc("/show_all_textbook/", show_all_textbook)

	//生徒情報　表示(教師用)
	templates["teacher_student"]=loadTemplate("teacher_student")
	http.HandleFunc("/teacher_student/",handleStudentDetail)

	//生徒情報　表示(生徒用)
	templates["teacher_student_for_stu"]=loadTemplate("teacher_student_for_stu")
	http.HandleFunc("/teacher_student_for_stu/",handleStudentDetailForStu)


	templates["info_student"] = loadTemplate("info_student")
	http.HandleFunc("/info_student/", info_student)

	//生徒登録
	templates["register_student"] = loadTemplate("register_student")
	http.HandleFunc("/register_student/", regist_student)

	//教材登録
	templates["register_textbook"] = loadTemplate("register_textbook")
	http.HandleFunc("/register_textbook/", register_textbook)

	//タスク更新
	http.HandleFunc("/update_task/",update_task)
	//タスク新規登録
	http.HandleFunc("/create_task/",create_task)

    // templates["onestudent"] = loadTemplate("onestudent")
    // templates["onetask"] = loadTemplate("onetask")
	// templates["alltask"] = loadTemplate("alltask")
	//http.HandleFunc("/",handleLogin)

	//http.HandleFunc("/", index)


	http.HandleFunc("/student/",handleRequestStudent)
	http.HandleFunc("/alltask/",showAll_task)
	http.HandleFunc("/task/",handleRequestTask)
	
	server.ListenAndServe()
	
}

var templates = make(map[string]*template.Template)

func index(w http.ResponseWriter, r *http.Request) {

    data := struct {
        Title  string
        Footer string
    }{
        Title:  "Go template Lesson",
        Footer: "2020 Go template Lesson",
    }
    if err := templates["index"].Execute(w, data); err != nil {
        log.Printf("failed to execute template: %v", err)
    }

}

func first(w http.ResponseWriter, r *http.Request){
	fmt.Printf("first called\n")
	http.Redirect(w, r, "/login", 301)
}
func login(w http.ResponseWriter, r *http.Request) {

    data := struct {
        Title  string
        
    }{
        Title:  "Go template Lesson",
        
    }
    if err := templates["login"].Execute(w, data); err != nil {
        log.Printf("failed to execute template: %v", err)
    }
}

func teacher(w http.ResponseWriter, r *http.Request) {
	data := struct {
        Title  string
        Name   string
        Footer string
    }{
        Title:  "Go template Lesson",
        Name:   r.FormValue("name"),
        Footer: "2020 Go template Lesson",
    }
	if err := templates["teacher"].Execute(w,data); err != nil {
        log.Printf("failed to execute template: %v", err)
    }
}
func info_student(w http.ResponseWriter, r *http.Request){
	data := struct {
        Title  string
        Name   string
        Footer string
    }{
        Title:  "Go template Lesson",
        Name:   r.FormValue("name"),
        Footer: "2020 Go template Lesson",
    }
	if err := templates["info_student"].Execute(w,data); err != nil {
        log.Printf("failed to execute template: %v", err)
    }
}
func loadTemplate(name string) *template.Template {
    t, err := template.ParseFiles("template/" + name + ".html")
    if err != nil {
        log.Fatal("ParseFiles: ", err)
    }
    return t
}
//生徒情報の登録
func regist_student(w http.ResponseWriter, r *http.Request){

	grade,err:=strconv.Atoi(r.FormValue("grade"))
	if err!=nil{
		log.Println(err)
		return
	}
	student_name:=r.FormValue("student_name")
	teacher_name:=r.FormValue("teacher_name")
	teacher_id,err:=getOneTeacherByName(teacher_name)

	nowSchool:=r.FormValue("school")
	wantSchool:=r.FormValue("school_of_choise")
	pass:="newPass"
	_,err=db.Query("INSERT INTO students(name,pass,grade,teacher_id,nowSchool,wantSchool) VALUES(?,?,?,?,?,?)",student_name,pass,grade,teacher_id,nowSchool,wantSchool)
	if err!=nil{
		log.Println(err)
		return
	}
	data:=struct {
		Grade int
		Teacher_name string
		NowSchool string
		WantSchool string
	}{
		Grade:grade,
		Teacher_name:teacher_name,
		NowSchool:nowSchool,
		WantSchool:wantSchool,
	}


	if err := templates["register_student"].Execute(w,data); err != nil {
        log.Printf("failed to execute template: %v", err)
    }
}

func register_textbook(w http.ResponseWriter, r *http.Request){
	textName:=r.FormValue("text_title")
	subject_name:=r.FormValue("subject")
	subject_id,err:=getSubjectIdBysubjectName(subject_name)
	if err!=nil{
		log.Println(err)
		return
	}
	pages,err:=strconv.Atoi(r.FormValue("pages"))
	_,err=db.Query("insert into textbooks(name,subject_id,page) VALUES(?,?,?)",textName,subject_id,pages)
	data:=struct {
		Name string
		Subject_id int
		Subject_name string
		Page int
	}{
		Name:textName,
		Subject_id:subject_id,
		Subject_name:subject_name,
		Page:pages,
	}
	if err := templates["register_textbook"].Execute(w,data); err != nil {
        log.Printf("failed to execute template: %v", err)
    }

}

//教材一覧


func show_all_textbook(w http.ResponseWriter,r *http.Request){
	
	fmt.Println("show_all_textbook called")
	
	
	
	post,err:=getAllTextbook()
	if err !=nil{
		return
	}
	data:=[]Textbook{}

	for _,v:=range post{
		var tb Textbook
		tb.Id=v.Id
		tb.Name=v.Name
		tb.Subject_name=v.Subject_name
		tb.Page=v.Page
		data=append(data,tb)
	}

	if err := templates["show_all_textbook"].Execute(w, data); err != nil {
        log.Printf("failed to execute template: %v", err)
    }
	
}
func auth(w http.ResponseWriter,r *http.Request){
	println("auth called")
	role,err:=strconv.Atoi(r.FormValue("role"))
	
	if err!=nil{
		log.Println(err)
	}
	name:=r.FormValue("username")
	pass:=r.FormValue("password")
	fmt.Println("role",role,"name",name)
	var Id int =0
	var truepass string=""
	//生徒のとき
	if role==1{
		err=db.QueryRow("select id,pass from students where name=?",name).Scan(&Id,&truepass)
		fmt.Println("role:生徒 id=",Id,"name=",name)

		if err!=nil{
			log.Println(err)
		}else if Id==0{
			http.Redirect(w, r, "../login", 301)
		}else if truepass==""{
			http.Redirect(w, r, "../login", 301)
		}else if pass!=truepass {
			http.Redirect(w, r, "../login", 301)
		}
		http.Redirect(w, r, "../teacher_student_for_stu/"+strconv.Itoa(Id), 301)
	}else{
		//教師のとき
		err:=db.QueryRow("select id,pass from teachers where name=?",name).Scan(&Id,&truepass)
		if err!=nil{
			log.Println(err)
		}else if Id==0{
			http.Redirect(w, r, "../login", 301)
		}else if truepass==""{
			http.Redirect(w, r, "../login", 301)
		}else if pass!=truepass {
			http.Redirect(w, r, "../login", 301)
		}
		http.Redirect(w, r, "../teacher", 301)
	}
}
