CREATE DATABASE educationWebsite;
USE educationWebsite;
CREATE TABLE users  (
    uID int AUTO_INCREMENT PRIMARY KEY, 
    email VARCHAR(40) NOT NULL, 
    uName VARCHAR(30) NOT NULL,
    pw VARCHAR(40) NOT NULL
);

CREATE TABLE cards (
    userOWN INT NOT NULL,
    setID INT  AUTO_INCREMENT PRIMARY KEY,
    setName VARCHAR(50),
    FOREIGN KEY (userOWN)
        REFERENCES users (uID)
);

CREATE TABLE terms (
    setID INT NOT NULL,
    term1 VARCHAR(140) NOT NULL,
    term2 VARCHAR(200) NOT NULL,
    FOREIGN KEY (setID)
        REFERENCES cards (setID)
);