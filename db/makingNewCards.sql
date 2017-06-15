USE educationWebsite;
INSERT INTO cards (userOwn) values(1);
INSERT INTO terms VALUES(1, "Hello", "World");
INSERT INTO terms VALUES(1, "print", "A function which outputs data to the console");
SELECT * FROM terms WHERE setID = 1;