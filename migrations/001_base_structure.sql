-- Enable gen_random_uuid() in Postgres
CREATE EXTENSION IF NOT EXISTS pgcrypto;

-- Modify the users table (drop & recreate if early stage)
DROP TABLE IF EXISTS users;

CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    userName VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    firstName VARCHAR(100),
    lastName VARCHAR(100),
    createdAt TIMESTAMPTZ DEFAULT now(),
    updatedAt TIMESTAMPTZ DEFAULT now()
);

CREATE TABLE audit_logs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    requestId TEXT NOT NULL,
    requestIp INET,
    requestTimestamp TIMESTAMP WITH TIME ZONE,
    userId UUID,
    action TEXT NOT NULL,
    resource TEXT,
    extra JSONB,
    createdAt TIMESTAMPTZ NOT NULL DEFAULT now()
);
