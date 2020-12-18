USE sample_db;

CREATE TABLE users (
  id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(10)
);

INSERT INTO users (name) VALUES ("user1"), ("user2"), ("user3");
