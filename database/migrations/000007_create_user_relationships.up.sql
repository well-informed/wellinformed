CREATE TABLE IF NOT EXISTS user_relationships
(
  id BIGSERIAL PRIMARY KEY,
  follower_id int REFERENCES users(id),
  followee_id int REFERENCES users(id),
  created_at timestamp with time zone NOT NULL,
  updated_at timestamp with time zone NOT NULL,
  UNIQUE(follower_id, followee_id)
);