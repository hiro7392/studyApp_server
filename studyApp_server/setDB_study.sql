DROP TABLE IF EXISTS students;
DROP TABLE IF EXISTS task_math;
DROP TABLE IF EXISTS task_japa;
DROP TABLE IF EXISTS task_eng;


CREATE TABLE students(
    id INT UNSIGNED NOT NULL AUTO_INCREMENT,
    name VARCHAR(100) NOT NULL,
    pass VARCHAR(100) NOT NULL,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL,
    PRIMARY KEY(id)
);


CREATE TABLE task_math(
    id INT UNSIGNED NOT NULL AUTO_INCREMENT,
    student_id INT UNSIGNED NOT NULL,
    name VARCHAR(100) NOT NULL,
    content VARCHAR(100) NOT NULL,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL,
    deadline DATETIME NOT NULL,
    achivement INT UNSIGNED NOT NULL,
    FOREIGN KEY (student_id) REFERENCES students(id),
    PRIMARY KEY (id)
);
CREATE TABLE task_japa(
    id INT UNSIGNED NOT NULL AUTO_INCREMENT,
    student_id INT UNSIGNED NOT NULL,
    name VARCHAR(100) NOT NULL,
    content VARCHAR(100) NOT NULL,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL,
    deadline DATETIME NOT NULL,
    achivement INT UNSIGNED NOT NULL,
    FOREIGN KEY (student_id) REFERENCES students(id),
    PRIMARY KEY (id)
);

CREATE TABLE task_eng(
    id INT UNSIGNED NOT NULL AUTO_INCREMENT,
    student_id INT UNSIGNED NOT NULL,
    name VARCHAR(100) NOT NULL,
    content VARCHAR(100) NOT NULL,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL,
    deadline DATETIME NOT NULL,
    achivement INT UNSIGNED NOT NULL,
    FOREIGN KEY (student_id) REFERENCES students(id),
    PRIMARY KEY (id)
);

insert into students(id,name,,created_at,updated_at) 