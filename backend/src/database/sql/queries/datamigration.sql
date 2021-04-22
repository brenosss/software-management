-- datamigration.sql

-- name: populate-languages
INSERT INTO language(name, popularity, learning, created_at) VALUES
('java', 'very high', 'moderate', CURRENT_TIMESTAMP),
('c', 'medium', 'moderate', CURRENT_TIMESTAMP),
('python', 'very high', 'moderate', CURRENT_TIMESTAMP),
('c++', 'high', 'hard', CURRENT_TIMESTAMP),
('c#', 'high', 'moderate', CURRENT_TIMESTAMP),
('php', 'high', 'easy', CURRENT_TIMESTAMP),
('javascript', 'very high', 'moderate', CURRENT_TIMESTAMP),
('sql', 'very high', 'moderate', CURRENT_TIMESTAMP),
('golang', 'low', 'moderate', CURRENT_TIMESTAMP);

-- name: populate-skills
INSERT INTO skill(name, type, created_at) VALUES
('Algorithms', 'tech', CURRENT_TIMESTAMP),
('User experience', 'tech', CURRENT_TIMESTAMP),
('Designer', 'tech', CURRENT_TIMESTAMP),
('Programming languages', 'tech', CURRENT_TIMESTAMP),
('Mentoring', 'soft', CURRENT_TIMESTAMP),
('Write', 'soft', CURRENT_TIMESTAMP),
('Speaking', 'soft', CURRENT_TIMESTAMP),
('Planing', 'soft', CURRENT_TIMESTAMP),
('Creativity', 'personality', CURRENT_TIMESTAMP),
('Focus', 'personality', CURRENT_TIMESTAMP),
('Proactivity', 'personality', CURRENT_TIMESTAMP);