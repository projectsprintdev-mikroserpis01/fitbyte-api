CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    name VARCHAR(100) DEFAULT NULL,
    preference VARCHAR(100) DEFAULT NULL,
    weight_unit VARCHAR(10) DEFAULT NULL,
    height_unit VARCHAR(10) DEFAULT NULL,
    weight INT DEFAULT NULL,
    height INT DEFAULT NULL,
    image_uri VARCHAR(4096) DEFAULT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create an index on the email field for login optimization
CREATE INDEX idx_email ON users(email);