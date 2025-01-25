ALTER TABLE posts
    ADD COLUMN tags VARCHAR(255)[]; -- array of strings

ALTER TABLE posts
    ADD COLUMN updated_at timestamp(0) WITH TIME ZONE NOT NULL DEFAULT NOW(); -- timestamp of last update

    