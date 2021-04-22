CREATE TABLE IF NOT EXISTS project (
  project_id INTEGER PRIMARY KEY NOT NULL,
  name VARCHAR(255) NOT NULL,

  created_at DATETIME,
  updated_at DATETIME
);

CREATE TABLE IF NOT EXISTS task (
  task_id INTEGER PRIMARY KEY NOT NULL,
  project_id INTEGER NOT NULL REFERENCES project(project_id),
  person_id INTEGER REFERENCES person(person_id),

  name VARCHAR(255) NOT NULL,
  type VARCHAR(255) NOT NULL,
  description TEXT NOT NULL,

  complexity INTEGER,
  duration INTEGER,
  progress INTEGER,

  created_at DATETIME,
  updated_at DATETIME
);

ALTER TABLE person ADD COLUMN project_id INTEGER REFERENCES project(project_id);