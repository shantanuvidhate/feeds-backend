-- GIN stands for Generalized Inverted Index. GIN is designed for handling cases where the items to be indexed are composite values,
--  and the queries to be handled by the index need to search for element values that appear within the composite items. 
--  For example, the items could be documents, and the queries could be searches for documents containing specific words.

-- Uses gin because we have to search for substring in the content of the posts and comments. 

CREATE EXTENSION IF NOT EXISTS "pg_trgm";


-- indexes on content using gin
CREATE INDEX IF NOT EXISTS idx_comments_content ON comments USING gin (content gin_trgm_ops);

-- indexes on title using gin
CREATE INDEX IF NOT EXISTS idx_post_title ON posts USING gin (title gin_trgm_ops);

-- indexes on tags using gin
CREATE INDEX IF NOT EXISTS idx_posts_tags ON posts USING gin (tags);




-- indexes on username
CREATE INDEX IF NOT EXISTS idx_users_username ON users (username);

-- indexes on user_id
CREATE INDEX IF NOT EXISTS idx_posts_user_id ON posts (user_id);

-- indexes on post_id
CREATE INDEX IF NOT EXISTS idx_comments_post_id ON comments (post_id);