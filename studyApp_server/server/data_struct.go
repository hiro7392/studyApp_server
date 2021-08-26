package main


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
	Ave     int
	Message string
	Good	bool
	EndpageThisweek int
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
	Endpage int
}

type TaskBox struct{
	TaskJapa 	[]Gnt
	TaskMath 	[]Gnt
	TaskEng		[]Gnt
	TaskSci		[]Gnt
	TaskSoc		[]Gnt
}