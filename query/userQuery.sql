CREATE TABLE users(
    id VARCHAR(55) PRIMARY KEY,
    username VARCHAR(55) NOT NULL,
    email VARCHAR(55) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    address TEXT,
    photo VARCHAR(255) DEFAULT 'https://www.weact.org/wp-content/uploads/2016/10/Blank-profile.png',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    is_deleted BOOLEAN DEFAULT false
);
