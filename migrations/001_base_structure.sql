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

CREATE TABLE auditlogs (
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
