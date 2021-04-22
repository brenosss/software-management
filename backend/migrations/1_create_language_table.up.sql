CREATE TABLE IF NOT EXISTS language (
  language_id INTEGER PRIMARY KEY NOT NULL,
  name VARCHAR(255) NOT NULL,
	popularity VARCHAR(255) NOT NULL,
	learning VARCHAR(255) NOT NULL,

  created_at DATETIME,
  updated_at DATETIME
);