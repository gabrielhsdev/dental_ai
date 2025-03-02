CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    userName VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    firstName VARCHAR(100),
    lastName VARCHAR(100),
    createdAt TIMESTAMPTZ,
    updatedAt TIMESTAMPTZ
);