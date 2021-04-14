migrations-up:
	sqlite3 software.db < migrations/1_create_language_table.up.sql
	sqlite3 software.db < migrations/2_create_skill_table.up.sql

migrations-down:
	sqlite3 software.db < migrations/1_create_language_table.down.sql
	sqlite3 software.db < migrations/2_create_skill_table.down.sql