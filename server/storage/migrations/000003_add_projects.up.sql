CREATE TABLE IF NOT EXISTS projects (
    id          VARCHAR(36) PRIMARY KEY,
    name       VARCHAR(255) NOT NULL,
    description TEXT,
    owner_id       VARCHAR(255) NOT NULL,
    created_at  TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE users
ADD COLUMN projects text[];

ALTER TABLE bugs
ADD COLUMN project_id text[];