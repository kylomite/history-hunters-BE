-- Seed Players
INSERT INTO players (email, password_digest, avatar, score, created_at, updated_at)
VALUES
('player1@example.com', 'hashed_password_1', 'avatar1.png', 100, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('player2@example.com', 'hashed_password_2', 'avatar2.png', 200, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Seed Stages
INSERT INTO stages (title, background_img, difficulty, created_at, updated_at)
VALUES
('Easy Stage', 'stage1.png', 1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('Medium Stage', 'stage2.png', 2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('Hard Stage', 'stage2.png', 3, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Seed Player Sessions
INSERT INTO player_sessions (player_id, stage_id, lives, created_at, updated_at)
VALUES
(1, 1, 3, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(1, 3, 3, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(2, 2, 3, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Seed Questions
INSERT INTO questions (player_session_id, question_text, created_at, updated_at)
VALUES
(1, 'When did the Byzantine Empire collapse??', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(1, 'The collapse of the Soviet Union took place in which year?', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(2, 'Which of these facilities was not present on the Titanic?', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(2, 'The Thirty Years War ended with which treaty?', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(3, 'In the Seven Wonders of the World, which wonder is the only that has survived to this day?', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(3, 'During the Spanish Civil War (1936), Francisco Franco fought for which political faction?', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);


-- Seed Answers
INSERT INTO answers (question_id, answer_text, correct, created_at, updated_at)
VALUES
(1, '1453', true, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(1, '1299', false, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(1, '1891', false, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(1, '1990', false, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(2, '1991', true, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(2, '1995', false, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(2, '1981', false, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(2, '1990', false, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(3, 'Fainting room', true, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(3, 'Turkish baths', false, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(3, 'Kennel', false, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(3, 'Squash court', false, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(4, 'Peace of Westphalia', true, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(4, 'Treaty of Versailles', false, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(4, 'Treaty of Paris', false, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(4, 'Peace of Prague', false, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(5, 'Great Pyramids of Egypt', true, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(5, 'Colossus of Rhodes', false, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(5, 'Lighthouse of Alexandria', false, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(5, 'Statue of Zeus at Olympia', false, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(6, 'Nationalist Spain', true, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(6, 'Republican Spain', false, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(6, 'Popular Front', false, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(6, 'Papal State', false, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
