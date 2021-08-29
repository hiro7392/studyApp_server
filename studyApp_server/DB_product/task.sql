CREATE TABLE tasks(
    id INT UNSIGNED NOT NULL AUTO_INCREMENT,
    student_id INT UNSIGNED NOT NULL,
    task_class INT UNSIGNED NOT NULL,
    textbook_id INT UNSIGNED ,
    startday    INT NOT NULL,
    deadline    INT NOT NULL,
    nowpage INT UNSIGNED NOT NULL,
    allpage INT UNSIGNED NOT NULL,
    FOREIGN KEY (student_id) REFERENCES students(id),
    PRIMARY KEY (id)
);