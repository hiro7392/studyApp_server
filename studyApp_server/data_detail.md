# データテーブル本番用

###### tags: `BBWCハッカソン`




### 生徒テーブル

```
insert students(name,pass,grade,teacher_id,nowSchool,wantSchool) values("外丸大樹","JVUabnCK",3,1,"C中学校","A高校");

insert students(name,pass,grade,teacher_id,nowSchool,wantSchool) values("落 絢來","FQBPY",3,3,"D中学校","C高校");
```


```
+----+--------------+----------+-------+------------+------------+------------+
| id | name         | pass     | grade | teacher_id | nowSchool  | wantSchool |
+----+--------------+----------+-------+------------+------------+------------+
|  1 | 外丸大樹     | JVUabnCK |     3 |          1 | C中学校    | A高校         |
+----+--------------+----------+-------+------------+------------+------------+

```

### 高校テーブル

```
insert into schools(name,level) values("A高校",55);
```

```
+----+---------+-------+
| id | name    | level |
+----+---------+-------+
|  1 | A高校   |    55  |
+----+---------+-------+
1 row in set (0.00 sec)
```

### 教師テーブル

```
insert into teachers(name,pass) values("斉藤","losnwme");
```

insert into teachers(name,pass) values("高橋","ksiencms");
```
+----+--------+---------+
| id | name   | pass    |
+----+--------+---------+
|  1 | 斉藤   | losnwme |
+----+--------+---------+
```


### ログイン用テーブル

role(生徒or教師について)
1-生徒
2-教師

```
insert into users(name,pass,role) values("高橋","csFLDpDY",2);

insert into users(name,pass,role) values("外丸 大樹","bROsLdLQ",1);

```


### 教材テーブル


```
insert into textbooks(name,subject_id,page) values("チャート式 英語",3,360)

insert into textbooks(name,subject_id,page) values("DB3000",3,360)

insert into textbooks(name,subject_id,page) values("ハイクラステスト数学",2,150)

```

```
+----+------------------------+------------+------+
| id | name                   | subject_id | page |
+----+------------------------+------------+------+
|  1 | チャート式 英語        |          3 |  360 |
+----+------------------------+------------+------+
1
```

### 科目テーブル

```
insert into subjects(id,name) values(1,"国語"),(2,"数学"),(3,"英語"),(4,"理科"),(5,"社会");
```

```
+----+--------+
| id | name   |
+----+--------+
|  1 | 国語   |
|  2 | 数学   |
|  3 | 英語   |
|  4 | 理科   |
|  5 | 社会   |
+----+--------+
```


### タスクテーブル

デモ用のデータ
```
insert into tasks(student_id,task_class,textbook_id,startday,deadline,nowpage,allpage) value(1,2,1,120,213,0,360);

```




```
+----+------------+------------+-------------+----------+----------+---------+---------+
| id | student_id | task_class | textbook_id | startday | deadline | nowpage | allpage |
+----+------------+------------+-------------+----------+----------+---------+---------+
|  1 |          1 |          2 |           1 |      120 |      213 |       0 |     360 |
+----+------------+------------+-------------+----------+----------+---------+---------+
```