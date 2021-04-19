CREATE TABLE IF NOT EXISTS project (
  project_id VARCHAR(255) NOT NULL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,

  created_at DATETIME,
  updated_at DATETIME
);

CREATE TABLE IF NOT EXISTS task (
  task_id VARCHAR(255) NOT NULL PRIMARY KEY,
  project_id  VARCHAR(255) NOT NULL REFERENCES project(project_id),
  person_id VARCHAR(255) REFERENCES person(person_id),

  name VARCHAR(255) NOT NULL,
  type VARCHAR(255) NOT NULL,
  description TEXT NOT NULL,

  complexity INTEGER,
  duration INTEGER,
  progress INTEGER,

  created_at DATETIME,
  updated_at DATETIME
);

ALTER TABLE person ADD COLUMN project_id VARCHAR(255) REFERENCES project(project_id);