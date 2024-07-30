CREATE DATABASE applications;
\c applications

CREATE TABLE assignments (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE,
    position VARCHAR(255) NOT NULL,
    version INTEGER NOT NULL,
    doc BYTEA
);

CREATE TABLE applications (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    position VARCHAR(255) NOT NULL,
    assignment_id BIGINT NOT NULL,
    assignment_sent TIMESTAMP WITH TIME ZONE,
    FOREIGN KEY (assignment_id) REFERENCES assignments(id)
);

CREATE INDEX idx_assignments_deleted_at ON assignments(deleted_at);
CREATE INDEX idx_applications_deleted_at ON applications(deleted_at);

GRANT ALL PRIVILEGES ON DATABASE applications TO postgres;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO postgres;

-- assignments
INSERT INTO assignments(position, version, doc) VALUES ('Junior Go Developer', 1, null);
INSERT INTO assignments(position, version, doc) VALUES ('Junior Go Developer', 2, null);
INSERT INTO assignments(position, version, doc) VALUES ('Junior Go Developer', 3, null);