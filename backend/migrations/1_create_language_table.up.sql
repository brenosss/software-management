CREATE TABLE IF NOT EXISTS language (
  language_id  VARCHAR(255) NOT NULL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
	popularity VARCHAR(255) NOT NULL,
	learning VARCHAR(255) NOT NULL,

  created_at DATETIME,
  updated_at DATETIME
);