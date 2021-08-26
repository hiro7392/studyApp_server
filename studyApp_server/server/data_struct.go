package main

//import "time"


// type Task struct{
// 	Id 			int		`json:"id"`
// 	StudentId 	int		`json:"student_id"`
// 	Taskclass 	int 	`json:"task_class"`
// 	Name		string	`json:"name"`
// 	Content		string	`json:"content"`
// 	Createdat  time.Time `json:"created_at"`
// 	Updatedat  time.Time `json:"updated_at"`
// 	Deadline   time.Time `json:"deadline"`
// 	Achivement 	int		`json:"achivement"`
// }
type Task struct{
	Id int
	StudentId int
	Taskclass int
	TextName string
	Textbook_id int
	Startday int
	Deadline int
	Nowpage int
	Allpage int
}
type Student struct{ 
	Id			int
	Teacher_id  int
	Name 		string `json:"name"`
	Password 	string `json:"pass"`
	NowSchool 	string
	WantSchool	string
	Grade 		int
}

type Textbook struct{
	Id				int
	Name			string
	Subject_name 	string
	Subject_id		int
	Page			int
}
type Taskmath struct{
	Name	string
	Start	int
	End		int
}
type Gnt struct{
	Name	string
	Start	int
	End		int
}

type TaskBox struct{
	TaskJapa 	[]Gnt
	TaskMath 	[]Gnt
	TaskEng		[]Gnt
	TaskSci		[]Gnt
	TaskSoc		[]Gnt
}