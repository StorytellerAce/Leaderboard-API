CREATE TABLE  scores (

    id SERIAL PRIMARY KEY,

    player VARCHAR(100) NOT NULL,

    score INTEGER NOT NULL

);

INSERT INTO scores (player, score)
VALUES
('Alice', 950),
('Bob', 820),
('Charlie', 1100),
('David', 500),
('Eve', 760);