CREATE TABLE users (
    id CHAR(36) PRIMARY KEY DEFAULT (UUID()), -- Unique identifier for the user
    email VARCHAR(255) UNIQUE NOT NULL,       -- User's email address
    phone VARCHAR(15) UNIQUE,                 -- User's phone number
    username VARCHAR(50) UNIQUE NOT NULL,     -- User's username
    password VARCHAR(255) NOT NULL,           -- User's password (hashed)
    auth_type VARCHAR(50) NOT NULL DEFAULT 'email', -- Authentication type (e.g., email, google)
    is_verified BOOLEAN NOT NULL DEFAULT FALSE, -- Whether the user is verified
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, -- Timestamp when the user was created
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, -- Timestamp when the user was last updated
    deleted_at TIMESTAMP                      -- Timestamp when the user was deleted (soft delete)
);