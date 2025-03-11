CREATE TABLE IF NOT EXISTS user_invitations (
    token bytea PRIMARY KEY,
    userId bigint NOT NULL
);