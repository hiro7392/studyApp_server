package main

import "log"

func getSubjectIdBysubjectName(Name string) (Id int, err error) {
	
	//var tmp_id int
	err = db.QueryRow("SELECT id FROM subjects WHERE name=?", Name).Scan(&Id)
	if err != nil {
		log.Println(err)
	}
	//fmt.Println("retrive check 3",post.Id,post.Content)
	return
}
func getSubjectNameById(Id int) (Name string, err error) {
	
	//var tmp_id int
	err = db.QueryRow("SELECT name FROM subjects WHERE id=?", Id).Scan(&Name)
	if err != nil {
		log.Println(err)
	}
	//fmt.Println("retrive check 3",post.Id,post.Content)
	return
}
func getAllTextbook()(textbooks []Textbook,err error){

	textbooks=[]Textbook{}
	rows,err:=db.Query("select * from textbooks")
	for rows.Next() {
		var textb Textbook //構造体Person型の変数personを定義
		err := rows.Scan(&textb.Id,&textb.Name,&textb.Subject_id,&textb.Page)
		textb.Subject_name,err=getSubjectNameById(textb.Subject_id)
		textbooks=append(textbooks,textb)
		if err != nil {
			panic(err.Error())
		}
	}
	return 
}
