CREATE TABLE IF NOT EXISTS skill (
  id  VARCHAR(255) NOT NULL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
	type VARCHAR(255) NOT NULL,

  created_at DATETIME,
  updated_at DATETIME
);