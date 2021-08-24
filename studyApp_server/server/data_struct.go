package main

import "time"


type Task struct{
	Id 			int		`json:"id"`
	StudentId 	int		`json:"student_id"`
	Taskclass 	int 	`json:"task_class"`
	Name		string	`json:"name"`
	Content		string	`json:"content"`
	Createdat  time.Time `json:"created_at"`
	Updatedat  time.Time `json:"updated_at"`
	Deadline   time.Time `json:"deadline"`
	Achivement 	int		`json:"achivement"`

}
type Student struct{
	Name 		string `json:"name"`
	Password 	string `json:"pass"`
}