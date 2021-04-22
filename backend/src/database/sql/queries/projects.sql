-- projects.sql

-- name: create-project
INSERT INTO project (name) VALUES($1);
