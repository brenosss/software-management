CREATE TABLE IF NOT EXISTS person (
  person_id INTEGER PRIMARY KEY NOT NULL,
  name VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL,
  phone VARCHAR(255) NOT NULL,
  description TEXT NOT NULL,

  birthday DATETIME,
  created_at DATETIME,
  updated_at DATETIME
);

CREATE TABLE IF NOT EXISTS person_skill (
  person_skill_id INTEGER PRIMARY KEY NOT NULL,
  person_id INTEGER REFERENCES person(person_id),
  skill_id INTEGER REFERENCES skill(skill_id),

  value INTEGER,
  progress INTEGER,

  created_at DATETIME,
  updated_at DATETIME
);