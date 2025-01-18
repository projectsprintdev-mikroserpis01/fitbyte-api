CREATE TABLE activities (
    id VARCHAR(255) PRIMARY KEY,
    activity_type SMALLINT NOT NULL,
    done_at TIMESTAMP NOT NULL,
    duration_in_minutes INT NOT NULL,
    user_id INT NOT NULL,
    calories_burned INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL
);

ALTER TABLE activities ADD CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES users(id);

CREATE INDEX idx_user_id ON activities(user_id);
CREATE INDEX idx_done_at ON activities(done_at);
CREATE INDEX idx_activity_type ON activities(activity_type);
CREATE INDEX idx_calories_burned ON activities(calories_burned);
