CREATE TABLE `user`(
    `uid` INTEGER PRIMARY KEY AUTOINCREMENT, 
    `email` VARCHAR(30) NOT NULL,
    `passw` VARCHAR(30) NOT NULL
); 

CREATE TABLE `posts`(
    `pid` INT NOT NULL,
    `category` VARCHAR(10) NOT NULL, 
    `picture` BLOB NOT NULL,
    PRIMARY KEY(pid),
    FOREIGN KEY(pid) REFERENCES user(uid)
); 

CREATE TABLE `comments`(
    `cid`INT NOT NULL, 
    `content` VARCHAR(280) NOT NULL, 
    PRIMARY KEY(cid),
    FOREIGN KEY(cid) REFERENCES user(uid)
); 