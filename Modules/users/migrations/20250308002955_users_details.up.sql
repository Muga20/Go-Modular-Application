CREATE TABLE users_details (
    id CHAR(36) PRIMARY KEY DEFAULT (UUID()), -- Unique identifier for the user details
    user_id CHAR(36) UNIQUE NOT NULL,         -- Foreign key to users table
    first_name VARCHAR(50) NOT NULL,          -- User's first name
    last_name VARCHAR(50) NOT NULL,           -- User's last name
    profile_pic TEXT,                         -- URL to the user's profile picture
    gender VARCHAR(10),                       -- User's gender
    date_of_birth TIMESTAMP,                  -- User's date of birth
    about_me TEXT,                            -- User's bio or description
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, -- Timestamp when the details were created
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, -- Timestamp when the details were last updated
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE -- Foreign key constraint
);