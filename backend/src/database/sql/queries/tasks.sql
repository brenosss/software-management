-- tasks.sql

-- name: create-task
INSERT INTO task (
    project_id,
    name,
    type,
    description,
    complexity,
    duration,
    progress
) values($1, $2, $3, $4, $5, $6, $7);
