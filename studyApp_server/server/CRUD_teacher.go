package main

import "log"

func getOneTeacherByName(Name string) (teacher_id int , err error) {
	
	
	//参考書通りの実装
	
	//var tmp_id int
	err = db.QueryRow("SELECT id FROM  teachers WHERE name =?", Name).Scan(&teacher_id)
	if err != nil {
		log.Println(err)
	}
	//fmt.Println("retrive check 3",post.Id,post.Content)
	return
}

//teacher_idを渡してteacherの名前を返す
func getOneTeacherNameById(Id int) (teacher_name string , err error) {
	
	
	//参考書通りの実装
	
	//var tmp_id int
	err = db.QueryRow("SELECT name FROM  teachers WHERE id =?", Id).Scan(&teacher_name)
	if err != nil {
		log.Println(err)
	}
	//fmt.Println("retrive check 3",post.Id,post.Content)
	return
}