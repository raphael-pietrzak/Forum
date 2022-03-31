CREATE TABLE `user`(
    `uid` INTEGER PRIMARY KEY AUTOINCREMENT, 
    `email` VARCHAR(30) NOT NULL,
    `passw` VARCHAR(30) NOT NULL
); 

CREATE TABLE `posts`(
    `pid` INTEGER PRIMARY KEY AUTOINCREMENT,
    `category` VARCHAR(10) NOT NULL, 
    `picture` BLOB NOT NULL,
    FOREIGN KEY(pid) REFERENCES user(uid)
); 

CREATE TABLE `comments`(
    `cid` INTEGER PRIMARY KEY AUTOINCREMENT, 
    `content` VARCHAR(280) NOT NULL, 
    FOREIGN KEY(cid) REFERENCES user(uid)
); 

INSERT INTO user VALUES(1,"emailTest","passwordTest")
INSERT INTO comments VALUES(1,"thisIsAComment",1)
