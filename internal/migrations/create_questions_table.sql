CREATE TABLE questions (
    id SERIAL PRIMARY KEY,
    player_session_id INT NOT NULL,
    question_text VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_player_session
        FOREIGN KEY(player_session_id) 
        REFERENCES player_sessions(id) 
        ON DELETE CASCADE
);