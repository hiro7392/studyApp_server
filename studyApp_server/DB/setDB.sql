DROP TABLE IF EXISTS posts;

CREATE TABLE posts (
id INT UNSIGNED NOT NULL AUTO_INCREMENT,
author VARCHAR(100) NOT NULL,
content VARCHAR(100) NOT NULL,
PRIMARY KEY(id)
);
insert into posts(id,author,content) values(1,"hiromichi","hello_Obon!"),(2,"yamaguchi","GoodBye"),(3,"Arnold_Schwarzenegger","Goodnight");
