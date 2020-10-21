CREATE TABLE IF NOT EXISTS user_feeds
( id BIGSERIAL PRIMARY KEY,
  user_id int NOT NULL references users(id),
  engine_id int NOT NULL references engines(id),
  name varchar NOT NULL,
  created_at timestamp with time zone,
  updated_at timestamp with time zone,
  UNIQUE(user_id, name)
);

CREATE TABLE IF NOT EXISTS feed_subscriptions
( id BIGSERIAL PRIMARY KEY,
  feed_id int NOT NULL references user_feeds(id),
  source_type varchar NOT NULL,
  source_id int NOT NULL,
  created_at timestamp with time zone,
  updated_at timestamp with time zone,
  UNIQUE(feed_id, source_type, source_id)
);

ALTER TABLE users
  ADD COLUMN active_user_feed varchar NOT NULL;

ALTER TABLE users
  DROP COLUMN active_preference_set;