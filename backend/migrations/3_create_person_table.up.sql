CREATE TABLE IF NOT EXISTS person (
  person_id VARCHAR(255) NOT NULL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL,
  phone VARCHAR(255) NOT NULL,
  description TEXT NOT NULL,

  birthday DATETIME,
  created_at DATETIME,
  updated_at DATETIME
);

CREATE TABLE IF NOT EXISTS person_skill (
  person_skill_id VARCHAR(255) NOT NULL PRIMARY KEY,
  person_id  VARCHAR(255) REFERENCES person(person_id),
  skill_id VARCHAR(255) REFERENCES skill(skill_id),

  value INTEGER,
  progress INTEGER,

  created_at DATETIME,
  updated_at DATETIME
);