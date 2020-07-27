package database

type table struct {
	name string
	sql  string
}

//List of all table DDL in one place for use by database.CreateTables()
//Tables which depend on others must be lower in the list than others.
var tables = []table{
	{
		name: "users",
		sql: `
	CREATE TABLE IF NOT EXISTS users
	( id BIGSERIAL PRIMARY KEY,
		email varchar NOT NULL,
		first_name varchar NOT NULL,
		last_name varchar NOT NULL,
		user_name varchar NOT NULL,
		password varchar NOT NULL,
		active_preference_set varchar NOT NULL,
		created_at timestamp with time zone NOT NULL,
		updated_at timestamp with time zone NOT NULL
		)`,
	},
	{
		name: "src_rss_feeds",
		sql: `CREATE TABLE IF NOT EXISTS src_rss_feeds
	(	id SERIAL PRIMARY KEY,
		title varchar NOT NULL,
		description varchar,
		link varchar UNIQUE NOT NULL,
		feed_link varchar UNIQUE NOT NULL ,
		updated timestamp with time zone,
		last_fetched_at timestamp with time zone,
		language varchar,
		generator varchar
	)`,
	},
	{
		name: "user_subscriptions",
		sql: `CREATE TABLE IF NOT EXISTS user_subscriptions
	( id BIGSERIAL PRIMARY KEY,
		user_id int REFERENCES users(id),
		source_id int REFERENCES src_rss_feeds(id),
		created_at timestamp with time zone,
		UNIQUE(user_id, source_id)
	)`},
	{
		name: "content_items",
		sql: `CREATE TABLE IF NOT EXISTS content_items
	( id BIGSERIAL PRIMARY KEY,
	  source_id int NOT NULL REFERENCES src_rss_feeds(id),
	  source_title varchar NOT NULL,
	  source_link varchar NOT NULL,
	  title varchar NOT NULL,
	  description varchar,
	  content varchar,
	  link varchar NOT NULL,
	  updated timestamp with time zone,
	  published timestamp with time zone,
	  author varchar,
	  guid varchar NOT NULL,
	  image_title varchar,
	  image_url varchar,
	  UNIQUE (source_id, link)
	)`,
	},
	{
		name: "preference_sets",
		sql: `CREATE TABLE IF NOT EXISTS preference_sets
		( id BIGSERIAL PRIMARY KEY,
			user_id int NOT NULL REFERENCES users(id),
			name varchar NOT NULL,
			sort varchar NOT NULL,
			start_date timestamp with time zone,
			end_date timestamp with time zone,
			UNIQUE (user_id, name)
		)`,
	},
}
