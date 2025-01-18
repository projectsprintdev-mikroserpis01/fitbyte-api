CREATE TABLE activity (
    id VARCHAR(255) PRIMARY KEY,
    activity_type SMALLINT NOT NULL,
    done_at TIMESTAMP NOT NULL,
    duration_in_minutes INT NOT NULL,
    user_id INT NOT NULL,
    calories_burned INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL
);

ALTER TABLE activity ADD CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES users(id);

CREATE INDEX idx_user_id ON activity(user_id);
CREATE INDEX idx_done_at ON activity(done_at);
CREATE INDEX idx_activity_type ON activity(activity_type);
CREATE INDEX idx_calories_burned ON activity(calories_burned);

