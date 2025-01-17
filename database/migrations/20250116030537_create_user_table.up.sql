CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    name VARCHAR(100) DEFAULT '',
    preference VARCHAR(100) DEFAULT '',
    weight_unit VARCHAR(10) DEFAULT '',
    height_unit VARCHAR(10) DEFAULT '',
    weight INT DEFAULT 0,
    height INT DEFAULT 0,
    image_uri VARCHAR(4096) DEFAULT '',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create an index on the email field for login optimization
CREATE INDEX idx_email ON users(email);