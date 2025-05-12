-- Enable gen_random_uuid() in Postgres
CREATE EXTENSION IF NOT EXISTS pgcrypto;

-- Modify the users table (drop & recreate if early stage)
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS audit_logs;
DROP TABLE IF EXISTS patients;
DROP TABLE IF EXISTS patient_images;

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

CREATE TABLE patients (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    userId UUID REFERENCES users(id),
    firstName VARCHAR(100) NOT NULL,
    lastName VARCHAR(100) NOT NULL,
    dateOfBirth DATE,
    gender VARCHAR(20),
    phoneNumber VARCHAR(20),
    email VARCHAR(100),
    notes TEXT,
    createdAt TIMESTAMPTZ DEFAULT now(),
    updatedAt TIMESTAMPTZ DEFAULT now()
);

CREATE TABLE patient_images (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    patientId UUID NOT NULL REFERENCES patients(id) ON DELETE CASCADE,
    imageData BYTEA NOT NULL,  -- stores the actual image binary data
    fileType VARCHAR(20) CHECK (fileType IN ('png', 'jpeg', 'jpg', 'bmp', 'tiff')),
    description TEXT,
    uploadedAt TIMESTAMPTZ DEFAULT now(),
    createdAt TIMESTAMPTZ DEFAULT now(),
    updatedAt TIMESTAMPTZ DEFAULT now()
);
