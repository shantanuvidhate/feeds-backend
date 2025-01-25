CREATE TABLE IF NOT EXISTS comments (
    id BIGSERIAL PRIMARY KEY, -- This is the primary key of the comments table
    post_id BIGSERIAL NOT NULL, -- This is the post that the comment belongs to
    user_id BIGSERIAL NOT NULL, -- This is the user who made the comment
    content TEXT NOT NULL,
    created_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW()
    );