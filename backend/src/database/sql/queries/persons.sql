-- persons.sql

-- name: get-person
SELECT
    person_id, name, email, phone, description
FROM person
WHERE person.person_id=$1


-- name: get-skills
SELECT
    person_skill_id,
    value,
    progress,
    skill.skill_id "skill.skill_id",
    skill.name "skill.name",
    skill.type "skill.type"
FROM person_skill
JOIN skill ON skill.skill_id = person_skill.skill_id
WHERE person_skill.person_id=$1

-- name: list-persons
SELECT person_id, name, email FROM person;

-- name: create-person
INSERT INTO person (
    person_id,
    name,
    email,
    phone,
    description,
    birthday,
    created_at
) values($1, $2, $3, $4, $5, $6, CURRENT_TIMESTAMP)
