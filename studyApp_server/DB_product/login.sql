CREATE TABLE users(
    id INT UNSIGNED NOT NULL AUTO_INCREMENT,

    name VARCHAR(100) NOT NULL,
    pass VARCHAR(100) NOT NULL,
    role INT UNSIGNED NOT NULL,
    PRIMARY KEY(id)
);